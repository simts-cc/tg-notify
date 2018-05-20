package main

import (
	_ "github.com/joho/godotenv/autoload"
	"tg.notify/src"
)

func main() {
	tg.NewService().Start()
}
