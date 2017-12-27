package Servers

import (
	"net/http"
)

type LowLevelRequestHandler func(rw http.ResponseWriter, r *http.Request)
type JsonRequestHandler func(data interface{}) JsonResponse
type RealtimeRequestHandler func(inChannel chan []byte, outChannel chan []byte, done chan bool)

type LowLevelHandler struct {
	Route   string
	Handler LowLevelRequestHandler
	Verb    string
}

type JsonHandler struct {
	Route      string
	Handler    JsonRequestHandler
	JsonObject interface{}
}

type JsonResponse struct {
	Error string `json:"Error"`
	Data  interface{}
}

type JsonData map[string]interface{}

type RealtimeHandler struct {
	Route   string
	Handler RealtimeRequestHandler
}

type ApiServer struct {
	lowLevelHandlers []LowLevelHandler
	jsonHandlers     []JsonHandler
	realtimeHandlers []RealtimeHandler
}