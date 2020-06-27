package main

import (
	"encoding/json"
	ri "github.com/sp0x/mediareleaseinfo"
	"log"
	"os"
)

func main() {
	result, err := ri.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(result)
}
