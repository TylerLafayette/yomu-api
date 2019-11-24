package main

import "log"

func main() {
	config := CreateConfig()
	router := CreateRouter()
	app := NewApp(config, router)

	err := app.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
