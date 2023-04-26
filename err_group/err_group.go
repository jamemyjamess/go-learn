package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Test struct {
	no    int
	value int
}

var m sync.Mutex

func main() {
	errorGroupWithTimeOut()
}

func errorGroup() {
	ctx := context.Background()
	testList := &[]Test{}
	rows := 100000
	limit := 2000
	section := int(math.Ceil(float64(rows) / float64(limit)))
	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < section; i++ {
		// log.Println("start section:", i)
		no := i
		offset := no * limit
		fun := func() error {
			return doWorkErrorGroup(ctx, testList, no, limit, offset)
		}
		g.Go(fun)
	}
	// log.Println("pending g.Wait()")
	if err := g.Wait(); err != nil {
		log.Println(err.Error())
	}
	//time.Sleep(10 * time.Second)
	fmt.Println("total test: ", len(*testList))
}

func doWorkErrorGroup(ctx context.Context, testList *[]Test, no, limit, offset int) error {
	err := error(nil)
	testListItem := &[]Test{}
	for i := 0; i < limit; i++ {
		testItem := &Test{}
		testItem.no = i + no
		testItem.value = rand.Intn(100)
		*testListItem = append(*testListItem, *testItem)
	}
	m.Lock()
	*testList = append(*testList, *testListItem...)
	m.Unlock()
	return err
}

func errorGroupWithTimeOut() {
	testList := &[]Test{}
	rows := 100000
	limit := 2000
	section := int(math.Ceil(float64(rows) / float64(limit)))
	ctx, cancel := context.WithTimeout(context.Background(), 100000000*time.Second)
	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	// prevent memory leak
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i < section; i++ {
		no := i
		offset := no * limit
		fn := func() error {
			return doWorkErrorGroupWithTimeOut(ctx, testList, no, limit, offset)
		}
		g.Go(fn)
	}
	if err := g.Wait(); err != nil {
		log.Println("g.Wait() err:", err.Error())
	}
	log.Println("total test: ", len(*testList))

}

func doWorkErrorGroupWithTimeOut(ctx context.Context, testList *[]Test, no, limit, offset int) error {
	done := make(chan struct{}, 1)
	errCh := make(chan error, 1)
	//time.Sleep(5 * time.Second)
	go func() {
		testListItem := &[]Test{}
		// example error
		if errCh == nil {
			errCh <- fmt.Errorf("example error on work no: %v", no)
		}
		for i := 0; i < limit; i++ {
			testItem := &Test{}
			testItem.no = i + no
			testItem.value = rand.Intn(100)
			*testListItem = append(*testListItem, *testItem)
		}
		m.Lock()
		*testList = append(*testList, *testListItem...)
		m.Unlock()
		close(done) //completed
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	case err := <-errCh:
		log.Println(err.Error())
		close(errCh)
		return err
	}
}
