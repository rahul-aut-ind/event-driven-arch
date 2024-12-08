package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_URI       = "127.0.0.1:3306"
	DB_USER      = "root"
	DB_USER_PSWD = "root12345"
	DB_NAME      = "employees"
)

type (
	Employee struct {
		emp_no     int    `json:"emp_no"`
		birth_date string `json:"birth_date"`
		first_name string `json:"first_name"`
		last_name  string `json:"last_name"`
		gender     string `json:"gender"`
		hire_date  string `json:"hire_date"`
	}
)

func main() {

	db := initDB()

	getFirstEmployee(db)
	getEmployeeNo(db, 10002)

	cleanup(db)

}

func getEmployeeNo(db *sql.DB, empNo int) {
	// defer db.Close()
	// Execute the query
	//res, err := db.Query("SELECT first_name from employees where emp_no=" + strconv.Itoa(empNo))
	res, err := db.Query("SELECT * FROM employees where emp_no=" + strconv.Itoa(empNo))
	if err != nil {
		log.Fatalf("Unable to get data from DB. Err %v", err.Error())
	}

	// Fetch data
	for res.Next() {
		var emp Employee
		err = res.Scan(&emp.emp_no, &emp.birth_date, &emp.first_name, &emp.last_name, &emp.gender, &emp.hire_date)
		if err != nil {
			log.Fatalf("Unable to parse data from DB response. Err %v", err.Error())
		}
		fmt.Printf("%#v\n", emp)
		fmt.Println("-----------------------------------")
	}

}

func getFirstEmployee(db *sql.DB) {
	// defer db.Close()
	// Execute the query
	rows, err := db.Query("SELECT * FROM employees limit 1")
	if err != nil {
		log.Fatalf("Unable to get data from DB. Err %v", err.Error())
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Unable to get columns from rows. Err %v", err.Error())
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatalf("Unable to parse data from DB response. Err %v", err.Error())
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("error getting data from DB. Err %v", err.Error())
	}

}

func initDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, DB_USER_PSWD, DB_URI, DB_NAME)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to DB. Err %v", err.Error())
		return nil
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func cleanup(db *sql.DB) {
	db.Close()
}
