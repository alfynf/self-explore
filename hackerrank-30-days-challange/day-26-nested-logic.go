/*
Day 28: Nested Logic

Task: Your local library needs your help! Given the expected and actual return dates for a library book, create a program that calculates the fine (if any). 
The fee structure is as follows:
	If the book is returned on or before the expected return date, no fine will be charged (i.e.: fine = 0).
	If the book is returned after the expected return day but still within the same calendar month and year as the expected return date,
	fine = 15 hackos x (number days days late).
	If the book is returned after the expected return month but still within the same calendar year as the expected return date, the
	fine = 500 hackos x (number of month late).
	If the book is returned after the calendar year in which it was expected, there is a fixed fine of 10000 hackos.
	*/

package main
import (
    "fmt"
    "bufio"
    "os"   
    "strings"
    "strconv"
)


func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    input := bufio.NewScanner(os.Stdin)
    var day []int
    var month []int
    var year []int
    var count int
    for count < 2 && input.Scan() {
       inputString := strings.Split(input.Text()," ")
       inputDay, _ := strconv.Atoi(inputString[0])
       day = append(day, inputDay)
       inputMonth, _ := strconv.Atoi(inputString[1])
       month = append(month, inputMonth)
       inputYear, _ := strconv.Atoi(inputString[2])
       year = append(year,inputYear)
       count++
    }

    if year[0] > year[1] { 
        fmt.Println(10000)
    } else if year[0] == year[1] {
        if month[0] > month[1] {
        fmt.Println(500*(month[0]-month[1]))
        } else if month[0] == month[1] {
           if day[0] > day[1] {
           fmt.Println(15*(day[0]-day[1]))
           } else {
               fmt.Println(0)
           }
        } else {
            fmt.Println(0)
        }
    } else {
        fmt.Println(0)
    }
}