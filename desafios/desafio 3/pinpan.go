package main

import (
	"fmt"
)

func pinpan(limite int) {
	for i := 1; i <= limite; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	pinpan(100)
}
