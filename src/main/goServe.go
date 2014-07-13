package main

import (
	"server"
	"runtime"
	"net/http"
	"fmt"
)

func main() {
	var exit = make(chan bool)
	go func() {
		runtime.GOMAXPROCS(10)
		var done = make(chan bool)
		server.Start()
		<-done
	}()
	go shutDownServer(exit)

	<-exit

	fmt.Print("Service ShutDown\n")

}


func shutDownServer(exit chan bool) {
	http.HandleFunc("/exitSrv", func(w http.ResponseWriter, r *http.Request) {
			exit <- true
		})

	if err := http.ListenAndServe(":8088", nil); err != nil {
		fmt.Printf(err.Error())
	}
}
