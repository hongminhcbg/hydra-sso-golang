package main

import (
	"fmt"

	"hydra-sso/app"
)

func main()  {
	fmt.Println("hello world")
	application := app.NewApp()
	application.Init()
	application.Run()
}