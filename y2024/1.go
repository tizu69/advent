package y2024

import (
	"slices"
	"strconv"

	"g.tizu.dev/advent/c"
)

type D1 struct{}

func (*D1) Run(ctx c.Ctx) error {
	left, right := []int{}, []int{}
	current, isRight := "", false
	output := 0

	apply := func() error {
		i, err := strconv.Atoi(current)
		if err != nil {
			return err
		}

		if !isRight {
			left = append(left, i)
		} else {
			right = append(right, i)
		}

		current, isRight = "", !isRight
		return nil
	}

	ctx.Data = append(ctx.Data, '\n')
	for _, b := range ctx.Data {
		switch b {
		case '\r':
			continue
		case '\n', ' ':
			if current == "" {
				continue
			}
			if err := apply(); err != nil {
				return err
			}
		default:
			current = current + string(b)
		}
	}

	if !ctx.Two {
		slices.Sort(left)
		slices.Sort(right)
		for i, l := range left {
			r := right[i]
			output += c.Abs(l - r)
		}
	} else {
		for _, l := range left {
			r := len(c.SlicesFilter(right, func(i int) bool { return i == l }))
			output += l * r
		}
	}

	c.Output(output, ctx.Two)
	return nil
}
