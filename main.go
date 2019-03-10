package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Irainia/tdd_golang/human"
)

func main() {
	person := human.NewPerson()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("guess: ")
		inputAgeString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		inputAgeString = strings.Replace(inputAgeString, "\n", "", -1)

		inputAgeInt, err := strconv.Atoi(inputAgeString)
		if err != nil {
			log.Fatalln(err)
		}

		err = person.GuessAge(inputAgeInt)
		if err == nil {
			fmt.Println("correct")
			break
		}

		fmt.Println(err.Error())
	}
}
