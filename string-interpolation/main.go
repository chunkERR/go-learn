package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	userName        string
	age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.userName = readString("What is your name?")
	user.age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnsADog = readYesOrNo("Do you own a dog? [Y/N]")

	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.2f. Owns a dog: %t.\n",
		user.userName,
		user.age,
		user.FavouriteNumber,
		user.OwnsADog)
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value.")
		} else {
			return userInput
		}

	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number.")
		} else {
			return num
		}

	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number.")
		} else {
			return num
		}

	}
}

func readYesOrNo(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Pleae type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}

	}
}
