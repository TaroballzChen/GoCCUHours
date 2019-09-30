package main

import (
	"fmt"
	"github.com/curtis992250/GoCCUHours/MainMenu"
)

func main(){
	mm := MainMenu.InitMainMenu()


	if err := mm.Show(); err != nil {
		fmt.Println(err)
	}
}

