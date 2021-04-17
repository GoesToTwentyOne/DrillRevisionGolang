package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "NIHAD "
	format := `
	<!DOCTYPE html>
<html>
<head>
<title>Page Title</title>
</head>
<body>

<h1>` + name + `</h1>

</body>
</html>
	`

	msg := fmt.Sprintf(format)
	nFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	io.Copy(nFile, strings.NewReader(msg))

}
