package req

import (
	"encoding/xml"
	"lastfm/lib"
)

type UserInfo struct {
	XMLName    xml.Name       `xml:"user"`
	Name       string         `xml:"name"`
	Images     []lib.Image    `xml:"image"`
	URL        string         `xml:"url"`
	Country    string         `xml:"country"`
	Age        uint8          `xml:"age"`
	Gender     string         `xml:"gender"`
	Subscriber bool           `xml:"subscriber"`
	Playcount  uint           `xml:"playcount"`
	Playlists  uint           `xml:"playlists"`
	Boostrap   uint           `xml:"bootstrap"`
	Registered RegisteredTime `xml:"registered"`
}

type RegisteredTime struct {
	XMLName   xml.Name `xml:"registered"`
	UnixTime  string   `xml:"unixtime,attr"`
	InnerTime string   `xml:",innerxml"`
}
