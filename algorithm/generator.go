package algorithm

import (
	"fmt"
	"strconv"
	"time"
)

func GeneratePeselsForDate(date time.Time, passChannel chan string) {
	peselDate := date.Format("060102")
	if date.Year() >= 2000 {
		peselDate = peselDate[:2] + "2" + peselDate[3:]
	}
	for i := 0; i < 10000; i++ {
		peselWithoutValidation := peselDate + fmt.Sprintf("%04d", i)
		checksum := ControlSum(peselIntArray(peselWithoutValidation))
		pesel := peselWithoutValidation + strconv.Itoa(checksum)
		passChannel <- pesel
	}
	fmt.Println("Nothing found for: ", date)
	close(passChannel)
}

func peselIntArray(pesel string) [10]int {
	var peselArray [10]int
	for idx, digit := range pesel {
		peselArray[idx] = int(digit) - '0'
	}
	return peselArray
}
