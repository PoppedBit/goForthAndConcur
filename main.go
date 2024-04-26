package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Pet struct {
	name string
	age  int
}

type Cat struct {
	hasClaws bool
}

func main() {

	// ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Arrays
	myList := []int{1, 2, 3, 4, 5}

	// For Loop
	for i := 0; i < len(myList); i++ {
		// If, else if, else
		if myList[i] < 3 {

		} else if myList[i] < 5 {
			continue
		} else {
			break
		}
	}

	// While Loops
	n := 1
	for n < 5 {
		n += 1
	}

	// For Each
	for i, s := range []string{"Hello", "World", "!"} {
		fmt.Println(i, s)
	}

	// HTTP
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":"+port, nil)
}
