package dragonbone

import (
	"encoding/xml"
	"io/ioutil"
)

type Transform struct {
	X   float64 `xml:"x,attr"`
	Y   float64 `xml:"y,attr"`
	SkX float64 `xml:"skX,attr"`
	SkY float64 `xml:"skY,attr"`
	ScX float64 `xml:"scX,attr"`
	ScY float64 `xml:"scY,attr"`
	PX  float64 `xml:"pX,attr"`
	PY  float64 `xml:"pY,attr"`
}

type ColorTransform struct {
	AO int `xml:"aO,attr"`
	RO int `xml:"rO,attr"`
	GO int `xml:"gO,attr"`
	BO int `xml:"bO,attr"`
	AM int `xml:"aM,attr"`
	RM int `xml:"rM,attr"`
	GM int `xml:"gM,attr"`
	BM int `xml:"bM,attr"`
}

type Bone struct {
	Name           string         `xml:"name,attr"`
	Transform      Transform      `xml:"transform"`
	ColorTransform ColorTransform `xml:"colorTransform"`
}

type Display struct {
	Name      string    `xml:"name,attr"`
	Type      string    `xml:"type,attr"`
	Transform Transform `xml:"transform"`
}

type Slot struct {
	Name      string  `xml:"name,attr"`
	Parent    string  `xml:"parent,attr"`
	Z         int     `xml:"z,attr"`
	BlendMode string  `xml:"blendMode,attr"`
	Display   Display `xml:"display"`
}

type Skin struct {
	Name string `xml:"name,attr"`
	Slot []Slot `xml:"slot"`
}

type Frame struct {
	Event          string         `xml:"event,attr"`
	Z              int            `xml:"z,attr"`
	TweenEasing    int            `xml:"tweenEasing,attr"`
	Duration       int            `xml:"duration,attr"`
	DisplayIndex   int            `xml:"displayIndex,attr"`
	Transform      Transform      `xml:"transform"`
	ColorTransform ColorTransform `xml:"colorTransform"`
}

type Timeline struct {
	Frames []Frame `xml:"frame"`
	Name   string  `xml:"name,attr"`
	Scele  int     `xml:"scele,attr"`
	Offset int     `xml:"offset,attr"`
}

type Animation struct {
	Name        string     `xml:"name,attr"`
	Timelines   []Timeline `xml:"timeline"`
	Loop        int        `xml:"loop,attr"`
	FadeInTime  int        `xml:"fadeInTime,attr"`
	Duration    int        `xml:"duration,attr"`
	Scale       int        `xml:"scale,attr"`
	AutoTween   int        `xml:"autoTween,attr"`
	TweenEasing string     `xml:"tweenEasing,attr"`
}

type Armature struct {
	Name       string      `xml:"name,attr"`
	Bones      []Bone      `xml:"bone"`
	Skin       Skin        `xml:"skin"`
	Animations []Animation `xml:"animation"`
}

type DragonBones struct {
	Name      string   `xml:"name,attr"`
	FrameRate int      `xml:"frameRate,attr"`
	Version   string   `xml:"version,attr"`
	Armature  Armature `xml:"armature"`
}

func DecodeSkeleton(path string) (db *DragonBones, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = xml.Unmarshal(data, &db)
	return
}
