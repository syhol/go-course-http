package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Could not reach http://google.com")
		os.Exit(1)
	}

	read4Bytes4Times(resp)
	fmt.Println()
	read1024Bytes(resp)
	fmt.Println()
	copyEverythingToStdout(resp)
	fmt.Println()

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func copyEverythingToStdout(resp *http.Response) {
	io.Copy(os.Stdout, resp.Body)
}

func read1024Bytes(resp *http.Response) {
	body := make([]byte, 1024)
	resp.Body.Read(body)
	fmt.Println(string(body))
}

func read4Bytes4Times(resp *http.Response) {
	body := []byte{0, 0, 0, 0}
	resp.Body.Read(body)
	fmt.Println(string(body))
	resp.Body.Read(body)
	fmt.Println(string(body))
	resp.Body.Read(body)
	fmt.Println(string(body))
	resp.Body.Read(body)
	fmt.Println(string(body))
}
