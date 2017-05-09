package main

import (
	"log"
	"os"

	"github.com/yuuki/cookwares/pkg/resource/aptpackage"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("Starting to cook localhost...")

	if err := aptpackage.Install("nginx", ""); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
