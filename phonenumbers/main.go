package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
	"unicode"
)

func main() {
	input := `1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892`

	db, err := sql.Open("sqlite3", "phonenumbers/db.sqlite")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stmt, _ := db.Prepare("INSERT INTO phone_numbers(phone_number_original, phone_number_normalized) values(?,?)")

	results := make(chan string)
	numbers := 0
	for _, pn := range strings.Split(input, "\n") {
		numbers++
		go func(phoneNumber string){
			normalizedPhoneNumber := ""
			for _, c := range phoneNumber {
				if unicode.IsDigit(c) {
					normalizedPhoneNumber += string(c)
				}
			}
			res, err := stmt.Exec( strings.TrimSpace(phoneNumber), normalizedPhoneNumber)
			if err != nil {
				results <- fmt.Sprintf("Issue adding %s to database: %s\n", phoneNumber, err)
			} else {
				id, _ := res.LastInsertId()
				results <- fmt.Sprintf("Inserted %s with uid %d\n", phoneNumber, id)
			}
		}(pn)
	}
	for numbers > 0 {
		fmt.Println(<- results)
		numbers--
	}

	rows, err := db.Query("SELECT uid, phone_number_normalized FROM phone_numbers")
	var pn string
	var uid int
	for rows.Next() {
		err := rows.Scan(&uid, &pn)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("#%d: %s\n", uid, pn)
		}
	}

}
