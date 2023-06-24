package lib

type Websocket struct {
	Exchange string   `json:"exchange"`
	Market   string   `json:"market"`
	Streams  []string `json:"streams"`
}
