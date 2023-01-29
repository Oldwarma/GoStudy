package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	myUrl string
	delay = 5
	w     sync.WaitGroup
)

type myData struct {
	r   *http.Response
	err error
}

func main() {
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(10)*time.Second)
	defer cancel()

	fmt.Printf("connecting to %s\n", myUrl)
	w.Add(1)
	go connect(c)
	w.Wait()
	fmt.Println("exiting...")
}
func connect(c context.Context) error {
	defer w.Done()
	data := make(chan myData, 1) //缓冲管道
	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}    //客户端
	req, _ := http.NewRequest("GET", myUrl, nil) //http请求
	go func() {
		response, err := httpClient.Do(req) //http回复
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()
	select {
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("The request was cancelled")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("error select", err)
			return err
		}
		defer resp.Body.Close()
		realHTTPData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error select:", err)
			return err
		}
		fmt.Printf("server Response :%s\n", realHTTPData)
	}
	return nil
}
