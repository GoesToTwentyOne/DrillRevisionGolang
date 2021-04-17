package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	res, _ := http.Get("https://github.com/GoesToEleven/GolangTraining/blob/master/21_interfaces/01_interface/05_io-copy/01_no-error-checking/main.go")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}
