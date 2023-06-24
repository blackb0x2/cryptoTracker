package profile

import (
	lib "cryptoTracker/lib"

	"fmt"
	"strings"
)

// Loop to create the stream for every Trading Pair
func binanceLoopStream(streamSlice, pairSlice []string, market, streamType, timef string) []string {
	answer = lib.BooleanQuestion("Do you want to track all the saved Trading Pairs with this stream?")
	if answer == false {
		pairSlice = lib.MultipleInputSelection(pairSlice, "Please enter the Trading Pair : ")
	}

	n = len(pairSlice)

	for x = 0; x != n; x++ {
		if streamType == "kline" || streamType == "ticker" && market == "spot" {
			stream = fmt.Sprintf("%s@%s_%s", pairSlice[x], streamType, timef)
		} else {
			stream = fmt.Sprintf("%s@%s", pairSlice[x], streamType)
		}
		streamSlice = append(streamSlice, stream)
	}
	return streamSlice
}

// Function to create a Binance Spot websocket profile
func CreateBinanceSpotProfile(pairSlice []string) (streamSlice []string) {
	// Type of streams
	streamsType := []string{"Kline", "MiniTicker", "Rolling Window"}
	klineIntervals := []string{"1s", "1m", "3m", "5m", "15m", "30m",
		"1h", "2h", "4h", "6h", "8h", "12h", "1d", "3d", "1w", "1M"}
	tickerIntervals := []string{"1h", "4h", "1d"}

	fmt.Println("*** Spot Profile Creation ***")
	for {
		if len(streamSlice) != 0 {
			answer = lib.BooleanQuestion("Do you want to add more streams?")
			if answer == false {
				break
			}
		}

		x = lib.MultipleSelection(streamsType, "Please select stream type for the websocket")

		switch streamsType[x-1] {
		case "Kline":
			x = lib.MultipleSelection(klineIntervals, "Please select Kline interval")
			interval = klineIntervals[x-1]
			streamSlice = binanceLoopStream(streamSlice, pairSlice, "spot", "kline", interval)

		case "MiniTicker":
			streamSlice = binanceLoopStream(streamSlice, pairSlice, "spot", "miniTicker", "")

		case "Rolling Window":
			x = lib.MultipleSelection(tickerIntervals, "Please select the interval")
			interval = tickerIntervals[x-1]
			streamSlice = binanceLoopStream(streamSlice, pairSlice, "spot", "ticker", interval)
		}
	}
	return
}

// Function to create a Binance USDM websocket profile
func CreateBinanceUSDMProfile(pairSlice []string) (streamSlice []string) {
	// Type of streams
	streamsType := []string{"Kline", "MarkPrice", "MiniTicker", "Ticker"}
	klineIntervals := []string{"1s", "1m", "3m", "5m", "15m", "30m",
		"1h", "2h", "4h", "6h", "8h", "12h", "1d", "3d", "1w", "1M"}

	fmt.Println("*** USDM Profile Creation ***")
	for {
		if len(streamSlice) != 0 {
			answer = lib.BooleanQuestion("Do you want to add more streams?")
			if answer == false {
				break
			}
		}

		x = lib.MultipleSelection(streamsType, "Please select stream type for the websocket")

		switch streamsType[x-1] {
		case "Kline":
			x = lib.MultipleSelection(klineIntervals, "Please select Kline interval")
			interval = klineIntervals[x-1]
			streamSlice = binanceLoopStream(streamSlice, pairSlice, "usdm", "kline", interval)

		case "MarkPrice":
			streamSlice = binanceLoopStream(streamSlice, pairSlice, "usdm", "markPrice@1s", "")

		case "MiniTicker":
			streamSlice = binanceLoopStream(streamSlice, pairSlice, "usdm", "miniTicker", "")

		case "Ticker":
			streamSlice = binanceLoopStream(streamSlice, pairSlice, "usdm", "ticker", "")
		}
	}
	return
}

// General function to create a Binance websocket profile
func CreateBinanceProfile() (ws Websocket) {
	ws.Exchange = "binance"

	markets := []string{"Spot", "USDM"}

	// Select the Binance market for the websockets
	fmt.Println("*** Markets ***")
	x = lib.MultipleSelection(markets, "Which Binance Market do you want to connect to?")
	b = lib.BooleanQuestion("Do you want to add the stream/s later?")

	ws.Market = strings.ToLower(markets[x-1])

	if b == false {
		pairSlice := TradingPairs("lower")

		switch ws.Market {
		case "spot":
			ws.Streams = CreateBinanceSpotProfile(pairSlice)

		case "usdm":
			ws.Streams = CreateBinanceUSDMProfile(pairSlice)
		}
	}
	return
}
