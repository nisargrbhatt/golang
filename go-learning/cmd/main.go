package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Variable starts with Uppercase name will be exported to package level
var ExportedVariable int = 12

// Kind of like ENUM
const (
	a = iota
	b
)

func variables() {

	// Will set zero value of assigned type
	var i int
	fmt.Println(i)
	var j int = 34
	// := initialize and need value, which will be used to auto infer type
	k := 48

	i = 42

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	i = 27

	fmt.Println(i)
}

func variableConversion() {
	i := 42

	var j float32
	j = float32(i)
	fmt.Printf("%v %T\n", j, j)

	// To Convert Int to String, use strconv package. string(i) is finding the unicode character for int i
	k := string(i)
	fmt.Printf("%v %T \n", k, k)

	l := strconv.Itoa(i)
	fmt.Printf("%v %T \n", l, l)

}

func stringsAndRunes() {
	str := "Hello"
	fmt.Printf("%v %T\n", str, str)

	// int32 or ascii character stored in rune. Can use ReadRune and ReadByte for getting character out.
	// rune is defined with single quote and string is defined with double quote
	var r rune = 'a'
	rune1 := 'A'
	fmt.Printf("%v %T\n", r, r)
	fmt.Printf("%v %T\n", rune1, rune1)
}

func constants() {
	// Block level
	const myConst float64 = 2.58
	// Package level
	const MyConst float64 = 2.58

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	// Constant can not be something which needs to be derived. Means you can not assign value from function return to a constant
	// const myConst float64 = math.Sin(1.57)

}

func arrays() {
	grades := [...]int{22, 33, 45}
	// c is a copy of grades
	c := grades
	fmt.Printf("c: %v\n", c)
	fmt.Printf("grades: %v\n", grades)
	// d is ref of grades
	d := &grades
	fmt.Printf("d: %v\n", d)
	fmt.Printf("grades: %v\n", grades)
}

// Slice is a projection of an array/slice
func slices() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// In Slice, `c` is a ref of `a`
	c := a
	c[1] = 5
	fmt.Printf("a: %v\n", a)
	fmt.Printf("c: %v\n", c)

	bigData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Works with both slice and array
	// bigData := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := bigData[:]  // Slice of all elements
	s2 := bigData[3:] // Slice from 4th element to end
	s3 := bigData[:6] // Slice first 6 elements
	// left side index is inclusive and right side index is exclusive
	s4 := bigData[3:6] // Slice the 4th, 5th and 6th elements

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)

	makeSlice := make([]int, 3, 100)
	fmt.Printf("makeSlice: %v\n", makeSlice)
	fmt.Printf("len: %v\n", len(makeSlice))
	fmt.Printf("cap: %v\n", cap(makeSlice))

	// Spread operator
	makeSlice = append(makeSlice, []int{1, 2, 3, 4}...)
	fmt.Printf("makeSlice: %v\n", makeSlice)

	// Shift operation -> remove element from start
	makeSlice = makeSlice[1:]
	fmt.Printf("makeSlice: %v\n", makeSlice)

	// Pop Operation -> Remove element from last
	makeSlice = makeSlice[:len(makeSlice)-1]
	fmt.Printf("makeSlice: %v\n", makeSlice)

	// Remove element from middle. remove index 2
	makeSlice = append(makeSlice[:2], makeSlice[3:]...)
	fmt.Printf("makeSlice: %v\n", makeSlice)
}

func maps() {
	mapEx := map[string]int{
		"Record1": 235,
		"Record2": 2225,
		"Record3": 18556,
		"Record4": 8764,
	}
	fmt.Printf("mapEx: %v\n", mapEx)
	fmt.Printf("mapEx[\"Record1\"]: %v\n", mapEx["Record1"])

	delete(mapEx, "Record1")
	fmt.Printf("mapEx: %v\n", mapEx)

	record1, ok := mapEx["Record1"]
	fmt.Println(record1, ok)

	// cp is reference of mapEx
	cp := mapEx
	fmt.Printf("cp: %v\n", cp)
}

type User struct {
	name string
	age  int
}

type Animal struct {
	// third property is called tags. which is not useful for golang. but mostly used for validation or other side effect libs.
	name   string `required max:"100"`
	origin string
}

// This is called composition / embeddings where we are combining two structs into one
type Bird struct {
	Animal
	canFly bool
	speed  int
}

func structs() {
	aUser := User{
		name: "Nisarg",
		age:  10,
	}

	// bUser is a copy of aUser
	bUser := aUser
	// bUser is a copy of aUser
	cUser := &aUser
	bUser.age = 23
	fmt.Printf("aUser: %v\n", aUser)
	fmt.Printf("bUser: %v\n", bUser)
	fmt.Printf("cUser: %v\n", cUser)

	oneBird := Bird{
		Animal: Animal{
			name:   "pe",
			origin: "awda",
		},
		canFly: true,
		speed:  24,
	}
	fmt.Printf("oneBird: %v\n", oneBird)
	twoBird := Bird{}
	twoBird.name = "awd"
	twoBird.origin = "awdaa"
	twoBird.canFly = false
	fmt.Printf("twoBird: %v\n", twoBird)

	tagsOfAnimal := reflect.TypeOf(Animal{})
	field, _ := tagsOfAnimal.FieldByName("name")
	fmt.Printf("Tag: %v\n", field.Tag)
}

func condition() {
	if true {
		fmt.Println("This is true")
	}

	mapEx := map[string]int{
		"Record1": 12,
	}

	if record1, ok := mapEx["Record1"]; ok {
		fmt.Printf("record1: %v\n", record1)

	}

}

func switches() {
	switch 4 {
	case 1, 5, 10:
		fmt.Println("1,5,10")
	default:
		fmt.Println("default")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or eq 10")
		// It will execute next case even if is not satisfied. It will ignore condition.
		fallthrough
	case i <= 20:
		fmt.Println("less than or eq 20")
	default:
		fmt.Println("default")
	}
}

func main() {

	// variables()
	// variableConversion()
	// stringsAndRunes()
	// constants()
	// arrays()
	// slices()
	// maps()
	// structs()
	// condition()
	switches()
}
