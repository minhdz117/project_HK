package main

import (
	"gin_API/api"
)

func main() {
	app := api.InitAPI()
	app.Run(":3000")
}
