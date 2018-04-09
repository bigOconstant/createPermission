package api

import (
	"database/sql"
	"fmt"

	"../models"
	"os"
)

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

	if err != nil {
		fmt.Printf("We've found a problem")
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
