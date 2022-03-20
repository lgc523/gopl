package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/net/context"
)

func doApiCall(ctx context.Context, wg *sync.WaitGroup) error {

	defer wg.Done()

	req, err := http.NewRequest("GET", "https://httpstat.us/200", nil)

	if err != nil {
		return err
	}

	req.URL.Query()
	sleepMin := 1000
	sleepMax := 5000

	q.Set("sleep", fmt.Sprintf("%d", rand.Intn(sleepMax-sleepMin)+sleepMin))
	req.URL.RawQuery = q.Encode()

	c := make(chan error, 1)
	go func() {
		_, err := http.DefaultClient.Do(req)
		c <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-c:
		return err

	}
}

func main() {
	var (
		closing     = make(chan struct{})
		ticker      = time.NewTicker(1 * time.Second)
		logger      = log.New(os.Stderr, "", log.LstdFlags)
		batchSize   = 6
		wg          sync.WaitGroup
		ctx, cancel = context.WithCancel(context.Background())
	)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
		<-signals
		cancel()
		close(closing)
	}()
loop:
	for {
		select {
		case <-closing:
			break loop
		case <-ticker.C:
			for n := 0; n < batchSize; n++ {
				wg.Add(1)
				go doAPICall(ctx, &wg)
			}
			wg.Wait()
			logger.Printf("Completed doing %d things.", batchSize)
		}
	}

}
func init() {
	rand.Seed(time.Now().Unix())
}
