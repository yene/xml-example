package main

import (
	"io/ioutil"
	"testing"
)

func TestXMLParsing(t *testing.T) {
	t.Run("Testing parsing against file", func(t *testing.T) {
		dat, err := ioutil.ReadFile("response.xml")
		if err != nil {
			t.Error("Test not setup up correctly, could not find file")
			return
		}

		xml := parseXML(dat)
		if xml.Body != "Don't forget me this weekend!" {
			t.Error("Invalid XML body")
		}
	})
}
