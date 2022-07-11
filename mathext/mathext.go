package mathext

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//KeepDecimal keep a specified number of decimals
//v: the number
//n: how many decimals will be kept
func KeepDecimal(v float64, n int32) float64 {
	dividor := float64(1)
	for i := int32(0); i < n; i++ {
		dividor *= 10
	}
	return (float64(int64(v*dividor)) / dividor)
}

//RandByMinMax get a rand number from [min, max]
func RandByMinMax(min, max int) int {
	if min >= max {
		return min
	}
	return min + rand.Intn(max+1-min)
}

func TryParseIntVal(strVal string) (int64, error) {
	strVal = strings.TrimSpace(strVal)
	val, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		return 0, err
	}
	intVal := int64(val)
	if float64(intVal) != val {
		return 0, fmt.Errorf("parse to int failed")
	}
	return intVal, nil
}

func TryParseFloatVal(strVal string) (float64, error) {
	strVal = strings.TrimSpace(strVal)
	val, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func TryParseStringVal(strVal string) (string, error) {
	return strVal, nil
}

func TryParseBoolVal(strVal string) (bool, error) {
	strVal = strings.TrimSpace(strVal)
	val, err := strconv.ParseBool(strVal)
	if err != nil {
		return false, err
	}
	return val, nil
}
