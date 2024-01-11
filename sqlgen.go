package main

import "fmt"

func GeneratePuzzleSQL() {
	fmt.Println("Hello from sql")
}

func createTables() {
	puzzleTable := `
        CREATE TABLE IF NOT EXISTS puzzles (
            id SERIAL PRIMARY KEY,
            puzzle_id VARCHAR(5) NOT NULL,
            fen VARCHAR(100) NOT NULL,
            moves VARCHAR(500) NOT NULL,
            rating SMALLINT NOT NULL,
            ratingDeviation SMALLINT NOT NULL,
            popularity SMALLINT NOT NULL,
            nb_plays SMALLINT NOT NULL,
            themes VARCHAR(150) NOT NULL,
            game_url VARCHAR(150),
            opening_tags TEXT
        )`
	fmt.Println(puzzleTable)
}
