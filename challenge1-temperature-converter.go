// Challenge 1: Thermometric Converter
// Converts temperature from Kelvin to Celsius and identifies water phase transitions

package main

import "fmt"

// ConvertKelvinToCelsius converts temperature from Kelvin to Celsius
// and provides additional context for water's phase transition points
func ConvertKelvinToCelsius(kelvin float64) (float64, string) {
    // Convert using the standard formula: C = K - 273.15
    celsius := kelvin - 273.15
    
    var message string
    
    // Check for water's critical temperatures
    switch celsius {
    case 100.0:
        message = "This is the boiling point of water!"
    case 0.0:
        message = "This is the freezing point of water!"
    default:
        message = ""
    }
    
    return celsius, message
}

// main function handles user input and displays the conversion results
func main() {
    var kelvin float64

    // Prompt user for temperature input
    fmt.Print("Enter the temperature in Kelvin (K): ")
    
    // Read user input
    fmt.Scanf("%v", &kelvin)
    
    // Perform conversion and get additional information
    celsius, message := ConvertKelvinToCelsius(kelvin)
    
    // Display the converted temperature with proper formatting
    fmt.Printf("The equivalent temperature in Celsius is: %.2f°C\n", celsius)
    
    // Display additional information if available
    if message != "" {
        fmt.Println(message)
    }
}
