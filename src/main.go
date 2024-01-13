package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	db, _ := NewDB()

	db.CreateTable("../sql/puzzles-table.csv")
}
