package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)


func printHello(c *cli.Context) error {
	fmt.Println("Hello friend!")
	return nil
}

func run(args []string)  {
	
	app := &cli.App{
		Name: "greet",
		Usage: "fight the loneliness!",
		Action: printHello,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	run(os.Args)
}