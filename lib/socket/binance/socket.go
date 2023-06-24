package socket

import (
	lib "cryptoTracker/lib"
	number "cryptoTracker/lib/number"
	"sync"

	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Function to edit the JSON map from Binance Websockets, and remove unnecessary fields
func editMap(dataMap *map[string]any, streamType string, market string) {
	deleteKey = nil

	switch streamType {
	case "kline":
		deleteKey = append(deleteKey, "t", "T", "f", "L", "v", "n", "x", "V", "Q", "B", "Q")

	case "markPriceUpdate":
		deleteKey = append(deleteKey, "E", "P", "r", "T")

	case "24hrMiniTicker":
		deleteKey = append(deleteKey, "E", "v")

	case "24hrTicker":
		deleteKey = append(deleteKey, "E", "w", "v", "O", "C", "F", "L", "n", "Q")

	// 1hTicker, 4hTicker, 1dTicker
	default:
		deleteKey = append(deleteKey, "E", "w", "v", "O", "C", "F", "L", "n")
	}

	n = len(deleteKey)
	for x = n; x != 0; x-- {
		delete(*dataMap, deleteKey[x-1])
	}
}

// Function to print the data collected by the Binance Websocket
func binancePrint(binanceMap *map[string]map[string]map[string]any) {
	lib.ClearTerminal()

	sort.Strings(streams)

	titleColor := color.New(color.BgWhite, color.FgHiBlack, color.Bold)
	contentColor := color.New(color.FgWhite)

	for _, value := range streams {
		streamTypeMap = (*binanceMap)[value]
		keys := SortMap(streamTypeMap, o)
		o = len(streamTypeMap)

		switch value {
		// Print kline streams if exists
		case "kline":
			content = []string{"Symbol", "Timeframe", "Open", "High", "Low", "Close", "Quote Volume"}
			padding = []int{colLen[0], colLen[1], colLen[2], colLen[2], colLen[2], colLen[2], colLen[4]}

			ColLen = -1
			n = len(padding)
			for x = 0; x < n; x++ {
				ColLen += padding[x] + 1
			}

			title = "Kline"
			titleSpaceLeft, titleSpaceRight = lib.CenterText(title, ColLen)

			titleColor.Printf("%s%s%s \n",
				strings.Repeat(" ", titleSpaceLeft), title, strings.Repeat(" ", titleSpaceRight))

			for x = 0; x < 7; x++ {
				if x != 6 {
					titleColor.Printf("%*s ", padding[x], content[x])
				} else {
					titleColor.Printf("%*s \n", padding[x], content[x])
				}
			}

			for x = 0; x < o; x++ {
				stream = keys[x]
				dataMap = streamTypeMap[stream]

				content = []string{dataMap["s"].(string), dataMap["i"].(string), dataMap["o"].(string),
					dataMap["h"].(string), dataMap["l"].(string), dataMap["c"].(string), dataMap["q"].(string)}

				m = len(content)
				for y = 0; y < m; y++ {
					if y != (m - 1) {
						contentColor.Printf("%*s ", padding[y], content[y])
					} else {
						contentColor.Printf("%*s \n", padding[y], content[y])
					}
				}
			}

		// Print markPriceUpdate streams if exists
		case "markPriceUpdate":
			content = []string{"Symbol", "Mark", "Index"}
			padding = []int{colLen[0], colLen[2], colLen[2]}

			ColLen = -1
			n = len(padding)
			for x = 0; x < n; x++ {
				ColLen += padding[x] + 1
			}

			title = "Mark Price"
			titleSpaceLeft, titleSpaceRight = lib.CenterText(title, ColLen)

			titleColor.Printf("%s%s%s \n",
				strings.Repeat(" ", titleSpaceLeft), title, strings.Repeat(" ", titleSpaceRight))

			for x = 0; x < 3; x++ {
				if x != 2 {
					titleColor.Printf("%*s ", padding[x], content[x])
				} else {
					titleColor.Printf("%*s \n", padding[x], content[x])
				}
			}

			for x = 0; x < o; x++ {
				stream = keys[x]
				dataMap = streamTypeMap[stream]

				content = []string{dataMap["s"].(string), dataMap["p"].(string), dataMap["i"].(string)}

				m = len(content)
				for y = 0; y < m; y++ {
					if y != (m - 1) {
						contentColor.Printf("%*s ", padding[y], content[y])
					} else {
						contentColor.Printf("%*s \n", padding[y], content[y])
					}
				}
			}

		// Print 24hrMiniTicker streams if exists
		case "24hrMiniTicker":
			content = []string{"Symbol", "Open", "High", "Low", "Close", "Quote Volume"}
			padding = []int{colLen[0], colLen[2], colLen[2], colLen[2], colLen[2], colLen[4]}

			ColLen = -1
			n = len(padding)
			for x = 0; x < n; x++ {
				ColLen += padding[x] + 1
			}

			title = "24hr Mini Ticker"
			titleSpaceLeft, titleSpaceRight = lib.CenterText(title, ColLen)

			titleColor.Printf("%s%s%s \n",
				strings.Repeat(" ", titleSpaceLeft), title, strings.Repeat(" ", titleSpaceRight))

			for x = 0; x < 6; x++ {
				if x != 5 {
					titleColor.Printf("%*s ", padding[x], content[x])
				} else {
					titleColor.Printf("%*s \n", padding[x], content[x])
				}
			}

			for x = 0; x < o; x++ {
				stream = keys[x]
				dataMap = streamTypeMap[stream]

				content = []string{dataMap["s"].(string), dataMap["o"].(string), dataMap["h"].(string),
					dataMap["l"].(string), dataMap["c"].(string), dataMap["q"].(string)}

				m = len(content)
				for y = 0; y < m; y++ {
					if y != (m - 1) {
						contentColor.Printf("%*s ", padding[y], content[y])
					} else {
						contentColor.Printf("%*s \n", padding[y], content[y])
					}
				}
			}

		// Print 24hrTicker, 1hTicker, 4hTicker or/and 1dTicker streams if exists
		default:
			content = []string{"Symbol", "Open", "High", "Low", "Close", "Change", "Change %", "Quote Volume"}
			padding = []int{colLen[0], colLen[2], colLen[2], colLen[2], colLen[2], colLen[3], colLen[3], colLen[4]}

			ColLen = -1
			n = len(padding)
			for x = 0; x < n; x++ {
				ColLen += padding[x] + 1
			}

			switch value {
			case "24hrTicker":
				title = "24hr Ticker"
			case "1hTicker":
				title = "Rolling Window: 1h Ticker"
			case "4hTicker":
				title = "Rolling Window: 4h Ticker"
			case "1dTicker":
				title = "Rolling Window: 1d Ticker"
			}
			titleSpaceLeft, titleSpaceRight = lib.CenterText(title, ColLen)

			titleColor.Printf("%s%s%s \n",
				strings.Repeat(" ", titleSpaceLeft), title, strings.Repeat(" ", titleSpaceRight))

			for x = 0; x < 8; x++ {
				if x != 7 {
					titleColor.Printf("%*s ", padding[x], content[x])
				} else {
					titleColor.Printf("%*s \n", padding[x], content[x])
				}
			}

			for x = 0; x < o; x++ {
				stream = keys[x]
				dataMap = streamTypeMap[stream]

				content = []string{dataMap["s"].(string), dataMap["o"].(string), dataMap["h"].(string),
					dataMap["l"].(string), dataMap["c"].(string), dataMap["o"].(string),
					dataMap["P"].(string), dataMap["q"].(string)}

				m = len(content)
				for y = 0; y < m; y++ {
					if y != (m - 1) {
						contentColor.Printf("%*s ", padding[y], content[y])
					} else {
						contentColor.Printf("%*s \n", padding[y], content[y])
					}
				}
			}
		}
		fmt.Printf("\n")
	}
}

// Function to read []byte data comming from byte channel with Binance Websocket
func ReadBinanceByte(wg *sync.WaitGroup, msg []byte, binanceMap *map[string]map[string]map[string]any, market string) {
	// Convert JSON byte array to JSON map
	json.Unmarshal([]byte(msg), &jsonMap)

	stream = jsonMap["stream"].(string)
	dataMap = jsonMap["data"].(map[string]any)
	streamType = dataMap["e"].(string)

	for _, value := range streams {
		if value == streamType {
			streamExists = true
			break
		} else {
			streamExists = false
		}
	}

	if streamExists != true {
		streams = append(streams, streamType)
	}

	if dataMap["e"] == "kline" {
		dataMap = dataMap["k"].(map[string]any)
		dataMap["e"] = "kline"
	}

	editMap(&dataMap, streamType, market)

	// Round values
	if streamType == "markPriceUpdate" {
		roundValues = []string{"e", "p", "i"}
	} else if streamType != "24hrMiniTicker" && strings.HasSuffix(streamType, "Ticker") == true {
		roundValues = []string{"o", "l", "h", "c", "p", "P", "q"}
	} else {
		roundValues = []string{"o", "l", "h", "c", "q"}
	}

	n = len(roundValues)
	for x = 0; x < n; x++ {
		_, exist := dataMap[roundValues[x]]
		if exist {
			k = roundValues[x]
			value, _ = strconv.ParseFloat(dataMap[roundValues[x]].(string), 64)
			value = number.RoundValue(value)
			dataMap[k] = strconv.FormatFloat(value, 'f', -1, 64)
		}
	}

	_, ok := (*binanceMap)[streamType]
	if !ok {
		(*binanceMap)[streamType] = make(map[string]map[string]any)
	}
	(*binanceMap)[streamType][stream] = dataMap

	binancePrint(binanceMap)
	wg.Done()
}
