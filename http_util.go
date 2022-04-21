package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http访问工具，访问url获取http响应数据内容

func postUrl(url string) (result string) {

	//
	return
}

//uri string
func getUrl(url string) (result string) {

	//
	// r, err := http.Get("https://api.github.com/events")
	r, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)

	return
}
