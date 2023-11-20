package main

import (
	"context"
	"fmt"
	"runtime"
)

func main() {
	c := context.Background()
	someFunc(c)
}

func someFunc(ctx context.Context) {
	something := "cheese"
	fmt.Println("ctx: ", ctx)
	fmt.Println(something)
	pc, file, line, boolean := runtime.Caller(0)
	fmt.Println("file: ", file)
	fmt.Println("line: ", line)
	fmt.Println("boolean: ", boolean)
	fc := runtime.FuncForPC(pc)
	fmt.Println("name: ", fc.Name())
	fmt.Println("entry: ", fc.Entry())
}
