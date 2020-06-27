package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sp0x/mediareleaseinfo"
)

func main() {
	result, err := mediareleaseinfo.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}
