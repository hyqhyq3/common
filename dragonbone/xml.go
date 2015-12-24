package dragonbone

import (
	"encoding/xml"
	"io/ioutil"
)

type SubTexture struct {
	Name    string  `xml:"name,attr"`
	X       float32 `xml:"x,attr"`
	Y       float32 `xml:"y,attr"`
	Width   float32 `xml:"width,attr"`
	Height  float32 `xml:"height,attr"`
	Rotated bool    `xml:"rotated,attr"`
}

type TextureAtlas struct {
	SubTexture []SubTexture
	ImagePath  string `xml:"imagePath,attr"`
	Width      int    `xml:"-"`
	Height     int    `xml:"-"`
}

func DecodeXml(path string) (texture *TextureAtlas, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = xml.Unmarshal(data, &texture)
	return
}
