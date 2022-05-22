package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	name = strings.TrimSpace(name)

	b := newBill(name)

	fmt.Println("Created bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you choose to save to the bill", b.name)
	default:
		fmt.Println("invalid option, try again ...")
		promptOptions(b)
	}
}

func mainBill() {
	mybill := createBill()
	promptOptions(mybill)
}
