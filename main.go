package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Println("from init fail container.........")

	time.Sleep(41 * time.Second)
	os.Exit(1)
}
