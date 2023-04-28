package lib

import "encoding/xml"

type Image struct {
	XMLName   xml.Name `xml:"image"`
	Fulltrack string   `xml:"size,attr"`
	URL       string   `xml:",innerxml"`
}

type Artist struct {
	XMLName xml.Name `xml:"artist"`
	Name    string   `xml:",innerxml"`
	MBID    string   `xml:"mbid"`
}

type Streamable struct {
	XMLName   xml.Name `xml:"streamable"`
	Fulltrack string   `xml:"fulltrack,attr"`
	Value     uint     `xml:",innerxml"`
}

type Track struct {
	XMLName    xml.Name   `xml:"track"`
	NowPlaying bool       `xml:"nowplaying,attr"`
	Rank       string     `xml:"rank,attr"`
	Name       string     `xml:"name"`
	Duration   uint       `xml:"duration"`
	Playcount  uint       `xml:"playcount"`
	MBID       string     `xml:"mbid"`
	URL        string     `xml:"url"`
	Streamable Streamable `xml:"streamable"`
	Artist     Artist     `xml:"artist"`
	Image      Image      `xml:"image"`
}
