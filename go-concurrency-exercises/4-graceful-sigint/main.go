//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	proc := MockProcess{}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	done := make(chan bool)
	stopping := false

	go func() {
		proc.Run()
		done <- true
	}()

	for {
		select {
		case <-c:
			if stopping {
				fmt.Println("process aborted")
				os.Exit(1)
			}
			stopping = true
			go func() {
				proc.Stop()
				done <- false
			}()
		case success := <-done:
			if success {
				fmt.Println("process finished successfully")
			} else {
				fmt.Println("process stopped successfully")
			}
			return
		}
	}
}
