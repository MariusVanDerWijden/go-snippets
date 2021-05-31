package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	done := make(chan error, 1)

	for i := 0; i < 100; i++ {
		bin := "/home/matematik/go/src/github.com/ethereum/go-ethereum/build/bin/geth"
		cmd := exec.Command(bin, "--goerli")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			panic(err)
		}
		// Sleep 5 minutes
		time.Sleep(5 * time.Minute)
		// Send kill signal
		if err := cmd.Process.Signal(os.Interrupt); err != nil {
			fmt.Printf("killing: %v\n", err)
		}
		fmt.Println("Killed")
		time.Sleep(10 * time.Second)
		go func() {
			done <- cmd.Wait()
		}()
		select {
		case err := <-done:
			fmt.Printf("Killed, err %v\n", err)
		case <-time.After(2 * time.Minute):
			panic("Shutdown took longer than 2 minutes")
		}
	}

}
