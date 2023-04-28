package main

import (
	"fmt"
	"io/ioutil"
	"lastfm/lib"
	"lastfm/req"
	"lastfm/util"
	"os"
	"strconv"
	"time"
)

func getExampleResp() []byte {
	xmlFile, err := os.Open("example.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened example.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	return byteValue
}

func getUserInfo(user string) req.UserInfo {
	// Make header request to lastfm API
	resp := util.GetReq(fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getInfo&api_key=%s&user=%s", util.GetConfig().API_KEY, user))
	lfm := util.SafeUnmarshal(resp, req.LastFMStatus{})

	// If request failed...
	if lfm.Status == "failed" {
		lfmerror := util.SafeUnmarshal(lfm.Body, req.LFMError{})
		panic(fmt.Sprintf("Last FM API call failed: %s, %s", lfmerror.Error, lfmerror.Code))
	}

	// Request suceeded!
	userinfo := util.SafeUnmarshal(lfm.Body, req.UserInfo{})

	return *userinfo
}

func getRecentTracks(timestamp_from string, timestamp_to string, page uint) []byte {
	// Timestamp shenanigans
	cache_file := fmt.Sprintf("rsc/cache/%s-%s-%d.xml", timestamp_from, timestamp_to, page)

	// File exists, read and return
	if _, err := os.Stat(cache_file); err == nil {
		resp, err := os.ReadFile(cache_file)
		util.Check(err)
		return resp
	}

	// File does not exist, query and create cache
	// API call will never result in different results due to the nature of time
	fmt.Println("Cache does not exist for timestamps, querying and creating cache...")
	resp := util.GetReq(fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getRecentTracks&api_key=%s&user=bratevi&from=%s&to=%s&page=%d", util.GetConfig().API_KEY, timestamp_from, timestamp_to, page))

	// Write to cache
	err := os.WriteFile(cache_file, resp, 0644)
	util.Check(err)

	return resp
}

func getRecentTracksObject(timestamp_from string, timestamp_to string, page uint) req.RecentTracks {
	// Make header request to lastfm API
	lfm := util.SafeUnmarshal(getRecentTracks(timestamp_from, timestamp_to, page), req.LastFMStatus{})

	// If request failed...
	if lfm.Status == "failed" {
		lfmerror := util.SafeUnmarshal(lfm.Body, req.LFMError{})
		panic(fmt.Sprintf("Last FM API call failed: %s, %s", lfmerror.Error, lfmerror.Code))
	}

	// Request suceeded!
	recenttracks := util.SafeUnmarshal(lfm.Body, req.RecentTracks{})

	return *recenttracks
}

func getRecentTracksAllPages(timestamp_from string, timestamp_to string) []lib.Track {
	firstPage := getRecentTracksObject(timestamp_from, timestamp_to, 1)
	curPageNumber := firstPage.Page
	totalPageNumber := firstPage.TotalPages

	reqs := firstPage.Tracks

	for curPageNumber < totalPageNumber {
		curPage := getRecentTracksObject(timestamp_from, timestamp_to, curPageNumber+1)
		reqs = append(reqs, curPage.Tracks...)
		curPageNumber = curPage.Page
	}

	return reqs
}

func lastMonthDiscovered(user string) map[util.MonthRange][]lib.Track {
	// Get the month the user registered
	userInfo := getUserInfo(user)
	userRegisteredTime, err := strconv.ParseInt(userInfo.Registered.UnixTime, 10, 64)
	if err != nil {
		panic(err)
	}
	userRegisteredUnix := time.Unix(userRegisteredTime, 0)

	// Start at registration month, and go until current month
	monthRange := util.GetMonthRangeDate(userRegisteredUnix)
	lastMonthRange := util.GetMonthRange()
	discoveredMap := make(map[lib.Track]util.MonthRange)
	reversedMap := make(map[util.MonthRange][]lib.Track)
	for monthRange.StartOfMonth != lastMonthRange.EndOfMonth {
		fmt.Println(monthRange)
		reversedMap[monthRange] = []lib.Track{}
		monthTracks := getRecentTracksAllPages(fmt.Sprint(monthRange.StartOfMonth.Unix()), fmt.Sprint(monthRange.EndOfMonth.Unix()))

		// Count tracks
		trackMap := make(map[lib.Track]uint)
		for _, track := range monthTracks {
			_, ok := trackMap[track]
			if !ok {
				trackMap[track] = 0
			}
			if !track.NowPlaying {
				trackMap[track] += 1
			}
		}

		// For all tracks, see if
		for track := range trackMap {
			if trackMap[track] >= 3 {
				_, ok := discoveredMap[track]
				// Song is not yet discovered, time to add it!
				if !ok {
					discoveredMap[track] = monthRange
				}
			}
		}

		monthRange = monthRange.NextMonth()
	}

	// Reverse map to go from MonthRange -> []Tracks
	for track := range discoveredMap {
		reversedMap[discoveredMap[track]] = append(reversedMap[discoveredMap[track]], track)
	}

	return reversedMap
}

func main() {
	// Let's find the start and end of the current month, and go from there
	discovered := lastMonthDiscovered("bratevi")
	discoveredThisMonth := discovered[util.GetMonthRangeDate(time.Date(2022, 04, 28, 0, 0, 0, 0, time.Local))]
	for _, track := range discoveredThisMonth {
		fmt.Println(fmt.Sprintf("%s - %s", track.Name, track.Artist.Name))
	}
}
