package main

import (
	"context"
	"fmt"
)

func main() {

	// Create a context to allow cancellation of all workers on error.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a channel to send tasks amount 10 to workers.
	taskCh := make(chan string, 10)
	// Submit some tasks to the worker pool.
	for i := 0; i < 10; i++ {
		task := fmt.Sprintf("task-%d", i)
		taskCh <- task
	}
	// Close the task channel to indicate no more tasks.
	close(taskCh)

	// Number of workers in the pool
	numWorkers := 8
	// Create a channel to communicate errors from workers.
	catchWorkerErrorCh := make(chan error)
	// Start the worker pool follow by numWorkers.
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, taskCh, catchWorkerErrorCh)
	}

	// Check for any errors from workers.
	err := <-catchWorkerErrorCh
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// Close the error channel.
	//close(errorCh)

	fmt.Println("All tasks completed!")
}

// worker is a goroutine that performs tasks.
func worker(ctx context.Context, taskCh <-chan string, catchWorkerErrorCh chan<- error) {
	for {
		select {
		case <-ctx.Done():
			// Context canceled, worker should stop.
			return
		case task, more := <-taskCh:
			if !more {
				// No more tasks, worker should stop.
				return
			}
			// Perform the task.
			if err := doTask(task); err != nil {
				// Send the error to the error channel.
				catchWorkerErrorCh <- err
				//close(catchWorkerErrorCh)
				return
			}
		}
	}
}

// performTask simulates performing a task and returns an error if needed.
func doTask(task string) error {
	// Simulate work by printing the task.
	fmt.Println("Doing work on task:", task)

	// Simulate an error condition.
	if task == "task-3" {
		return fmt.Errorf("error: task-3 encountered an error")
	}

	return nil
}
