package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

func main() {
	http.HandleFunc("/", echoHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// 读取请求的URL、Header和Body
	url := r.URL.String()
	header := r.Header
	body, _ := ioutil.ReadAll(r.Body)

	// 将Header的key按字母顺序排序
	keys := make([]string, 0, len(header))
	for key := range header {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 将URL、Header和Body写入响应的Body中
	fmt.Fprintf(w, "URL: %s\n\n", url)
	fmt.Fprintf(w, "Header:\n")
	for _, key := range keys {
		values := header[key]
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", key, value)
		}
	}
	fmt.Fprintf(w, "\nBody:\n%s", body)
}
