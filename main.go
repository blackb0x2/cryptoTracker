package main

import (
	lib "cryptoTracker/lib"
	profile "cryptoTracker/lib/profile"
	binanceSocket "cryptoTracker/lib/socket/binance"
	"sync"

	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	bytesChan                     chan []byte
	interruptChan                 chan os.Signal
	answer                        bool
	socketUrl, streamType, stream string
	exchange, market              *string
	binanceMap                    map[string]map[string]map[string]any
)

// Function to receive data
func receiver(connection *websocket.Conn) {
	// Close bytesChan if there was an error
	defer close(bytesChan)
	defer close(interruptChan)

	// Loop to get Websocket messages
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			break
		} else {
			bytesChan <- msg
		}
	}
}

func main() {
	fmt.Println("*** Welcome to CryptoTracker ***")
	fmt.Println("made by: blackb0x")
	ws := profile.ProfileSelection()

	lib.ClearTerminal()

	fmt.Println("*** Websocket ***")
	fmt.Println("Type 'ctrl' + 'c' to exit")
	answer = lib.BooleanQuestion("Do you want to start streaming?")
	if answer == false {
		os.Exit(0)
	} else {
		exchange = &ws.Exchange
		market = &ws.Market

		switch *exchange {
		case "binance":
			binanceMap = make(map[string]map[string]map[string]any)
			socketUrl = binanceSocket.GetSocketUrl(ws)
		}
	}

	// Websocket channels
	bytesChan = make(chan []byte)
	interruptChan = make(chan os.Signal)
	wg := new(sync.WaitGroup)

	// Notify the interrupt channel
	signal.Notify(interruptChan, os.Interrupt)

	// Connects to websocket
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		lib.CriticalError(err)
	}
	defer conn.Close()

	go receiver(conn)

	for {
		select {
		case <-bytesChan:
			switch *exchange {
			case "binance":
				msg, ok := <-bytesChan

				// Check if channel is closed
				if !ok {
					fmt.Println("*** Finished ***")
					os.Exit(0)
				} else {
					wg.Add(1)
					go binanceSocket.ReadBinanceByte(wg, msg, &binanceMap, *market)
					wg.Wait()
				}
			}
		case <-interruptChan:
			// Close websocket connection if the user typed 'ctrl' + 'c'
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			select {
			case <-bytesChan:
			case <-time.After(time.Duration(500) * time.Millisecond):
				os.Exit(0)
			}
		}
	}
}
