package y2024

import (
	"regexp"
	"strconv"

	"g.tizu.dev/advent/c"
)

type D3 struct{}

func (this *D3) Run(ctx c.Ctx) error {
	output, enabled := 0, true

	rgxAll := regexp.MustCompile(`do\(\)|don\'t\(\)|mul\(\d+,\d+\)`)
	rgxMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	cmds := rgxAll.FindAllString(string(ctx.Data), -1)

	for _, cmd := range cmds {
		if cmd == "do()" && ctx.Two {
			enabled = true
		} else if cmd == "don't()" && ctx.Two {
			enabled = false
		} else if match := rgxMul.FindStringSubmatch(cmd); len(match) == 3 && enabled {
			n1, _ := strconv.Atoi(match[1])
			n2, _ := strconv.Atoi(match[2])
			output += n1 * n2
		}
	}

	c.Output(output, ctx.Two)
	return nil
}
