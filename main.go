package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// struct Cars  {
// 	color string
// 	velocity float64
// }
func main() {
	name := "lucas"
	age := 26

	fmt.Println("my name is", name, "and my age is", age)
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("my name is %q and my age is %v \n", name, age)

	// var formatted = fmt.Sprintf("my name is %v and my age is %v \n", name, age);
	// fmt.Printf(formatted);

	// var ages [3]int = [3]int{20,25,30}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"lucas", "mario", "joao", "pedro"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices
	var scores = []int{100, 200, 60}
	scores = append(scores, 101)
	scores = append(scores, 102)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	greeting := fmt.Sprintf("hello folks!")
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	agesTwo := []int{3, 7, 5, 1, 2}
	// sort.Ints(agesTwo)
	fmt.Println(agesTwo)

	// sort desc

	sorting(agesTwo)

	fmt.Println(agesTwo)

	index := sort.SearchInts(agesTwo, 3)

	fmt.Println(index)

	namesSlice := []string{"lucas", "mario", "joao", "pedro"}

	sort.Strings(namesSlice)
	fmt.Println(namesSlice)
	fmt.Println(sort.SearchStrings(namesSlice, "lucas"))

	for i := 0; i < len(namesSlice); i++ {
		fmt.Println(namesSlice[i])
	}

	for _, name := range namesSlice {
		fmt.Println(name)
	}

	calc1 := circleArea(10.5)
	fmt.Println("my area is", calc1)

	//pointers

	namePointer := "lucas"
	// fmt.Println("memory addres namePointer is", &namePointer)

	m := &namePointer
	// fmt.Println("memory addres of namePointer", m)
	// fmt.Println("memory addres namePointer is", &m)
	// fmt.Println("value at memory addres", *m)
	fmt.Println(namePointer)
	updateName(m)
	fmt.Println(namePointer)

	// mainBill()
	mainShape()
}

func updateName(x *string) {
	*x = "lucas 2"
}

func circleArea(area float64) float64 {
	return math.Pi * area * area
}
func sorting(numbers []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
}
