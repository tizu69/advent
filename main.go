package main

import (
	"os"

	"g.tizu.dev/advent/c"
	"g.tizu.dev/advent/y2024"
	"github.com/alecthomas/kong"
)

var cli struct {
	Input string `help:"Input file path (or - for stdin)" type:"existingfile" required:""`
	Two   bool   `help:"Use part 2?"`

	Y2024 y2024.Y `cmd:"" name:"2024"`
}

func main() {
	ctx := kong.Parse(&cli, kong.UsageOnError())
	f, err := os.ReadFile(cli.Input)
	if err != nil {
		panic(err)
	}
	err = ctx.Run(c.Ctx{Data: f, Two: cli.Two})
	ctx.FatalIfErrorf(err)
}
