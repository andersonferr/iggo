package main

import (
	"github.com/andersonferr/iggo"
	_ "github.com/andersonferr/iggo/ixlib"
)

func app() {
	w1 := iggo.CreateWindow("Janela 1", 300, 200)
	w1.SetVisibility(true)
}

func main() {
	iggo.Use("X11")
	iggo.Run(app)
}
