package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")
	
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found")
	}
	fmt.Println("PORT:", portString)

}
