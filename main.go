package main

import (
	"encoding/json"
	"fmt"
)

// 함수 2024-02-24
// 함수 정의
func add(x int, y int) int {
	return x + y
}

// 두개의 값을 리턴할 수 있다.
func calc(x int, y int) (int, int) {
	return x + y, x - y
}

// 반환할 변수를 지정할 수 있다.
func returnVal(x int, y int) (val int) {
	val = x + y
	return // return 할 게 무엇인지 명시하지 않아도 된다.
}

// 함수는 scope 범위를 가진다.
func myName() string {
	name := "jack"
	return name
}

// 함수는 값을 복사해서 사용한다. -> 값을 변경할 수 없다.
func addOne(num int) {
	println("addOne 함수 호출 num: ", num)
	num += 1
	println("increase num : ", num, "in addOne function")
}

// 포인터를 이용해서 값을 변경할 수 있다.
func addOneByPointer(num *int) {
	println("addOneByPointer 함수 호출 num: ", *num)
	*num += 1
	println("increase num : ", *num, "in addOneByPointer function")
}

// Calculator /**
// 메소드 2024-02-24
// 메소드는 구조체에 속한 함수이다.
//
//	func (리시버) 메소드명(매개변수 매개변수타입) 반환타입 {
//		코드
//		return 반환값
//	}
type Calculator struct {
	X int
}

func (c Calculator) add(y int) int {
	return c.X + y
}

// 리시버를 포인터로 받으면 값이 변경된다.
func (c *Calculator) addByPointer(y int) {
	c.X += y
}

func main() {

	// 구조체 정의
	type Vertex struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	var v Vertex = Vertex{1, 2}
	fmt.Println(" ======= 구조체 정의 후 json 형태로 변환 =======")
	data, _ := json.Marshal(v)
	fmt.Println(string(data))

	// 포인터로 값 변경
	fmt.Println(" ======= 포인터로 값 변경 =======")
	num := 1
	fmt.Printf("before num: %d\n", num)
	numAddress := &num
	*numAddress = 2
	fmt.Printf("after num: %d\n", num)

	fmt.Println(" ======= 주소타입으로 값 변경 =======")
	var numPointer *int = &num // int 의 포인터를 만들어서 num 의 주소를 할당
	fmt.Println("num address: ", &num)
	fmt.Println("num value by reference : ", *numPointer)

	// 함수 호출
	fmt.Println(" ======= add 함수 호출 =======")
	addResult := add(42, 13)
	fmt.Println("addResult: ", addResult)

	fmt.Println(" ======= calc 함수 호출 =======")
	calcResult1, calcResult2 := calc(42, 13)
	fmt.Println("calcResult1: ", calcResult1, "calcResult2: ", calcResult2)

	fmt.Println(" ======= returnVal 함수 호출 =======")
	returnValResult := returnVal(42, 13)
	fmt.Println("returnValResult: ", returnValResult)

	fmt.Println(" ======= name 변수 호출 : fail =======")
	// fmt.Println("name: ", name) // name 은 함수 내부에서만 사용 가능한 변수이다.

	fmt.Println(" ======= addOne 함수 호출 =======")
	fmt.Println("함수는 값을 복사해서 사용한다. -> 값을 변경할 수 없다.")
	paramNum := 1
	fmt.Println("before paramNum: ", paramNum)
	addOne(paramNum)
	fmt.Println("after paramNum: ", paramNum)

	fmt.Println(" ======= addOneByPointer 함수 호출 =======")
	fmt.Println("before paramNum: ", paramNum)
	addOneByPointer(&paramNum)
	fmt.Println("after paramNum: ", paramNum)

	// 메소드 호출
	fmt.Println(" ======= 메소드 호출 =======")
	fmt.Println(" ======= 리시버를 구조체로 받으면 값이 변경되지 않는다. =======")
	calc := Calculator{10}
	val := calc.add(20)
	fmt.Println("val: ", val)
	fmt.Println("calc.X: ", calc.X)

	fmt.Println(" ======= 리시버를 포인터로 받으면 값이 변경된다. =======")
	calc.addByPointer(20)
	fmt.Println("calc.X: ", calc.X)

}
