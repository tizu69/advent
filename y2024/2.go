package y2024

import (
	"strconv"
	"strings"

	"g.tizu.dev/advent/c"
)

type D2 struct{}

func (this *D2) Run(ctx c.Ctx) error {
	if ctx.Two {
		c.Output(-1, ctx.Two)
		return nil
	}

	safe := 0

	for _, report := range strings.Split(string(ctx.Data), "\n") {
		if len(report) == 0 {
			continue
		}
		numbers, err := this.parseNumbers(report)
		if err != nil {
			return err
		}
		if this.isSafe(numbers) {
			safe++
		}
	}

	c.Output(safe, ctx.Two)
	return nil
}

func (*D2) parseNumbers(report string) ([]int, error) {
	var numbers []int
	for _, snum := range strings.Fields(report) {
		num, err := strconv.Atoi(snum)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func (*D2) isSafe(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}
	var isInc, isDec bool
	for i := 1; i < len(numbers); i++ {
		delta := c.Delta(numbers[i-1], numbers[i])
		if delta < 1 || delta > 3 {
			return false
		}
		if numbers[i] > numbers[i-1] {
			isInc = true
		} else if numbers[i] < numbers[i-1] {
			isDec = true
		}
	}
	return (isInc || isDec) && !(isInc && isDec)
}
