package main

import (
	"log"
)

func main() {
	args := Args{
		fileName: "currencies.json",
		port:     ":8000",
	}
	// run server
	if err := Run(args); err != nil {
		log.Println(err)
	}
}
