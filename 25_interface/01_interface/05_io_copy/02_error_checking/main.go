package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "My name is Nihad \n"
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)
	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	resp, err := http.Get("https://github.com/GoesToTwentyOne")
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Println(resp)
	io.Copy(os.Stdout, resp.Body)

}
