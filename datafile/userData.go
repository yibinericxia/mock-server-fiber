package datafile

import (
	"encoding/csv"
	"fmt"
	"io"
	"mock-server-fiber/model"
	"os"
	"strconv"
)

var users []model.User

func parseCSV4UserData() error {
	file, err := os.Open("datafile/user.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return err
		}
		user := model.User{
			ID:    id,
			Name:  row[1],
			Email: row[2],
		}
		users = append(users, user)
	}
	fmt.Println(users)
	return nil
}

func GetUserData() []model.User {
	if len(users) == 0 {
		if err := parseCSV4UserData(); err != nil {
			fmt.Println(err)
		}
	}
	return users
}
