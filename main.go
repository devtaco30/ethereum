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

	// 포인터로 값 변경
	num := 1
	fmt.Println("before num: %d", num)
	numAddress := &num
	*numAddress = 2
	fmt.Println("after num: %d", num)

	var numPointer *int = &num // int 의 포인터를 만들어서 num 의 주소를 할당
	fmt.Println("num address: ", &num)
	fmt.Println("num value by reference : ", *numPointer)

}
