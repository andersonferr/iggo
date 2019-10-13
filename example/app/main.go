package main

import (
	"github.com/andersonferr/iggo"
)

func main() {
	w1 := iggo.CreateWindow("Oi", 0, 0, 400, 300)

	w2 := iggo.CreateWindow("Hello", 0, 0, 600, 500)

	w3 := iggo.CreateWindow("Hello 2", 10, 10, 600, 500)

	app := iggo.NewApplication()
	app.AddWindow(w1)
	app.AddWindow(w2)
	app.AddWindow(w3)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
