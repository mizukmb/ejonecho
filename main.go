package main

import (
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"os"
)

func action(c *cli.Context) error {
	fmt.Printf(":ejoneco: < ")
	if terminal.IsTerminal(0) {
		for i := 0; i < c.NArg(); i++ {
			fmt.Printf("%s ", c.Args().Get(i))
		}
		fmt.Println()
	} else {
		b, _ := ioutil.ReadAll(os.Stdin)
		fmt.Printf(string(b))
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "ejonecho"
	app.Version = "0.0.1"
	app.Usage = "ejo090 ga echo suru yatsu."
	app.Action = action

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
