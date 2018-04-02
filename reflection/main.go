package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/lib/pq"
)

type X struct {
	Column1 string
	Column2 string
}

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/test?sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Openning connection: %v", err))
	}
	defer conn.Close()
	rows, err := conn.Query("select * from data")
	if err != nil {
		panic(fmt.Sprintf("Querying: %v", err))
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		panic(fmt.Sprintf("Reading columns: %v", err))
	}
	var x X
	var zv reflect.Value
	xv := reflect.ValueOf(&x)
	xvv := xv.Elem()
	var saved []interface{}
	for _, col := range cols {
		if f := xvv.FieldByName(strings.Title(col)); f != zv {
			saved = append(saved, f.Addr().Interface())
		}
	}
	if len(saved) != len(cols) {
		panic("Not enough variables to unmarshal into")
	}
	for rows.Next() {
		if err := rows.Scan(saved...); err != nil {
			panic(fmt.Sprintf("Scanning: %v", err))
		}
		fmt.Printf("%#v\n", x)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Sprintf("Iterating: %v", err))
	}
}
