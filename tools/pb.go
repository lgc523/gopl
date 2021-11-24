package main

import (
	"github.com/cheggaaa/pb/v3"
	"log"
	"time"
)

const count = 100

func main() {
	bar := pb.New(count)
	bar.SetRefreshRate(time.Millisecond)
	bar.SetTemplateString("fuck")
	if bar.Err() != nil {
		log.Fatal(bar.Err().Error())
	}
	bar.Start()
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(55 * time.Millisecond)
	}
	bar.Finish()
}
