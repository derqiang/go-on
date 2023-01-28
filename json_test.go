package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

type Element map[string]interface{}
type AnalyseInfo struct {
	CountInfo []struct {
		Color string  `json:"color"`
		Id    int     `json:"id"`
		RectH float32 `json:"rect_h"`
		RectW float32 `json:"rect_w"`
		RectX float32 `json:"rect_x"`
		RectY float32 `json:"rect_y"`
		State bool    `json:"state"`
	} `json:"count_info"`
	Meta struct {
		AnalyseTime string `json:"analyse_time"`
		ScanTime    string `json:"scan_time"`
	} `json:"meta"`
	QualityInfo string   `json:"quality_info"`
	SegmInfo    SegmInfo `json:"segm_info"`
}
type SegmInfo []struct {
	Angle    string  `json:"angle"`
	ClassId  string  `json:"class_id"`
	Color    string  `json:"color"`
	Id       int     `json:"id"`
	MaskPath string  `json:"mask_path"`
	RectH    float32 `json:"rect_h"`
	RectW    float32 `json:"rect_w"`
	RectX    float32 `json:"rect_x"`
	RectY    float32 `json:"rect_y"`
}

func TestJsonDiff(t *testing.T) {
	var root = "/Users/dongfuqiang/Desktop/PAN/qt/work/B20212318/slide2/cell23"
	var p1 = root + "/processingData/analyse.json"
	var p2 = root + "/analyse/analyse.json"

	var a1 = parseAndExtractFile(p1)
	var a2 = parseAndExtractFile(p2)

	sort.Slice(a1.SegmInfo, func(i, j int) bool {
		return a1.SegmInfo[i].Id-a1.SegmInfo[j].Id < 0
	})
	sort.Slice(a2.SegmInfo, func(i, j int) bool {
		return a2.SegmInfo[i].Id-a2.SegmInfo[j].Id < 0
	})

	r1 := ""
	r2 := ""

	if len(a1.SegmInfo) != len(a2.SegmInfo) {
		println(" 数目不等 ")
		return
	}

	for i, s := range a1.SegmInfo {
		//fmt.Printf(" >> %v ", s.Angle)

		r1 = fmt.Sprintf("id:%v,x:%v,y:%v,w:%v,h:%v,a:%v;", s.Id, s.RectX, s.RectY, s.RectW, s.RectH, s.Angle)

		ss := a2.SegmInfo[i]
		r2 = fmt.Sprintf("id:%v,x:%v,y:%v,w:%v,h:%v,a:%v;", ss.Id, ss.RectX, ss.RectY, ss.RectW, ss.RectH, ss.Angle)

		if r1 != r2 {
			fmt.Println(r1 + "\nVS\n" + r2 + "\n\n")
		}
	}
}

func parseAndExtractFile(fileLocation string) *AnalyseInfo {

	file, err := os.OpenFile(fileLocation, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("error reading!")
	}
	reader := bufio.NewReader(file)
	var analyse AnalyseInfo
	data, _ := ioutil.ReadAll(reader)
	_ = json.Unmarshal(data, &analyse)

	return &analyse
}

func TestStringOperator(t *testing.T) {

	input_str := "This is my country."
	fmt.Println(input_str[:2])
	input_str = "这是我的祖国。"

	fmt.Println(string([]rune(input_str)[:2]))
	t.Failed()

	arr := []string{"A"}
	arr = append(arr, "B")
	fmt.Println(arr)
	//t.Error(fmt.Sprintf("This is %v", time.Second))
	t.Log(" This is test log ... ")
}
