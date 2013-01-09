package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Grognak's Mod Patcher!")

	if !CheckSafety() {
		Pause(true)
	}

	for {
		choice := MainMenu()
		if choice >= 0 && choice <= 3 {
			break
		} else {
			Clr() // and then show the main menu again
		}
	}
}

func CheckSafety() bool {
	var fileToCheck string

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("It appears you are running OS X. Congratulations!\n")
		fileToCheck = "FTL_README.html"
	case "windows":
		fmt.Println("Looks like you're running Windows. Great!\n")
		fileToCheck = "FTLGame.exe"
	case "linux":
		fmt.Println("It seems like you're running Linux. Fantastic!\n")
		fileToCheck = "FTL"
	}

	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		fmt.Println("Sorry, but there was an error. You need to install this binary in it's right place.")
		fmt.Println("Try consulting the included readme for more information.")
		return false
	}

	return true
}

func MainMenu() int {
	var result string
	var resultint int

	fmt.Println("What would you like to do?\n")
	fmt.Println("1) Patch all mods")
	fmt.Println("2) Restore unmodded game")
	fmt.Println("3) Create updated backups")
	fmt.Println("0) Exit\n")
	fmt.Print(">> ")

	_, err := fmt.Scanln(&result)
	if err != nil {
		log.Fatal(err)
	}

	resultint, _ = strconv.Atoi(result)

	return resultint // The result converted into an int
}

func Pause(exiting bool) {
	var s string

	if exiting {
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln(&s)
		os.Exit(1)
	} else {
		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln(&s)
		return
	}
}

func Clr() {
	// TODO: Refine
	for i := 0; i < 64; i++ {
		fmt.Println("")
	}
}
