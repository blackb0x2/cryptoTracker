package profile

import (
	lib "cryptoTracker/lib"

	"fmt"
	"os"
	"strings"
)

// Add Trading Pairs that will be tracked
func TradingPairs(format string) (pairSlice []string) {
	fmt.Println("*** Trading Pairs ***")

	pairList := lib.MultipleInputString("Enter the Trading Pair: ", 6, 9)

	for _, pairs := range pairList {
		if format == "lower" {
			pairSlice = append(pairSlice, strings.ToLower(pairs))
		}
	}
	return
}

// Function to create the websocket profile
func CreateProfile(dir string) (ws Websocket) {
	ex := []string{"Binance"}

	// Select the exchange of the websocket
	if len(ex) == 1 {
		fmt.Printf("*** %s selected ***\n", ex[0])
		x = 1
	} else {
		fmt.Println("*** Exchanges ***")
		x = lib.MultipleSelection(ex, "Which Exchange do you want to connect to?")
	}

	switch ex[x-1] {
	case "Binance":
		ws = CreateBinanceProfile()
	}

	fmt.Println("*** Data ***")
	answer = lib.BooleanQuestion("Do you want to save this profile?")
	// Save profile
	if answer == true {
		fmt.Print("Please enter profile name: ")
		fmt.Scan(&name)

		filename := name + ".json"
		ws.WriteJson(filename)
	}

	answer = lib.BooleanQuestion("Do you want to save this as AutoMode?")
	// Save AutoMode
	if answer == true {
		os.Chdir(dir)
		ws.WriteJson("autoMode.json")
	}
	return
}
