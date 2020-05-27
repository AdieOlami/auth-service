package main

import "github.com/AdieOlami/auth-service/src/app"

func main() {
	start := app.App{}
	start.StartApp()
	start.Run(":9091")
}
