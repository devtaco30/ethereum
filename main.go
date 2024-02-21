package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 구조체 정의
	type Vertex struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	var v Vertex = Vertex{1, 2}

	data, _ := json.Marshal(v)
	fmt.Println(string(data))

}
