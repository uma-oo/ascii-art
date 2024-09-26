package main

import (
	"log"
	"os"
	"strings"

	ascii "ascii/functions"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("There is no arguments! Please provide a one!")
	}

	file, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	liste_of_letters := strings.Split(string(file[1:len(file)-1]), "\n\n")

	m := ascii.CreateMap(liste_of_letters)
	words := ascii.SplitNewLine(args[0])

	ascii.Print(words, m)
}
