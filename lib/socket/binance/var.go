package socket

import lib "cryptoTracker/lib"

type Websocket = lib.Websocket

var (
	streamExists                                           bool
	m, n, o, x, y, ColLen, titleSpaceLeft, titleSpaceRight int
	padding                                                []int
	colLen                                                 = []int{8, 9, 10, 10, 15}
	value                                                  float64
	k, stream, streamType, title                           string
	deleteKey, roundValues, data, streams, content         []string
	jsonMap, dataMap                                       map[string]any
	streamTypeMap                                          = make(map[string]map[string]any)
)
