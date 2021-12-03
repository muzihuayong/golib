package libhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpRequest(method, url string, header map[string]string, data interface{}, timeout time.Duration) ([]byte, error) {
	// json数据
	buf := bytes.NewBuffer(nil)
	if data != nil {
		encoder := json.NewEncoder(buf)
		if err := encoder.Encode(data); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func Post(url string, header map[string]string, data interface{}, timeout time.Duration) ([]byte, error) {
	return HttpRequest(http.MethodPost, url, header, data, timeout)
}

func Get(url string, header map[string]string, data interface{}, timeout time.Duration) ([]byte, error) {
	return HttpRequest(http.MethodGet, url, header, data, timeout)
}

func PostSimple(host string, params interface{}) (string, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(host, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func GetSimple(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
