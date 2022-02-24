package algorithm

import "strconv"

func getLastDigit(number int) int {
	strNumber := strconv.Itoa(number)
	lastDigit := strNumber[len(strNumber)-1:]
	result, _ := strconv.Atoi(lastDigit)
	return result
}

func ControlSum(peselArr [10]int) int {
	result := peselArr[0]*1 +
		peselArr[1]*3 +
		peselArr[2]*7 +
		peselArr[3]*9 +
		peselArr[4]*1 +
		peselArr[5]*3 +
		peselArr[6]*7 +
		peselArr[7]*9 +
		peselArr[8]*1 +
		peselArr[9]*3
	lastDigit := getLastDigit(result)
	if lastDigit == 0 {
		return 0
	} else {
		return 10 - getLastDigit(result%10)
	}
}
