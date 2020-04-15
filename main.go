package main

import (
	"fmt"
	"time"
)

// welcome information
func welcome() {
	fmt.Println("Welcome")
	fmt.Println("Please choose the function you want")
	fmt.Println("1.Get the date after several days\n2.How many days between two dates")
	fmt.Println("Just type the number and press Enter")
}

// function, the date after several days
func dayAfter() {
	var (
		inputTime string
		spend     time.Duration
	)
	fmt.Println("You want to get the date after several days\nyou need to give a start date")
	fmt.Println("Please input the start date and press Enter\nFormat sample：2006/01/02 means January 2nd 2006")
	fmt.Println("If start date is today，just input now")
	fmt.Scanln(&inputTime)
	switch inputTime {
	case "now":
		inputTime = time.Now().Format("2006/01/02")
	}
	startTime, err := time.Parse("2006/01/02", inputTime)
	if err != nil {
		fmt.Println("Can not parse，try again")
	}
	fmt.Println("Input days between two dates")
	fmt.Scanln(&spend)
	endTime := startTime.Add(24 * time.Hour * spend)
	fmt.Printf("The date %d days after %s is:\n", spend, inputTime)
	fmt.Println(endTime.Format("2006/01/02 Mon"))
}

// How many days between two dates
func dayToDay() {
	var (
		startTime string
		endTime   string
	)
	fmt.Println("You want to know how many days between two dates")
	fmt.Println("Please input the start date and press Enter\nFormat sample：2006/01/02 means January 2nd 2006")
	fmt.Println("If start date is today，just input now")
	fmt.Scanln(&startTime)
	switch startTime {
	case "now":
		startTime = time.Now().Format("2006/01/02")
	}
	startTimeObj, err := time.Parse("2006/01/02", startTime)
	if err != nil {
		fmt.Println("Can not parse，try again")
		return
	}
	fmt.Println("Please input the end date and press Enter\nFormat sample：2006/01/02 means January 2nd 2006")
	fmt.Println("If the date is today just input now")
	fmt.Scanln(&endTime)
	switch endTime {
	case "now":
		endTime = time.Now().Format("2006/01/02")
	}
	endTimeObj, err := time.Parse("2006/01/02", endTime)
	if err != nil {
		fmt.Println("Can not parse，try again")
		return
	}
	spend := endTimeObj.Sub(startTimeObj)
	spendHour := spend.Hours()
	spendDay := spendHour / 24
	fmt.Printf("It's %d days between %s to %s\n", int(spendDay), startTime, endTime)
}

func main() {
	// looping
	for {
		// welcome infomation
		welcome()
		// choose the function
		var choice1 int
		fmt.Scanln(&choice1)
		switch choice1 {
		case 1:
			// The date after several days
			dayAfter()
		case 2:
			// Days between two dates
			dayToDay()
		default:
			fmt.Println("Can not parse，try again")
			continue
		}
		// continue or exit
		var end string
		fmt.Println("type c to continue，type something else or press Enter to exit")
		fmt.Scanln(&end)
		switch end {
		case "c", "C":
			// continue looping
			continue
		default:
			// break the loop, then exit
			goto Swap
		}
	}
Swap:
	fmt.Println("Thanks for using，bye")
}
