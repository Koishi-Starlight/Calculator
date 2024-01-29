package calc

import "log"

const (
	add         = "+"
	subtract    = "-"
	multiply    = "*"
	divide      = "/"
	divisionErr = "Zero division error"
)

type calculator struct{}

func NewCalc() calculator {
	return calculator{}
}

func (c *calculator) Calculate(num1, num2 float64, operator string) float64 {
	switch operator {
	case add:
		return c.add(num1, num2)
	case subtract:
		return c.subtract(num1, num2)
	case multiply:
		return c.multiply(num1, num2)
	case divide:
		return c.divide(num1, num2)
	default:
		log.Printf("unsupported operator: %s\n", operator)
		return 0
	}
}

func (c *calculator) add(num1, num2 float64) float64 {
	return num1 + num2
}

func (c *calculator) subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func (c *calculator) divide(num1, num2 float64) float64 {
	if num2 == 0 {
		log.Println("Zero division error")
		return 0
	}
	return num1 / num2
}

func (c *calculator) multiply(num1, num2 float64) float64 {
	return num1 * num2
}
