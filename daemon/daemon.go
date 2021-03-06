package daemon

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var cmd *exec.Cmd

/* Start start child process
   exit: child process allow exit count, if <=0 not limit,
         otherwise if child process exit more than exit count,
         the parent process will exit
   pid: pid file save path
   username: change child process by user only supported linux, allow empty
   arg: child process startup arguments
*/
func Start(exit int, pid, username string, arg ...string) {
	chExit := make(chan struct{})
	onExit := false
	var wg sync.WaitGroup
	if len(pid) > 0 {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		wg.Add(1)
		go func() {
			defer wg.Done()
			sig := <-c
			onExit = true
			cmd.Process.Signal(sig)
			<-chExit
			os.Remove(pid)
			os.Exit(0)
		}()
	}
	var cnt int
	for {
		begin := time.Now()
		run(chExit, pid, username, arg...)
		if onExit {
			break
		}
		if time.Since(begin).Minutes() > 10 {
			cnt = 0
		}
		cnt++
		if exit > 0 && cnt > exit {
			break
		}
	}
	wg.Wait()
}

func run(ch chan struct{}, pid, username string, arg ...string) {
	cmd = makeCommand(username, arg...)
	if err := cmd.Start(); err == nil {
		writePidFile(pid, os.Getpid())
		cmd.Wait()
		select {
		case ch <- struct{}{}:
		default:
		}
	} else {
		fmt.Println("create child process failed")
		os.Exit(1)
	}
}

func writePidFile(dir string, pid int) {
	ioutil.WriteFile(dir, []byte(fmt.Sprintf("%d", pid)), 0644)
}
