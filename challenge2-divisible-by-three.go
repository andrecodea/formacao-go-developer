// Challenge 2: Numbers Divisible by 3 Detector
// Detects and displays numbers divisible by 3 within a given range

package main

import "fmt"

// checkDivisibilityByThree analyzes numbers from 1 to limit
// and identifies which ones are divisible by 3
func checkDivisibilityByThree(limit int) {
    fmt.Printf("Checking numbers from 1 to %d for divisibility by 3:\n\n", limit)
    
    for i := 1; i <= limit; i++ {
        // Use modulo operator to check divisibility
        // If remainder is 0, number is divisible by 3
        if i%3 == 0 {
            fmt.Printf("✓ %d is divisible by 3\n", i)
        } else {
            fmt.Printf("  %d\n", i)
        }
    }
}

// main function demonstrates the divisibility checker
func main() {
    // Test with numbers from 1 to 100
    checkDivisibilityByThree(100)
}
