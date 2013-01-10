package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var dirCurrent, _ = os.Getwd() // Directory that the binary is in
var dirResources = filepath.Join(dirCurrent, "resources")

func main() {
	var choice int

	Clr()
	fmt.Println("Welcome to Grognak's Mod Patcher!")

	if !CheckSafety() {
		Pause(true)
	}

	for {
		choice = MainMenu()
		if choice >= 1 && choice <= 4 {
			break
		} else {
			Clr() // and then show the main menu again
		}
	}

	switch choice {
	case 1:
		// Patch all mods
		StartPatch()
	case 2:
		// Restore unmodded game
		RestoreBackups()
	case 3:
		// Create updated backups
		UpdateBackups()
	case 4:
		// Exit
		os.Exit(0)
	}

	log.Println("Operation completed successfully!")
	Pause(true) // End the program
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
		fmt.Println("Sorry, but there was an error (#1):\n You need to install this binary in it's correct location.")
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
	fmt.Println("4) Exit\n")
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
		os.Exit(0)
	} else {
		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln(&s)
		return
	}
}

func StartPatch() {
	// TODO
}

func RestoreBackups() {
	log.Println("Restoring data file backups...")
	os.Remove(filepath.Join(dirResources, "data.dat"))
	os.Remove(filepath.Join(dirResources, "resources.dat"))
	_, err := CopyFile(filepath.Join(dirResources, "data.dat.bak"), filepath.Join(dirResources, "data.dat"))
	_, err = CopyFile(filepath.Join(dirResources, "resource.dat.bak"), filepath.Join(dirResources, "resource.dat"))
	if err != nil {
		log.Println("There was an error restoring backups. Are you sure that")
		log.Println("you have patched at least once?")
		log.Fatal(err)
	}
}

func UpdateBackups() {
	// TODO
}

func Clr() {
	// TODO: Refine
	for i := 0; i < 32; i++ {
		fmt.Println("")
	}
}

func CopyFile(src, dest string) (written int64, err error) {
	sf, err := os.Open(src)

	if err != nil {
		return 0, err
	}

	defer sf.Close()
	df, err := os.Create(dest)

	if err != nil {
		return 0, err
	}

	defer df.Close()
	return io.Copy(df, sf)
}
