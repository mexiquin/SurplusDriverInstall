package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const networkFlag int = 1000
const dcuFlag int = 1001

var targetInstall int

func main() {
	installChoice, err := strconv.Atoi(os.Args[1])
	if err == nil {
		switch installChoice {
		case networkFlag:
			networkInstall()
		case dcuFlag:
			driverInstall()
		}
	} else {
		fmt.Println("Error while inputting args")
	}

}

// findExecutible will search current working
// directory for the desired exe file and return its name as a string
func findExecutable(directory string, keywords ...string) (string, bool) {
	// get all files into one variable
	allfiles := GetAllFiles(directory)

	// check through all files for matching keywords
	for _, item := range allfiles {
		isCompleteMatch, _ := checkSubstrings(item, keywords)

		if isCompleteMatch == true {
			return item, true
		}
	}

	return "Item not found", false
}

func setTarget(target int) bool {
	if target == dcuFlag || target == networkFlag {
		targetInstall = target
		return true
	}

	return false
}

// getwd gets current working directory
func getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
		return "Nope, couldn't get current working directory"
	}

	return dir
}

// GetAllFiles returns all files within a directory as string in a slice
func GetAllFiles(directory string) []string {

	var listOffiles []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		print("OOPS, something went wrong with ReadDir()")
	}

	for _, item := range files {
		listOffiles = append(listOffiles, item.Name())
	}

	return listOffiles
}

/*
checkSubstrings goes through to see if your chosen string contains all
substrings. returns true if contains all substrings. false otherwise
*/
func checkSubstrings(str string, subs []string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches++
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}

func networkInstall() {
	// get the name of the executable (SDITool)
	sdiExe, isFound := findExecutable(getwd()+"/Scripts", "SDI", "x64")

	// execute the installer
	if isFound {
		cmd := exec.Command(sdiExe, "-autoinstall", "-nogui", "-showconsole", "-autoclose")

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to execute %s\n", sdiExe)
		}
	}

}

// driverInstall silently installs the DCU program
func driverInstall() {
	// get the name of the executable (DCU)
	dciExe, isFound := findExecutable(getwd()+"/Scripts", "DCU")

	// execute the installer
	if isFound {
		cmd := exec.Command(dciExe, "/s")

		err := cmd.Run()

		if err != nil {
			fmt.Printf("Failed to execute %s\n", dciExe)
		}
	}
}
