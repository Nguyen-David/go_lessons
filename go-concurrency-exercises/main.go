
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
	"go-concurrency-exercises/mockstream"
	"time"
)

func producer(stream mockstream.Stream, tweetChan chan *mockstream.Tweet) {
	defer close(tweetChan)
	for {
		tweet, err := stream.Next()
		if err == mockstream.ErrEOF {
			return
		}
		tweetChan <- tweet
	}
}

func consumer(tweetChan chan *mockstream.Tweet) {
	for {
		select {
		case tweet, ok := <-tweetChan:
			if ok {
				if tweet.IsTalkingAboutGo() {
					fmt.Println(tweet.Username, "\ttweets about golang")
				} else {
					fmt.Println(tweet.Username, "\tdoes not tweet about golang")
				}
			} else {
				return
			}
		}
	}
}

func main() {
	start := time.Now()
	stream := mockstream.GetMockStream()

	// Channels
	tweetChan := make(chan *mockstream.Tweet)

	// Producer
	go producer(stream, tweetChan)

	// Consumer
	consumer(tweetChan)

	fmt.Printf("Process took %s\n", time.Since(start))
}