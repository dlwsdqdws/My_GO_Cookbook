package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	// avoid same value every time
	rand.Seed(time.Now().UnixNano())

	setNum := rand.Intn(maxNum)

	fmt.Println("Please Enter A Guess Value")

	// use stream to deal with I/O
	reader := bufio.NewReader(os.Stdin)
	for {
		// read by line
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Unable read your input, please try again", err)
			continue
		}
		// delete \n
		input = strings.Trim(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please try again", err)
			continue
		}

		if guess > setNum {
			fmt.Println("your guess is bigger, please try again")
		} else if guess < setNum {
			fmt.Println("your guess is smaller, please try again")
		} else {
			fmt.Println("Correct!")
			break
		}
	}
}
