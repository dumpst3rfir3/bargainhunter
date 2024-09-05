package main

import (
	"bargainhunter"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s PRODUCT_PAGE_URL\n", os.Args[0])
		os.Exit(1)
	}
	url := os.Args[1]
	pagecontent, err := bargainhunter.Fetch(url)
	if err != nil {
		fmt.Printf("Something went wrong with Fetch: %s\n", err)
		os.Exit(1)
	}
	price, err := bargainhunter.ExtractPrice(pagecontent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Available now! Just $%.2f\n", price)
}
