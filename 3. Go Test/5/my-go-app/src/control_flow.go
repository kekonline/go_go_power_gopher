// control_flow.go
package main

import "fmt"

// This function demonstrates control flow statements in Go.
func controlFlowDemo() {
    // If-else statement example
    number := 10
    if number > 0 {
        fmt.Println("The number is positive.")
    } else if number < 0 {
        fmt.Println("The number is negative.")
    } else {
        fmt.Println("The number is zero.")
    }

    // For loop example
    fmt.Println("Counting from 1 to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // Infinite loop example (commented out to prevent infinite execution)
    // for {
    //     fmt.Println("This will run forever!")
    // }

    // Switch statement example
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    default:
        fmt.Println("Weekend")
    }

    // Breaking and continuing loops example
    fmt.Println("Counting with break and continue:")
    for j := 1; j <= 10; j++ {
        if j == 5 {
            continue // Skip the number 5
        }
        if j == 8 {
            break // Stop the loop when reaching 8
        }
        fmt.Println(j)
    }
}