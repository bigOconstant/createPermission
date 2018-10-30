package CreateConfiguration

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"../api"
	"../models"
	_ "github.com/denisenkom/go-mssqldb"
)

func CreateConfiguration() {
	fmt.Println("Connection To DB and Loading Data...")
	reader := bufio.NewReader(os.Stdin)
	connObj := api.GetConnection("./connection.json")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", connObj.Server, connObj.User, connObj.Password, connObj.Database)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())

		os.Exit(1)
	} else {
		fmt.Println("Connection Sucessfully")
		fmt.Println("*********Generate Configuration Script***************")
		var list = api.GetConfigurationSections(conn)

		println("Press Enter to display a list of catigories and associated ids")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		printConfigurations(list)
		print("\nType the id of the catigory you would like to use or type -1 to create a new category:")

		//var ConfigurationId int

		var inputstring = ""
		//var category = 0
		var categoryName = ""
		var categorynew = false
		for inputstring == "" {
			fmt.Scan(&inputstring)
			if err != nil {
				fmt.Println("Problem here", nil)
			}

			if inputstring == "-1" {
				fmt.Print("Enter your new Catagory Name\n")
				fmt.Scan(&inputstring)
				fmt.Println("\nNew name:", inputstring)
				categoryName = inputstring
				categorynew = true
			} else {
				nameId, err := strconv.Atoi(inputstring)

				if err != nil || list[nameId] == "" {

					inputstring = ""
					fmt.Print("\nHmm value not found enter another one:")

				} else {
					fmt.Println("erro value below")
					fmt.Println(err)
					fmt.Println("number you chose:", nameId)
					fmt.Println("You choose category:", list[nameId])
					categoryName = list[nameId]
				}
			}
		}
		if categorynew {
			fmt.Println("New Category:", categoryName)
		} else {

			var ExistingConfigs = api.GetConfigurationsByName(categoryName, conn)
			println("Current params in the " + categoryName + " Category")
			printConfigurationList(ExistingConfigs)

		}
		print("\n\n Please Enter a new Param name:")
		var newparamName = ""
		var newDescription = ""
		fmt.Scan(&newparamName)
		fmt.Println("\n You chose:" + newparamName + "\n")
		print("\n\n Please Enter a new Param Description:")
		newDescription, _ = reader.ReadString('\n')
		newDescription, _ = reader.ReadString('\n')
		println("\n Description:" + newDescription)

		var enums = api.GetConfigurationDataType_Enums(conn)
		fmt.Println("Pick a type from the list Below")
		printEnumList(enums)

		fmt.Print("Enter Id:")

		var idType = ""
		var TypeLabel = ""
		var IdOfType = 0
		for TypeLabel == "" {
			fmt.Scan(&idType)
			numberIdType, err := strconv.Atoi(idType)
			if err == nil {
				//fmt.Println("Enum Object Below")
				//fmt.Println(enums[numberIdType])
				TypeLabel = enums[numberIdType].Name
				fmt.Println(TypeLabel)
				if TypeLabel == "" {
					fmt.Println("You Entered something incorrect")
					fmt.Println("Pick a type from the list Below")
					printEnumList(enums)
				} else {
					IdOfType = numberIdType
				}
			}

		}
		fmt.Println("You choose the object below")
		fmt.Println(enums[IdOfType])

		fmt.Println("Would you like to add valid values? Type y for yes or n for no")

		var yesOrNo = ""

		var ValueList []string
		reader := bufio.NewReader(os.Stdin)
		fmt.Scan(&yesOrNo)

		if strings.ToLower(yesOrNo) == "y" {
			fmt.Print("Enter a Valid Value:")
			validValues, _ := reader.ReadString('\n')
			validValues, _ = reader.ReadString('\n')
			ValueList = append(ValueList, validValues)
			var enterAnotherValue = ""
			fmt.Print("Would you like to Enter another Valid value? :")
			for strings.ToLower(enterAnotherValue) != "n" {
				fmt.Scan(&enterAnotherValue)
				fmt.Println("")
				fmt.Print("Enter a Valid Value:")
				validValues, _ := reader.ReadString('\n')
				validValues, _ = reader.ReadString('\n')
				ValueList = append(ValueList, validValues)
				fmt.Println("\n")
				fmt.Print("Would you like to Enter another Valid value:")
				fmt.Scan(&enterAnotherValue)

			}
		}
		fmt.Println("length of valid values,:", len(ValueList))
		fmt.Print("\nWould you like to enter a default Value? Y for yes N for no: ")
		var defaultVal = ""

		fmt.Scan(&defaultVal)
		if strings.ToLower(defaultVal) == "y" {
			fmt.Println("")
			fmt.Print("Enter Value:")
			defaultVal, _ := reader.ReadString('\n')
			defaultVal, _ = reader.ReadString('\n')
			fmt.Println("")
			fmt.Println("You Choose:" + defaultVal)
		}
	}

}

func printConfigurations(input map[int]string) {
	for i := 1; i < len(input)+1; i++ {
		if i < 10 {
			fmt.Println("Id :", i, " Name:", input[i])
		} else {
			fmt.Println("Id:", i, " Name:", input[i])
		}

	}
}
func printConfigurationList(input map[int]models.Configuration) {

	for key := range input {
		fmt.Printf("\nParam: " + input[key].Param)
	}
}
func printEnumList(input map[int]models.ConfigurationDataType_Enum) {

	var sortedListOfEnums []models.ConfigurationDataType_Enum

	for key := range input { // Put map in a list for easy sorting
		sortedListOfEnums = append(sortedListOfEnums, input[key])
	}
	sort.Slice(sortedListOfEnums, func(i, j int) bool { // Sort said list
		return sortedListOfEnums[i].ConfigurationDataTypeId < sortedListOfEnums[j].ConfigurationDataTypeId
	})
	for _, v := range sortedListOfEnums { // Print and Profit
		fmt.Println("Id: ", v.ConfigurationDataTypeId, " Param: "+v.Name)
	}

}
