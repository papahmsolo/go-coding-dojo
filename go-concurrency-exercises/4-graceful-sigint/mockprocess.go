//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"fmt"
	"log"
	"time"
)

// MockProcess for example
type MockProcess struct {
	isRunning bool
}

// Run will start the process
func (m *MockProcess) Run() {
	m.isRunning = true

	fmt.Print("Process running..")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Process finished")
	m.isRunning = false
}

// Stop tries to gracefully stop the process, in this mock example
// this will not succeed
func (m *MockProcess) Stop() {
	if !m.isRunning {
		log.Fatal("Cannot stop a process which is not running")
	}

	fmt.Println("Stopping process..")
	for i := 0; i < 2; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	m.isRunning = false
	fmt.Println("\nStopping stopped")
}
