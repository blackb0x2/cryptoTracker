package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Method to read Websocket struct as JSON
func (ws *Websocket) GetJson() []byte {
	// Write Json with four spaces of indentation
	contentJson, err := json.MarshalIndent(ws, "", "    ")
	if err != nil {
		CriticalError(err)
	}
	return contentJson
}

// Method to read JSON file as Websocket Struct
func (ws *Websocket) ReadJsonFile(filename string) {
	contentJson, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			PrintError("File doesn't exist")
		} else {
			CriticalError(err)
		}
	}

	err = json.Unmarshal(contentJson, &ws)
	if err != nil {
		CriticalError(err)
	}
	return
}

// Method to write Websocket struct as JSON
func (ws Websocket) WriteJson(filename string) {
	contentJson := ws.GetJson()

	file, err := os.Create(filename)
	if err != nil {
		CriticalError(err)
	}
	defer file.Close()

	// Save JSON
	file.Write(contentJson)
}
