package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./api"
	"./models"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	fmt.Println("Connection To DB and Loading Data...")

	connObj := getConnection("./connection.json")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", connObj.Server, connObj.User, connObj.Password, connObj.Database)

	fmt.Println(connString)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	} else {
		fmt.Println("Connection Sucessfully")
		fmt.Println("*********Generate Perimission Script***************")

		fmt.Println("Printing out security list")

		var list = api.GetSecurityRoles(conn)

		printList(list)
		var securityRoleId int
		fmt.Println("Which SecurtityRole should your permission fall under? Please Enter an ID from above")

		var inputstring = ""

		for inputstring == "" {
			fmt.Scan(&securityRoleId)
			if err != nil {
				fmt.Println("Problem here", nil)
			}
			inputstring = list[securityRoleId].Name
			if inputstring == "" {
				fmt.Println("Could not find an id with that input try again")
			} else {
				fmt.Println("You chose :", list[securityRoleId].Name)
			}
		}

		var list2 = api.GetSecurityActivityEnumMap(conn)

		fmt.Println("Enter a id corresponding to a section this permession should go into. a number value from above please")

		var SecurityActivityMap map[int]string
		SecurityActivityMap = make(map[int]string)
		SecurityActivityMap[1] = "1000-1999 reserved for configuration/settings data"
		SecurityActivityMap[2] = "3000-3999 - reserved for patient related data"
		SecurityActivityMap[3] = "5000-5999 reserved for API OAuth Application permissions"
		SecurityActivityMap[4] = "6000-6999 reserved for media"
		SecurityActivityMap[5] = "7000 - 7999 reserved for authorization mode"

		for i := 1; i < len(SecurityActivityMap)+1; i++ {
			fmt.Println("ID: ", i, "Description:", SecurityActivityMap[i])

		}

		inputstring = ""
		var inputid int
		for inputstring == "" {
			fmt.Scan(&inputid)
			if err != nil {
				fmt.Println("Problem here", nil)
			}
			inputstring = SecurityActivityMap[inputid]
			if inputstring == "" {
				fmt.Println("Could not find an id with that input try again")
			} else {
				fmt.Println("You chose :", SecurityActivityMap[inputid])
			}
		}

		fmt.Println("Testing new function")
		var newval = returnSecurityActivityNumber(inputid, list2)
		fmt.Println("Valinserting: ", newval)

		fmt.Println("\n Please enter your new permission  name with no spaces")

		fmt.Printf("Ending Application\n")

	}
	defer conn.Close()

}

func returnSecurityActivityNumber(section int, input map[int]models.SecurityActivityEnum) int {

	var counter = 0

	switch section {
	case 1:
		counter = 1000
	case 2:
		counter = 3000
	case 3:
		counter = 5000
	case 4:
		counter = 6000
	case 5:
		counter = 7000
	default:
		counter = 1000

	}
	var returnval = 0

	for begin := counter; begin < counter+1000; begin++ {
		if input[begin].Name == "" {
			returnval = begin
			break
		}
	}

	return returnval

}

func printList(inputL map[int]models.Security) {

	for i := 1; i < len(inputL)+1; i++ {
		if inputL[i].Id > 9 {
			fmt.Println("Id:", inputL[i].Id, " Name:", inputL[i].Name)
		} else {
			fmt.Println("Id:", inputL[i].Id, "  Name:", inputL[i].Name)
		}
	}

}

func printOtherList(inputL map[int]models.SecurityActivityEnum) {
	for i := 0; i < 1; i++ {
		fmt.Println("SID: ", inputL[i].SecurityActivityId, " Name: ", inputL[i].Name, " Description: ", inputL[i].Description, " FilterSecurityActivityId: ", inputL[i].FilterSecurityActivityId)

	}

}

func getConnection(filename string) models.Connection {
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c models.Connection

	json.Unmarshal(raw, &c)

	return c
}
