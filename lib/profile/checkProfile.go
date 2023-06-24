package profile

import (
	lib "cryptoTracker/lib"

	"fmt"
	"os"
	"path/filepath"
)

// Check if data directory exists, and if there are profiles in
func checkProfiles(dir string) (ws Websocket) {
	// Check if data directory exists
	_, err := os.Open("data")
	if err != nil {
		if os.IsNotExist(err) {
			// Create data directory
			err := os.Mkdir("data", 0755)
			if err != nil {
				lib.CriticalError(err)
			}
		} else {
			lib.CriticalError(err)
		}
	}

	// Change current working directory to data
	os.Chdir(filepath.Join(dir, "data"))

	answer = lib.BooleanQuestion("Do you want to use an existing profile?")
	// Check existing profiles
	if answer == false {
		ws = CreateProfile(dir)
	} else {
		profiles, err := filepath.Glob("*.json")
		if err != nil {
			lib.CriticalError(err)
		}

		n = len(profiles)

		switch n {
		case 0:
			lib.PrintError("File doesn't exist")
			ws = CreateProfile(dir)
		case 1:
			fmt.Printf("*** Profile %s selected ***\n", profiles[0])
			ws.ReadJsonFile(profiles[0])
		default:
			fmt.Println("*** Profiles ***")
			x = lib.MultipleSelection(profiles, "Which profile do you want to use?")
			ws.ReadJsonFile(profiles[x-1])
		}
	}
	return
}

// Check if the user want to use AutoMode
func ProfileSelection() (ws Websocket) {
	dir, err := os.Getwd()
	if err != nil {
		lib.CriticalError(err)
	}

	answer = lib.BooleanQuestion("Do you want to use AutoMode?")
	// Check if there is an AutoMode profile
	if answer == true {
		ws.ReadJsonFile("autoMode.json")
		fmt.Println("*** AutoMode Activated ***")
	} else {
		fmt.Println("*** AutoMode Disabled ***")
		ws = checkProfiles(dir)
	}
	return
}
