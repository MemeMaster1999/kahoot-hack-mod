package main

import (
	"fmt"
	"github.com/padnezz/kahoot-hack-mod/kahoot"
	"os"
)

func main() {
	var pin string
	var nickname string
	var crash string
	fmt.Print("Spelkod: ")
	fmt.Scanln(&pin)
	fmt.Print("Nickname: ")
	fmt.Scanln(&nickname)
	fmt.Println("Kopplar...")
	conn, err := kahoot.NewConnection(pin)
	if err != nil {
		fmt.Println("Fel:", err)
		os.Exit(1)
	}
	if err := conn.Register(nickname); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Enter to start crash:")
        fmt.Scanln(&crash)
		fmt.Println("CRASHING!")
		conn.SendCrashAnswer()
}
