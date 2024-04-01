//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(done chan struct{}, tweetChan chan Tweet, stream Stream) {
	defer close(tweetChan)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}

		tweetChan <- *tweet
	}
	done <- struct{}{}
}

func consumer(tweetChan chan Tweet) {
	for {
		t, ok := <-tweetChan
		if !ok {
			return
		}
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	done := make(chan struct{})
	tweetChan := make(chan Tweet)
	// Producer
	go producer(done, tweetChan, stream)

	// Consumer
	go consumer(tweetChan)

	<-done

	fmt.Printf("Process took %s\n", time.Since(start))
}
