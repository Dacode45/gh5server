package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
	"io"
	"bitbucket.org/JeremySchlatter/go-atexit"
)
const(
	DATABASE_CACHE_FILE = "tickets.json"
)
func main() {
	atexit.TrapSignals()
	defer atexit.CallExitFuncs()

	atexit.Run(func(){
		serilized, err := json.Marshal(gTickets)
		var file io.WriteCloser
		if err != nil{
			panic(err)
		}
		if file, err = os.OpenFile(DATABASE_CACHE_FILE, os.O_WRONLY|os.O_CREATE, 0666); err != nil{
			panic(err)
		}
		defer file.Close()
		if _,err = file.Write(serilized); err != nil{
			panic(err)
		}
	})

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
