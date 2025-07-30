package main

import (
	"fmt"

	"github.com/DKeshavarz/armis/internal/commands"
)

func main(){
	cmd := commands.New()
	fmt.Println("cmd:", cmd)
	cmd.Run()
}