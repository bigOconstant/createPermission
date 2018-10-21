package api

import (
	"database/sql"
	"fmt"

	"os"

	"../models"
)

func GetConfigurationDataType_Enums(conn *sql.DB) map[int]models.ConfigurationDataType_Enum {
	stmt, err := conn.Prepare("SELECT ConfigurationDataTypeId, Name FROM ConfigurationDataType_Enum order by ConfigurationDataTypeId")

	if err != nil {
		fmt.Println("failed:", err.Error())
		fmt.Println("Check your connection.json file and that you have tcp connections enabled and try again")
		os.Exit(1)
		return map[int]models.ConfigurationDataType_Enum{}
	}

	defer stmt.Close()
	rows, err := stmt.Query()

	if err != nil {
		fmt.Printf("Something bad happened\n")
	}

	var listmap map[int]models.ConfigurationDataType_Enum
	listmap = make(map[int]models.ConfigurationDataType_Enum)
	val := models.ConfigurationDataType_Enum{}

	for rows.Next() {
		err = rows.Scan(&val.ConfigurationDataTypeId, &val.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		listmap[val.ConfigurationDataTypeId] = val
	}
	return listmap
}

func GetConfigurationSections(conn *sql.DB) map[int]string {
	stmt, err := conn.Prepare("SELECT distinct Name FROM TblConfiguration ")
	if err != nil {
		fmt.Println("failed:", err.Error())
		fmt.Println("Check your connection.json file and that you have tcp connections enabled and try again")
		os.Exit(1)
		return map[int]string{}
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("Something bad happened\n")
	}
	var listmap map[int]string
	listmap = make(map[int]string)
	val := ""
	counter := 1
	for rows.Next() {
		err = rows.Scan(&val)
		if err != nil {
			fmt.Println(err)
			continue
		}
		listmap[counter] = val
		counter++
	}
	return listmap
}
func GetConfigurationsByName(name string, conn *sql.DB) map[int]models.Configuration {
	stmt, err := conn.Prepare("SELECT Id,Name,Description,Param,Value,ConfigurationDataTypeId,SecurityActivityId, ValidValues FROM TblConfiguration where Name = ?")

	if err != nil {
		fmt.Println("failed:", err.Error())
		fmt.Println("Check your connection.json file and that you have tcp connections enabled and try again")
		os.Exit(1)
		return map[int]models.Configuration{}
	}

	defer stmt.Close()
	rows, err := stmt.Query(name)

	if err != nil {
		fmt.Printf("Something bad happened\n")
	}

	var listmap map[int]models.Configuration
	listmap = make(map[int]models.Configuration)
	val := models.Configuration{}

	for rows.Next() {

		err = rows.Scan(&val.Id, &val.Name, &val.Description, &val.Param, &val.Value, &val.ConfigurationDataTypeId, &val.SecurityActivityId, &val.ValidValues)
		if err != nil {
			fmt.Println(err)
			continue
		}

		listmap[val.Id] = val
	}
	return listmap

}

func GetSecurityRoles(conn *sql.DB) map[int]models.Security {
	stmt, err := conn.Prepare("SELECT SecurityRoleId, Name FROM SecurityRole ")

	if err != nil {
		fmt.Println("failed:", err.Error())
		fmt.Println("Check your connection.json file and that you have tcp connections enabled and try again")
		os.Exit(1)
		return map[int]models.Security{}
	}

	defer stmt.Close()
	rows, err := stmt.Query()

	if err != nil {
		fmt.Printf("Something bad happened\n")
	}

	var listmap map[int]models.Security
	listmap = make(map[int]models.Security)
	val := models.Security{}

	for rows.Next() {
		err = rows.Scan(&val.Id, &val.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		listmap[val.Id] = val
	}
	return listmap
}

func GetSecurityActivityEnumMap(conn *sql.DB) map[int]models.SecurityActivityEnum {
	stmt, err := conn.Prepare("select SecurityActivityId,Name, Description, FilterSecurityActivityId from SecurityActivityEnum order by SecurityActivityId ")

	if err != nil {
		fmt.Println("Failed:", err.Error())

		fmt.Println("Check your connection.json file and that you have tcp connections enabled and try again")
		os.Exit(1)
		return map[int]models.SecurityActivityEnum{} //return empty array
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		fmt.Printf("Something did not go correctly with your query")

	}

	var smap map[int]models.SecurityActivityEnum

	smap = make(map[int]models.SecurityActivityEnum)

	val := models.SecurityActivityEnum{}

	for rows.Next() {
		err = rows.Scan(&val.SecurityActivityId, &val.Name, &val.Description, &val.FilterSecurityActivityId)
		if err != nil {
			fmt.Println(err)
			continue

		}
		smap[val.SecurityActivityId] = val

	}
	return smap

}
