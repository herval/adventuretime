package graphics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FramesLoader struct {
}

func (l *FramesLoader) Parse(fileName string) map[string]Frame {
	path, e := filepath.Abs(fileName)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	var data jsonobject
	json.Unmarshal(file, &data)

	// util.Debug(fmt.Sprintf("Results: %+v\n", data.Frames))

	return data.Frames
}

type Dimensions struct {
	X int
	Y int
	W int
	H int
}

type Frame struct {
	Dimensions Dimensions `json:"frame"`
}

type jsonobject struct {
	Frames map[string]Frame
}
