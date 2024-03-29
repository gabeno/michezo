package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gabeno/poker/v1"
)

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type: '{name} wins' to record a win")

	const dbFileName = "game.db.json"

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}

	game := poker.NewCli(store, os.Stdin)
	game.PlayPoker()
}
