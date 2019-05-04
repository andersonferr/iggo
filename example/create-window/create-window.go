package main

import (
	"github.com/andersonferr/iggo"
	_ "github.com/andersonferr/iggo/defaultenv"
)

func app() {
	w1 := iggo.CreateWindow("Janela 1", 300, 200)

	w1.SetVisibility(true)
}

func main() {
	iggo.Run(app)
}
