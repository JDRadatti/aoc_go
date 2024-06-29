package day02

type Context struct {
	red        Color
	green      Color
	blue       Color
	afterColon bool
	gameIDSum  int
	currCount  int
	gameCount  int
	validGame  bool
}

func (c *Context) validTurn() bool {
	return c.red.check() && c.green.check() && c.blue.check()
}

func (c *Context) updateColor(colorID byte) {
	switch colorID {
	case c.red.id:
		c.red.count += c.currCount
		c.currCount = 0
	case c.blue.id:
		c.blue.count += c.currCount
		c.currCount = 0
	case c.green.id:
		c.green.count += c.currCount
		c.currCount = 0
	}
}

func (c *Context) updateCount(val byte) {
	// multiply by 10 for base10.
	c.currCount = c.currCount*10 + int(val) - 48
}

func (c *Context) resetColor() {
	c.red.reset()
	c.green.reset()
	c.blue.reset()
}

func (c *Context) reset() {
	c.resetColor()
	c.afterColon = false
	c.currCount = 0
	c.gameCount += 1
	c.validGame = true
}

type Color struct {
	id    byte
	count int
	max   int
}

func (color *Color) check() bool {
	return color.count <= color.max

}
func (color *Color) reset() {
	color.count = 0
}

func SolutionA(input []byte) int {
	context := Context{
		red:        Color{id: 'r', count: 0, max: 13},
		green:      Color{id: 'g', count: 0, max: 12},
		blue:       Color{id: 'b', count: 0, max: 14},
		afterColon: false,
		gameIDSum:  0,
		gameCount:  1,
		validGame:  true,
	}

	for i := 0; i < len(input); i++ {
		if input[i] == ':' {
			context.afterColon = true
		}
		if !context.afterColon {
			continue
		}

		if input[i] >= 48 && input[i] <= 57 { // input[i] is int
			context.updateCount(input[i])
		} else if input[i] == 10 { // New line
			if !context.validTurn() {
				context.validGame = false
			}
			if context.validGame {
				context.gameIDSum += context.gameCount
			}
			context.reset()
		} else if  input[i] == 59 { // ';'
			if !context.validTurn() {
				context.validGame = false
			}
            context.resetColor()
            context.currCount = 0

        } else {
			context.updateColor(input[i])
		}
	}
	return context.gameIDSum
}
