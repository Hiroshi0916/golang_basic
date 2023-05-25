package main

import "fmt"

func generateVals() <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))

	go func() {
		defer close(chanStream)

		for i := 0; i < 10; i++ {
			stream := make(chan interface{}, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	valChan := make(chan interface{})

	go func() {
		defer close(valChan)
		for {
			select {
			case <-done:
				return

			case v, ok := <-c:
				if !ok {
					return
				}

				select {
				case valChan <- v:
				case <-done:
				}
			}
		}
	}()
	return valChan
}

func bridge(done <-chan interface{}, chanCh <-chan <-chan interface{}) <-chan interface{} {
	valCh := make(chan interface{})

	go func() {
		defer close(valCh)

		for {
			var ch <-chan interface{}
			select {
			case maybeCh, ok := <-chanCh:
				if !ok {
					return
				}
				ch = maybeCh
			case <-done:
				return
			}

			for v := range orDone(done, ch) {
				select {
				case valCh <- v:
				case <-done:
				}
			}
		}
	}()
	return valCh
}

func main() {
	done := make(chan interface{})

	for v := range bridge(done, generateVals()) {
		fmt.Println(v)
	}
}
