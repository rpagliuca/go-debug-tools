package tools

import (
	"encoding/json"
	"fmt"
	"time"
)

func JsonEncode(anything interface{}) string {
	str, err := json.MarshalIndent(anything, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(str)
}

type statsItem struct {
	Count                int
	TotalEllapsedSeconds float64
}

var statsData = map[string]*statsItem{}

func Lap(name string, fn func()) {
	if _, ok := statsData[name]; !ok {
		statsData[name] = &statsItem{}
	}
	before := time.Now()
	fn()
	after := time.Now()
	ellapsed := after.Sub(before).Seconds()
	statsData[name].Count += 1
	statsData[name].TotalEllapsedSeconds += ellapsed
}

func PrintStats(name string) {
	fmt.Println("Stats for", name, ":", JsonEncode(statsData[name]))
}
