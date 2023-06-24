package socket

import (
	lib "cryptoTracker/lib"
	profile "cryptoTracker/lib/profile"

	"os"
)

// Function to get Binance socket URL from JSON profile
func GetSocketUrl(ws Websocket) string {
	var socketUrl string

	// Append Socket Base Url if JSON fields are correct
	switch ws.Market {
	case "spot":
		socketUrl = "wss://stream.binance.com:443"
	case "usdm":
		socketUrl = "wss://fstream.binance.com"
	default:
		lib.PrintError("Market field modified... Exiting")
		os.Exit(0)
	}

	// For loop for Binance profiles with empty streams
	for {
		n = len(ws.Streams)

		if n != 0 {
			break
		} else {
			pairSlice := profile.TradingPairs("lower")

			if ws.Market == "spot" {
				ws.Streams = profile.CreateBinanceSpotProfile(pairSlice)
			} else {
				ws.Streams = profile.CreateBinanceUSDMProfile(pairSlice)
			}
		}
	}

	// Append Socket Path URL
	if n >= 1 {
		socketUrl += "/stream?streams="
	} else {
		socketUrl += "/ws/"
	}

	// For loop to create Socket URL
	for x = n; x != 0; x-- {
		if x != n {
			socketUrl += "/" + ws.Streams[x-1]
		} else {
			socketUrl += ws.Streams[x-1]
		}
	}
	return socketUrl
}
