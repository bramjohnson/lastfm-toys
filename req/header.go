package req

import (
	"encoding/xml"
)

type LFMError struct {
	XMLName xml.Name `xml:"error"`
	Code    string   `xml:"code,attr"`
	Error   string   `xml:",innerxml"`
}

type LastFMStatus struct {
	XMLName xml.Name `xml:"lfm"`
	Status  string   `xml:"status,attr"`
	Body    []byte   `xml:",innerxml"`
}
