package log

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type LogContext struct {
	Calls    []*LogContext          `json:"calls"`
	Order    int                    `json:"order"`
	Context  string                 `json:"context"`
	File     string                 `json:"file"`
	Parent   *LogContext            `json:"-"`
	Args     map[string]interface{} `json:"args,omitempty"`
	Response interface{}            `json:"response,omitempty"`
	Error    string                 `json:"error,omitempty"`
	Line     int                    `json:"line"`
}

type Log struct {
	ApplicationName string        `json:"applicationName"`
	Calls           []*LogContext `json:"calls"`
	ActiveContext   *LogContext   `json:"-"`
	UuidMaker       func() string `json:"-"`
	Priority        string        `json:"priority"`
}

func NewLog(applicationName string, uuidMaker func() string) *Log {
	return &Log{
		UuidMaker:       uuidMaker,
		ApplicationName: applicationName,
		ActiveContext:   nil,
		Calls:           []*LogContext{},
		Priority:        "INFO",
	}
}

func (l *Log) OpenContext() {
	pc, file, _, _ := runtime.Caller(1)

	previousContext := l.ActiveContext

	newContext := &LogContext{
		Calls:   []*LogContext{},
		Context: runtime.FuncForPC(pc).Name(),
		File:    file,
		Parent:  previousContext,
		Args:    map[string]interface{}{},
	}

	if l.ActiveContext != nil {
		newContext.Order = len(previousContext.Calls)
		previousContext.Calls = append(previousContext.Calls, newContext)
		l.ActiveContext = newContext
	} else {
		newContext.Order = len(l.Calls)
		l.Calls = append(l.Calls, newContext)
		l.ActiveContext = newContext
	}
}

func (l *Log) CloseContext() {
	_, _, line, _ := runtime.Caller(1)
	l.ActiveContext.Line = line
	ac := l.ActiveContext.Parent

	l.ActiveContext = ac
}

func (l *Log) AddArg(key string, arg interface{}) {
	ctx := l.ActiveContext

	ctx.Args[key] = arg
}

func (l *Log) AddResponse(res interface{}) {
	ctx := l.ActiveContext

	ctx.Response = res
}

func (l *Log) AddError(err error) {
	ctx := l.ActiveContext
	ctx.Error = err.Error()

	l.Priority = "ERROR"
}

func (l *Log) Write() string {
	indent, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	return string(indent)
}
