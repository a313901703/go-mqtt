package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func HttpRequest(path string, method string, data map[string]interface{}, headers map[string]string) ([]byte, error) {
	var resp *http.Response
	var err error

	values := url.Values{}
	for k, v := range data {
		values.Set(k, v.(string))
	}

	if method == "GET" && data != nil {
		path += "?"
		path += values.Encode()
		resp, err = http.Get(path)
	} else if data != nil {
		resp, err = http.PostForm(path, values)
	}
	// fmt.Println(path, method, values)
	// req, err = http.NewRequest(method, path, ioReader)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println("body", string(body))

	return body, nil
}

func HttpJson(path string, method string, data map[string]interface{}, headers map[string]string) (map[string]interface{}, error) {
	var resp *http.Response
	var err error

	client := &http.Client{}
	var dataByte []byte
	var req *http.Request

	if method == "GET" && data != nil {
		path += "?"
		values := url.Values{}
		for k, v := range data {
			values.Set(k, v.(string))
		}
		path += values.Encode()
	} else {
		dataByte, _ = json.Marshal(data)
	}

	fmt.Println(path, method, string(dataByte))

	req, err = http.NewRequest(method, path, bytes.NewReader(dataByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret map[string]interface{}

	json.Unmarshal(body, &ret)
	fmt.Println("body", string(body), ret)
	return ret, nil
}
