package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	request, err := http.NewRequest("GET", "8080", nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("error in Do():", err)
		return
	}
	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)

	fmt.Println("Status code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("character set:", characterSet[1])
	}
	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength", httpData.ContentLength)
	}
}
