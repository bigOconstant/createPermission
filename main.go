package main

import (
	"fmt"

	"./CreateConfiguration"
	"./CreatePermission"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	m := make(map[int]string)
	m[1] = "Create Permission"
	m[2] = "Create Configuration"

	var PickPath = ""
	var exit = true
	var choice = 0
	for exit {
		fmt.Println("\n\n     Permission       ( ͡° ͜ʖ ͡°)     Configuration     ")
		fmt.Println("\n\n                       Script                        \n\n")
		fmt.Println("Enter 1 to Create a Permission")
		fmt.Println("Enter 2 to Create a Configuration  (Beta, Not finished)")
		fmt.Println("Enter q to quit")
		fmt.Scan(&PickPath)

		if PickPath == "q" {
			choice = 3
			exit = false

		} else if PickPath == "1" {
			fmt.Println("You entered:", PickPath)
			choice = 1
			exit = false
		} else if PickPath == "2" {
			fmt.Println("You entered:", PickPath)
			choice = 2
			exit = false
		} else {
			fmt.Println("You entered Something wrong")
		}

	}

	if choice == 1 {
		CreatePermission.CreatePermission()
	} else if choice == 2 {
		CreateConfiguration.CreateConfiguration()
	}

	fmt.Println("\n＼( ･_･) GoodBye\n\n")

}
