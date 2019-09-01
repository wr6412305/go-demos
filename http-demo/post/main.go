package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postJSON() {
	url := "http://baidu.com"
	usrID := "liangjisheng"
	pwd := "123456"
	//json序列化
	post := "{\"UserId\":\"" + usrID + "\",\"Password\":\"" + pwd + "\"}"
	fmt.Println(url, "post", post)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(post)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func postString() {
	url := "http://localhost:9090/page"
	payload := strings.NewReader("username=ljs&usertext=teacher")
	// post string
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func postBytes() {
	requestURL := "http://localhost:9090/page"
	// 要post的参数
	form := url.Values{
		"username": {"ljs"},
		"usertext": {"teacher"},
	}
	// func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) {
	// 对form进行编码
	fmt.Println(form.Encode())
	body := bytes.NewBufferString(form.Encode())
	// post bytes数据给URL
	resp, err := http.Post(requestURL, "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}

func postForm() {
	requestURL := "http://localhost:9090/page"
	resp, err := http.PostForm(requestURL, url.Values{"username": {"ljs"}, "usertext": {"teacher"}})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	// postJSON()
	// postString()
	// postBytes()
	postForm()
}
