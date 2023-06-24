package number

import "math"

var (
	dataPercentChange, powTen, roundNumber float64
)

func roundFloat(number float64, precision uint) float64 {
	powTen = math.Pow(10, float64(precision))
	roundNumber = math.Round(number * powTen)
	return roundNumber / powTen
}

func RoundValue(value float64) float64 {
	if value >= 10 {
		value = roundFloat(value, 2)
	} else if value >= 0.1 {
		value = roundFloat(value, 3)
	} else if value >= 0.01 {
		value = roundFloat(value, 4)
	} else if value >= 0.0001 {
		value = roundFloat(value, 5)
	} else if value > 0 {
		value = 0
	} else if value > -0.0001 {
		value = roundFloat(value, 5)
	} else if value > -0.01 {
		value = roundFloat(value, 4)
	} else if value > -0.1 {
		value = roundFloat(value, 3)
	} else {
		value = roundFloat(value, 2)
	}
	return value
}

// Function to calculate percent change between two numbers
func DataChange(firstData, secondData float64) float64 {
	dataPercentChange = roundFloat(((secondData * 100) / firstData), 2)
	return dataPercentChange
}
