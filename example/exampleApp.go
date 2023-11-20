package main

import (
	"errors"
	"fmt"
	log "github.com/davidthorpe71/dt-logger/logger"
	"github.com/google/uuid"
)

type MyLogger interface {
	OpenContext()
	CloseContext()
	AddResponse(res interface{})
	AddArg(key string, arg interface{})
	AddError(err error)
	Write() string
}

func main() {
	uuidMaker := uuid.NewString

	l := log.NewLog("example-app", uuidMaker)

	first := SomeFunc(l, "abcd")

	res := whatNow(l, first)

	logStr := l.Write()

	fmt.Println("log: ", logStr)
	fmt.Println("res: ", res)
}

func whatNow(l MyLogger, someArg string) string {
	l.OpenContext()
	defer l.CloseContext()

	l.AddArg("someArg", someArg)

	strRes, err := hereYaAre(l, "cheese")

	if err != nil {
		l.AddResponse(someArg)
		return someArg
	}

	l.AddResponse(strRes + "-xyz")
	return someArg + "-qrstuvw-xyz"
}

func hereYaAre(l MyLogger, inputString string) (string, error) {
	l.OpenContext()
	defer l.CloseContext()

	l.AddArg("inputString", inputString)

	if inputString == "cheese" {
		err := errors.New("the input was cheese we can't go on")
		l.AddError(err)
		l.AddResponse("")
		return "", err
	}

	return inputString + "-qrstuvw", nil
}

func SomeFunc(l MyLogger, s string) string {
	l.OpenContext()
	defer l.CloseContext()

	l.AddArg("s", s)

	a := AnotherFunc(l, s+"-efg")

	l.AddResponse(a)
	return a
}

func AnotherFunc(l MyLogger, s string) string {
	l.OpenContext()
	defer l.CloseContext()

	l.AddArg("s", s)

	b := s + "-hijk-lmnop"

	l.AddResponse(b)
	return b
}
