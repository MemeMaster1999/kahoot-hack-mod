package main

import (
	"fmt"
	"github.com/padnezz/kahoot-hack-mod/kahoot"
	"os"
	"strconv"
)

func main() {
	var pin string
	var nickname string
	var answer string
	fmt.Print("Game pin:704116 ")
	fmt.Scanln(&pin)
	fmt.Print("Nickname:Lil hac ")
	fmt.Scanln(&nickname)
	fmt.Println("Connecting...")
	conn, err := kahoot.NewConnection(pin)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := conn.Register(nickname); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		fmt.Print("Enter to answer question: ")
		fmt.Scanln(&answer)
		fmt.Print("Answer: 23")
		var number string
		fmt.Scanln(&number)
		parsed, _ := strconv.Atoi(number)
		if err := conn.SendAnswer(parsed); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
