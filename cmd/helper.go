package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrCouldNotReadBody       = errors.New("could not read body")
	ErrCouldNotParseBody      = errors.New("could not parse body")
	ErrorCouldNotMarshallBody = errors.New("could not marshall body")
)

func getBody(r *http.Request, v interface{}) error {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return ErrCouldNotReadBody
	}

	err = json.Unmarshal(body, v)

	if err != nil {
		return ErrCouldNotParseBody
	}
	return nil
}

type HTTPRes struct {
	Status  int
	Message string
	Data    interface{}
}

func sendRes(rw http.ResponseWriter, status int, message string, data interface{}) {
	out, err := json.Marshal(HTTPRes{Status: status, Message: message, Data: data})

	fmt.Println(string(out))
	if err != nil {
		sendErrRes(rw, 500, "internal server error", nil)
	}
	rw.Write(out)
}

func sendErrRes(rw http.ResponseWriter, status int, message string, data interface{}) {
	out, _ := json.Marshal(HTTPRes{Status: status, Message: message, Data: data})
	rw.Write(out)
}
