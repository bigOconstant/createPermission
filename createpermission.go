package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

		os.Exit(1)
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
				fmt.Println("You choose :", list[securityRoleId].Name)
			}
		}

		var list2 = api.GetSecurityActivityEnumMap(conn)

		fmt.Println("Enter a id corresponding to a section this permession should go into. a number value from below please")

		var SecurityActivityMap map[int]string
		SecurityActivityMap = make(map[int]string)
		SecurityActivityMap[1] = "1000-1999 reserved for configuration/settings data"
		SecurityActivityMap[2] = "3000-3999 reserved for patient related data"
		SecurityActivityMap[3] = "5000-5999 reserved for API OAuth Application permissions"
		SecurityActivityMap[4] = "6000-6999 reserved for media"
		SecurityActivityMap[5] = "7000-7999 reserved for authorization mode"

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

		var newval = returnSecurityActivityNumber(inputid, list2)
		fmt.Println("Value inserting: ", newval)

		fmt.Println("Please choose a level of security below, low is fine in most cases")
		var SecurityLevel map[int]string
		SecurityLevel = make(map[int]string)
		SecurityLevel[1] = "Low"
		SecurityLevel[2] = "Medium - Vivify Support"
		SecurityLevel[3] = "High - Only developers"

		for i := 1; i < len(SecurityLevel)+1; i++ {
			fmt.Println("ID: ", i, " Level: ", SecurityLevel[i])
		}

		var level int

		var securityLevelChosen = 1034
		inputstring = ""
		for inputstring == "" {
			fmt.Scan(&level)
			inputstring = SecurityLevel[level]

			if inputstring == "" {
				fmt.Println("Could not find an id with that input try again")
			} else {
				fmt.Println("You chose :", SecurityLevel[level])
				if level == 1 {
					securityLevelChosen = 1034
				} else if level == 2 {
					securityLevelChosen = 1036
				} else {
					securityLevelChosen = 1038
				}
			}

		}

		fmt.Print("\nPlease enter your new permission  name with no spaces:")

		fmt.Scan(&inputstring)

		fmt.Println("\nName is: ", inputstring)

		reader := bufio.NewReader(os.Stdin)
		desctiptionnew, _ := reader.ReadString('\n')
		fmt.Print("Please Enter a description: ")
		desctiptionnew, _ = reader.ReadString('\n')

		CreateMigrateScript(list[securityRoleId].Name, newval, inputstring, strings.TrimSpace(desctiptionnew), securityLevelChosen)

		fmt.Println("********************Next Steps**********************")
		fmt.Println("Add the following line to Database/Data/dbo.SecurityActivifyEnum.Data.sql")

		fmt.Println("****************************************************")
		fmt.Printf("(%d,'%s',N'%s',%d)\n", newval, inputstring, strings.TrimSpace(desctiptionnew), securityLevelChosen)
		fmt.Println("****************************************************")
		fmt.Println("Add the following line to Database/Data/dbo.SeuciryActiityRoleREL.Data.sql")
		fmt.Println("****************************************************")
		fmt.Printf("(%d,%d)\n", newval, securityRoleId)
		fmt.Println("****************************************************")
		fmt.Println("Add the following line to Vivify.Platform/Components/Security/SecurityActivityEnum.cs")
		fmt.Println("****************************************************")
		fmt.Printf("%s = %d\n", inputstring, newval)
		fmt.Println("****************************************************")

		fmt.Printf("Ending Application\n")

	}
	defer conn.Close()

}

func CreateMigrateScript(SecurityRole string, id int, name string, description string, securityLevelChosen int) {

	output := fmt.Sprintf("IF NOT EXISTS (SELECT 1 FROM SecurityActivityEnum Where SecurityActivityId =  %d )\n    Begin\n", id)
	output = output + fmt.Sprintf("        INSERT INTO SecurityActivityEnum(SecurityActivityId, Name, Description, FilterSecurityActivityId)\n")
	output = output + fmt.Sprintf("        VALUES ( %d ,'%s' ,'%s',%d  )\n", id, name, description, securityLevelChosen)
	output = output + fmt.Sprintf("    End\n")
	output = output + fmt.Sprintf("IF NOT EXISTS (SELECT 1 FROM SecurityActivityRoleRel WHERE SecurityActivityId = %d AND SecurityRoleId = (SELECT SecurityRoleId FROM SecurityRole WHERE Name = '%s'))\n", id, SecurityRole)
	output = output + fmt.Sprintf("BEGIN\n")
	output = output + fmt.Sprintf("    INSERT INTO SecurityActivityRoleREL VALUES (%d, ( SELECT SecurityRoleId FROM SecurityRole WHERE Name = '%s'))\n", id, SecurityRole)
	output = output + fmt.Sprintf("END\n")

	filename := fmt.Sprintf("IHCP-____Add%sPermission.sql", name)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Something bad happened", err)
	} else {
		fmt.Println("Writting migration script to ", filename)
	}
	defer file.Close()

	fmt.Fprintf(file, output)

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

func getConnection(filename string) models.Connection {
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("No connection file found")
		CreateConnectionFile()
		return getConnection(filename)

		//os.Exit(1)
	}

	var c models.Connection

	json.Unmarshal(raw, &c)

	return c
}

func CreateConnectionFile() {

	conn := models.Connection{}

	fmt.Println("Hi, it looks like you don't have a connection file.\nLets go ahead and create one")
	fmt.Println("Please enter the server address, if its your localhost make sure to enable tcp connections")
	fmt.Print("Server: ")
	fmt.Scan(&conn.Server)
	fmt.Print("\nPlease Enter the DB User name \nUser: ")
	fmt.Scan(&conn.User)
	fmt.Println("\nPlease Enter the Database Name example, develop")
	fmt.Print("Database: ")
	fmt.Scan(&conn.Database)
	fmt.Println("\nPlease Enter the Password for the database")
	fmt.Print("Password: ")
	fmt.Scan(&conn.Password)

	JsonFile, _ := json.Marshal(conn)
	ioutil.WriteFile("connection.json", JsonFile, 0644)

}
