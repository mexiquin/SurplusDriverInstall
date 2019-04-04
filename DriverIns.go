package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

var scriptsDir = getwd() + "/Scripts/"
var versionNum float32 = 1.0

// Provides the user interface and root of method execution
func main() {

	// Create user interface
	quintons := figure.NewFigure("Quinton's", "", true)
	driver := figure.NewFigure("Driver", "", true)
	instslr := figure.NewFigure("Installer", "", true)

	quintons.Print()
	driver.Print()
	instslr.Print()

	fmt.Printf("\n%52s: %2.1f\n\n", "Version", versionNum)

	fmt.Printf("(ENTER) Install All\n(1)Install Network\n(2)Dell Command Update\n\n")
	var i int
	_, err := fmt.Scanf("%d", &i)

	os.Chdir(scriptsDir)

	if err != nil {
		networkInstall()
		driverInstall()
	} else if i == 1 {
		networkInstall()
	} else if i == 2 {
		driverInstall()
	} else {
		fmt.Println("Invalid argument")
		time.Sleep(time.Second * 5)
	}

	// Eventually want to implement: copy and paste MediaCreationTool to the desktop
}

// findExecutable will search current working
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
		time.Sleep(time.Second * 5)
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
	sdiExe, isFound := findExecutable(getwd(), "SDI", "x64")

	// execute the installer
	if isFound {
		cmd := exec.Command(sdiExe, "-autoinstall", "-nogui", "-showconsole", "-autoclose")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	} else {
		fmt.Println("Error: Could not find executable")
		fmt.Println(sdiExe)
		time.Sleep(time.Second * 5)
	}

}

// driverInstall silently installs the DCU program
func driverInstall() {
	// get the name of the executable (DCU)
	dciExe, isFound := findExecutable(getwd(), "DCU")

	// execute the installer
	if isFound {
		cmd := exec.Command(dciExe, "/s")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	} else {
		fmt.Println("Error: Could not find executable")
		time.Sleep(time.Second * 5)
	}
}

// MoveFile to copy file to new dir
func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	return nil
}
