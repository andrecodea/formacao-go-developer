// Challenge 3: PinPan
// Custom FizzBuzz implementation using for loop and modulo operator

package main

import "fmt"

// pinPan implements the complete PinPan logic
func pinPan(limit int) {
    for i := 1; i <= limit; i++ {
        // Important: Check divisibility by both 3 and 5 first
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("PinPan")
        } else if i%3 == 0 {
            fmt.Println("Pin")
        } else if i%5 == 0 {
            fmt.Println("Pan")
        } else {
            fmt.Println(i)
        }
    }
}

func main() {
    pinPan(100)
}
