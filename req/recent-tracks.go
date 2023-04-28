package req

import (
	"encoding/xml"
	"lastfm/lib"
)

type RecentTracks struct {
	XMLName    xml.Name    `xml:"recenttracks"`
	User       string      `xml:"user,attr"`
	Page       uint        `xml:"page,attr"`
	PerPage    uint        `xml:"perPage,attr"`
	TotalPages uint        `xml:"totalPages,attr"`
	Total      uint        `xml:"total,attr"`
	Tracks     []lib.Track `xml:"track"`
}
