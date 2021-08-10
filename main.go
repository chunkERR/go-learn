package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	whatTosay := doctor.Intro()

	fmt.Println(whatTosay)

	for {

		userInput, _ := reader.ReadString('\n')

		fmt.Println(userInput)
	}
}
