package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/galuszkak/pesel/algorithm"
)

func main() {
	channel := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	if len(os.Args) != 2 {
		fmt.Println("You need to provide path to PDF\nExample: go run main.go ./PIT.pdf ")
		os.Exit(1)
	}
	filePath := os.Args[1]
	// We are only interested in people older than 18 years
	// So start date 2004/2/29
	// Also we want to include people from 18 to 60 years = 42 years
	// 365 days * 42years
	startDate := time.Date(2004, time.February, 29, 0, 0, 0, 0, time.UTC)
	allDays := 365 * 42
	for i := 0; i < allDays; i++ {
		go algorithm.GeneratePeselsForDate(startDate, channel)
		go algorithm.DecryptPdf(filePath, channel, &wg)
		startDate = startDate.AddDate(0, 0, -1)
	}
	wg.Wait()
}
