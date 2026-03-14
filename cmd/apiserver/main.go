package main

import (
	"go_api/internal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
