package main

import (
	"bargainhunter"
	"fmt"
	"os"
)

func main() {
	pagecontent, err := bargainhunter.Fetch("https://www.amazon.com/Ubiquiti-Security-Gateway-USG-PRO-4-Renewed/dp/B07MC83QR4/")
	if err != nil {
		fmt.Printf("Something went wrong with Fetch: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(pagecontent[:100], "...")
}
