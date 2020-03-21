package main

import (
	"log"
	"os"
	"stock/src/controllers"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	controllers.FetchMainIndicator()
}
