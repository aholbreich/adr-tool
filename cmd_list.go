package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

// CLI Command
type ListCmd struct {
}

// Command Handler
func (r *ListCmd) Run() error {

	entries, err := os.ReadDir(configFolderPath)

	if err != nil {
		log.Fatal(err)
	}

	var adrs []string

	for _, e := range entries {
		if unicode.IsDigit(rune(e.Name()[0])) { //Starst with digit? Must be adr file
			adrs = append(adrs, e.Name())
		}
	}

	reverse(adrs)

	for _, adr := range adrs {
		fmt.Println(adr)
	}

	return nil

}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
