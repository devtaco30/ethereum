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

// 인터페이스 2024-02-24
type Ducky interface {
	DuckSound() string
	DuckWalk() string
	isSwim() string
}

type Bird struct {
	name string
}

func (b Bird) DuckSound() string {
	return "꽥꽥"
}

func (b Bird) DuckWalk() string {
	return "뒤뚱뒤뚱"
}

func (b Bird) isSwim() string {
	if b.name == "오리" {
		return "있다."
	} else {
		return "없다."
	}
}

// 빈 인터페이스를 매개변수로 하면, 어떤 값이든 받을 수 있다.
func typeCheck(i interface{}) {
	fmt.Printf("%T\n", i)
}

// 패키지
// 패키지는 코드를 구조화하고 재사용하기 위한 방법이다.

// defer : 자신을 둘러싼 함수가 종료할 때 까지 실행을 연기한다.
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// 주의점 : 함수의 매개변수는 호출시 이 값을 '복사' 한다.
func a() {
	i := 0
	defer fmt.Println(i)
	i++
}

// defer는 반환된 값을 읽고, 반환값을 지정한 변수에 할당한다.
func c() (i int) {
	defer func() { i++ }()
	return 1
}

// 가변 인자
// -> 슬라이스로 받는다.
func numbers(nums ...int) {
	fmt.Println(nums)
	fmt.Printf("%T \n", nums)
}

// 함수 리터럴 (이름 없는 함수 -> 익명함수)
// 익명함수 -> 데이터타입처럼 파라미터로 던질 수도 있다.
//익명 함수는 포인터로도 사용할 수 있다.

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

	// 인터페이스
	duck := Bird{name: "오리"}
	fmt.Printf("%s 는 %s 하고 울며, %s 걸으며, 수영을 할 수 %s \n", duck.name, duck.DuckSound(), duck.DuckWalk(), duck.isSwim())

	var n1 int8 = 1
	var n2 int16 = 2
	var n3 int32 = 3
	typeCheck(n1)
	typeCheck(n2)
	typeCheck(n3)

	// defer -> stack 구조로 쌓인다.
	fmt.Println(" ======= defer =======")
	defer fmt.Println("마지막 실행")
	defer fmt.Println("세 번째 실행")
	defer fmt.Println("두 번째 실행")
	fmt.Println("첫 번째 실행")

	fmt.Println(" ======= defer for =======")
	b()

	fmt.Println(" ======= defer는 매개변수를 복사한다 =======")
	a()

	fmt.Println(" ======= defer는 반환된 값을 읽고, 반환값을 지정한 변수에 할당한다. =======")
	c()

	fmt.Println(" ======= 가변 인자 =======")
	numbers(1, 2, 3, 4, 5) // -> retunr 값이 [] slice 로 된다.

	// 함수 리터럴
	fmt.Println(" ======= 함수 리터럴 =======")
	f := func() {
		fmt.Println("Hello, World!")
	}
	f()

	func() {
		fmt.Println("Hello, World!")
	}() // -> 익명함수를 바로 호출할 수 있다.

	fmt.Println(" ======= 익명함수를 파라미터로 넘기기 =======")
	i := 0
	func() { // 여기서는 값이 증가한다. 일반 함수에 포인터형태로 넘겨야만 가능하던 것을, 익명함수는 가능하게 한다.
		i++
	}()
	fmt.Println("i: ", i)

}
