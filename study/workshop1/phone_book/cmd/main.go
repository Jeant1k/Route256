package main

import (
	"fmt"
	"os"

	"\\wsl.localhost\Ubuntu\home\jeantik\programs\route256\study\workshop1\phone_book\internal\cli\cli.go"
	"\\wsl.localhost\Ubuntu\home\jeantik\programs\route256\study\workshop1\phone_book\internal\module\module.go"
	"\\wsl.localhost\Ubuntu\home\jeantik\programs\route256\study\workshop1\phone_book\internal\storage\storage.go"
)

const (
	fileName = "telephone_book.json"
)

func main() {
	storageJSON := storage.NewStorage(fileName)
	phoneBookService := module.NewModule(module.Deps{
		Storage: storageJSON
	})
	commands := cli.NewCLI(cli.Deps{Module: module})
	
	if err := commands.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

