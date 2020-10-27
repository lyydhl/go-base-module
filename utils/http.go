package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// get 请求
func GetRequest(url string) ([]byte, error) {
	get, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer get.Body.Close()

	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

// get 请求
func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")

	get, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer get.Body.Close()

	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}

	return all, nil
}

// post 请求
func PostRequest(url, parameter string) ([]byte, error) {
	client := &http.Client{}
	byteParameter := bytes.NewBuffer([]byte(parameter))
	request, _ := http.NewRequest("POST", url, byteParameter)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		return nil, errors.New("网络请求失败")
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("读取网络内容失败")
	}
	return all, nil
}
