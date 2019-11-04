package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "https://shiro.ch/response.xml"

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	xml := parseXML(body)
	fmt.Println(xml.Body)
}

// Note is generated with https://www.onlinetool.io/xmltogo/
type Note struct {
	XMLName xml.Name `xml:"note"`
	Text    string   `xml:",chardata"`
	To      string   `xml:"to"`
	From    string   `xml:"from"`
	Heading string   `xml:"heading"`
	Body    string   `xml:"body"`
}

func parseXML(b []byte) Note {
	var note Note
	err := xml.Unmarshal(b, &note)
	if err != nil {
		panic(err)
	}
	return note
}
