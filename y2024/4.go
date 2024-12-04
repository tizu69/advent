package y2024

import (
	"bytes"

	"g.tizu.dev/advent/c"
)

type D4 struct{}

func (this *D4) Run(ctx c.Ctx) error {
	if ctx.Two {
		c.Output(-1, ctx.Two)
		return nil
	}

	count := 0

	lines := bytes.Split(ctx.Data, []byte("\r\n"))
	if len(lines) < 1 {
		lines = bytes.Split(ctx.Data, []byte("\n"))
	}

	for y, line := range lines {
		for x := range line {
			// basically this sets dX and dY to be -1..1, which tells the check how much to offset.
			for dX := -1; dX <= 1; dX++ {
				for dY := -1; dY <= 1; dY++ {
					matcher := "XMAS"

					// XMAS horizontally would be offset dX 1, as it takes 1 step right.
					for c, char := range []byte(matcher) {
						pX, pY := x+(c*dX), y+(c*dY)
						if pY < 0 || pY >= len(lines) || pX < 0 || pX >= len(lines[pY]) || lines[pY][pX] != char {
							break
						}
						if c == len(matcher)-1 {
							count++
						}
					}
				}
			}
		}
	}

	c.Output(count, ctx.Two)
	return nil
}
