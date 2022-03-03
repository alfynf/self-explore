/*
Day 25: Running Time and Complexity

A prime is a natural number greater than 1 that has no positive divisors other than 1 and itself. 
Given a number, n, determine and print whether it is or Prime or Not Prime.

Note: If possible, try to come up with a O^1/2n primality algorithm, or see what sort of optimizations you come up with for an 
O^1/2n algorithm. Be sure to check out the Editorial after submitting your code.
*/

package main
import (
    "fmt"
    "math"
    "bufio"
    "os"
    "strconv"
)

func isPrime(num int)string{
	if num == 1 {
        return "Not prime"
    }
	
    for i:=2;i<= int(math.Sqrt(float64(num)));i++{
        if num%i == 0 {
            return "Not prime"
        }
    }
    return "Prime"
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 var inputCount int
 fmt.Scanln(&inputCount)
 input := bufio.NewScanner(os.Stdin)
 for input.Scan(){
     inputInt,_ := strconv.Atoi(input.Text())
     fmt.Println(isPrime(inputInt))
 }
}

