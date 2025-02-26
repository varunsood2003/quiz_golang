package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Printf("Problem with opening the file")
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	var result int
	timeLimit := 10 * time.Second
	timer := time.NewTimer(timeLimit)
	for i, rec := range records {
		fmt.Printf("Question %v is %s \n", i+1, rec[0])
		answerChan := make(chan string, 1)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			answerChan <- ans
			fmt.Printf("You answered %s \n", ans)
		}()
		select {
		case userAnswer := <-answerChan:
			if userAnswer == rec[1] {
				result++
				fmt.Printf("Answer is correct %s \n", rec[1])
			}
		case <-timer.C:
			fmt.Printf("Time is up!, Your score is %v \n", result)
			return
		}
	}
	fmt.Printf("Result is %v", result)
}
