package cocos

import (
	"encoding/json"
	"io/ioutil"
)

type Armature struct {
	AnimationData []AnimationData `json:"animation_data"`
}

type AnimationData struct {
	Name    string    `json:"name"`
	MovData []MovData `json:"mov_data"`
}

type MovData struct {
	Name        string        `json:"name"`
	MovBoneData []MovBoneData `json:"mov_bone_data"`
}

type MovBoneData struct {
	Name      string      `json:"name"`
	FrameData []FrameData `json:"frame_data"`
}

type FrameData struct {
	Event      string `json:"evt"`
	FrameIndex int    `json:"fi"`
}

func DecodeArmatureJsonFile(path string) (a *Armature, e error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &a)
	return
}
