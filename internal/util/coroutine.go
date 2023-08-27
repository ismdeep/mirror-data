package util

import (
	"errors"
	"sync"
)

// run coroutines
func CoroutineRun(size int, f func() error) error {
	var errLst []error
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			if err := f(); err != nil {
				errLst = append(errLst, err)
			}
		}()
	}

	wg.Wait()
	return errors.Join(errLst...)
}
