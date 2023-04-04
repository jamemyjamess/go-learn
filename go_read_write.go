package main

import (
	"context"
	"errors"
	"log"
	"math"
	"math/rand"
	"sync"
)

type DataWriteExcel struct {
	sec    int
	offset int
	limit  int
	data   []interface{}
}

func main2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // WithCancel will be retained in memory indefinitely (until the program shuts down), causing a memory leak. Make sure it's called to release resources even if no errors
	done := make(chan bool)
	//err := error(nil)
	errFecthCh := make(chan error)
	errWriteExcelCh := make(chan error)
	//quiteFetch := make(chan error)
	dataWriteExcelCh := make(chan DataWriteExcel)
	requestStatus := 2
	limitPerList := 5
	currentRowExcel := 2
	// sec := int(math.Ceil(float64(amountTyreRequest) / float64(limitPerList)))
	// log.Println("sec", sec)
	optFecth := optionFetch{ctx: ctx, done: done, errCh: errFecthCh, limitPerList: limitPerList, requestStatus: requestStatus, dataWriteExcelCh: dataWriteExcelCh}
	go fetch(optFecth)
	optWriteExcel := optionWriteExcel{ctx: ctx, cancelCtx: cancel, done: done, errCh: errWriteExcelCh, currentRowExcel: currentRowExcel, dataWriteExcelCh: dataWriteExcelCh}
	go writeExcel(optWriteExcel)

	//isDone := <-done
	log.Println("done")

	// if err := catchError(isDone, errFecthCh, errWriteExcelCh); err != nil {
	// 	log.Println("err:", err.Error())
	// } else {
	// 	log.Println("no error")
	// }
	log.Println("End")

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // WithCancel will be retained in memory indefinitely (until the program shuts down), causing a memory leak. Make sure it's called to release resources even if no errors
	dataWriteExcelCh := make(chan DataWriteExcel)
	ishandleFetchDone, errHandleFetch := handleFetch(ctx, cancel, dataWriteExcelCh)
	ishandleWriteExcelDone, errhandleWriteExcel := handleWriteExcel(ctx, cancel, dataWriteExcelCh)

	if err := catchError(ishandleFetchDone, ishandleWriteExcelDone, errHandleFetch, errhandleWriteExcel); err != nil {
		log.Println("err:", err.Error())
	} else {
		log.Println("no error")
	}
	log.Println("End")

}

func catchError(ishandleFetchDone, ishandleWriteExcelDone chan bool, errFecthCh, errWriteExcelCh <-chan error) error {
	//ok := map[string]bool{"errFecthCh": false, "errWriteExcelCh": false}
	//err := error(nil)
	isHandleFetchDoneBool, ishandleWriteExcelDoneBool := false, false
	for {
		select {
		case err, ok := <-errFecthCh:
			log.Println("<-errFecthCh")
			if err != nil || !ok {
				return err
			}
		case err, ok := <-errWriteExcelCh:
			log.Println("<-errWriteExcelCh")
			if err != nil || !ok {
				return err
			}
		case <-ishandleFetchDone:
			isHandleFetchDoneBool = true
		case <-ishandleWriteExcelDone:
			ishandleWriteExcelDoneBool = true
		default:
			if isHandleFetchDoneBool && ishandleWriteExcelDoneBool {
				log.Println("NO Caugh error")
				return nil
			}
		}
	}
}

func handleFetch(ctx context.Context, cancelCtx context.CancelFunc, dataWriteExcelCh chan<- DataWriteExcel) (chan bool, <-chan error) {
	done := make(chan bool)
	errCh := make(chan error)
	requestStatus := 2
	limitPerList := 5
	optFecth := optionFetch{ctx: ctx, cancelCtx: cancelCtx, done: done, errCh: errCh, limitPerList: limitPerList, requestStatus: requestStatus, dataWriteExcelCh: dataWriteExcelCh}
	go fetch(optFecth)

	return done, errCh
}

func handleWriteExcel(ctx context.Context, cancelCtx context.CancelFunc, dataWriteExcelCh chan DataWriteExcel) (chan bool, <-chan error) {
	done := make(chan bool)
	errCh := make(chan error)
	currentRowExcel := 2
	optWriteExcel := optionWriteExcel{ctx: ctx, cancelCtx: cancelCtx, done: done, errCh: errCh, currentRowExcel: currentRowExcel, dataWriteExcelCh: dataWriteExcelCh}
	go writeExcel(optWriteExcel)

	return done, errCh
}

type optionFetch struct {
	ctx              context.Context
	cancelCtx        context.CancelFunc
	done             chan<- bool
	errCh            chan<- error
	limitPerList     int
	requestStatus    int
	dataWriteExcelCh chan<- DataWriteExcel
}

func fetch(opt optionFetch) {
	// จำนวนสัดส่วนทั้งหมด
	var wg sync.WaitGroup
	amountTyreRequest := 10
	sec := int(math.Ceil(float64(amountTyreRequest) / float64(opt.limitPerList)))
	//fetchDone := make(chan bool)
	for i := 0; i < sec; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case <-opt.ctx.Done():
				close(opt.dataWriteExcelCh)
				return // Error somewhere, terminate
			default: // Default is must to avoid blocking
				offset := i * opt.limitPerList
				limit := opt.limitPerList
				row := make([]interface{}, 50)
				for colID := 0; colID < 50; colID++ {
					row[colID] = rand.Intn(640000)
					if colID == 99 {
						opt.done <- true
						opt.errCh <- errors.New("func fetch Something went wrong.")
						//close(opt.dataWriteExcelCh)
						return
					}
				}
				log.Println("fetch i:", i)
				opt.dataWriteExcelCh <- DataWriteExcel{sec: i, offset: offset, limit: limit, data: row}
			}
		}(i)
	}
	wg.Wait()
	log.Println("wait done")
	log.Println("close(opt.errCh) from fetch")
	log.Println("close(opt.dataWriteExcelCh) from fetch")
	opt.done <- true
	opt.errCh <- nil
	close(opt.done)
	close(opt.errCh)
	close(opt.dataWriteExcelCh)
}

type optionWriteExcel struct {
	ctx              context.Context
	cancelCtx        context.CancelFunc
	currentRowExcel  int
	done             chan bool
	errCh            chan<- error
	dataWriteExcelCh chan DataWriteExcel
}

func writeExcel(opt optionWriteExcel) {
	for {
		select {
		case <-opt.ctx.Done():
			return // Error somewhere, terminate
		case dataWriteExcel, ok := <-opt.dataWriteExcelCh:
			if !ok { //if dataWriteExcel is closed
				log.Println("end writeExcel")
				opt.done <- true
				close(opt.done)
				return
			}
			rowNo := opt.currentRowExcel + dataWriteExcel.offset
			log.Println("sec:", dataWriteExcel.sec, "rowNo:", rowNo, "offset:", dataWriteExcel.offset, "limit:", dataWriteExcel.limit)
			if dataWriteExcel.sec == 0 {
				log.Println("error in write sec:", dataWriteExcel.sec)
				opt.errCh <- errors.New("func writeExcel Something went wrong.")
				opt.cancelCtx()
				close(opt.errCh)
				return
			}
		}
	}
}
