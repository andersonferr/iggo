package main

import (
	"github.com/andersonferr/iggo"
)

func main() {
	w := iggo.CreateWindow("Oi", 0, 0, 400, 300)

	w2 := iggo.CreateWindow("Hello", 0, 0, 20, 50)

	iggo.Use("XGB")
	app := iggo.NewApplication()
	app.AddWindow(w)
	app.AddWindow(w2)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
