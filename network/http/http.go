package tyhttp

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func DoHttpPost(url string, headers map[string]string, data string, timeout time.Duration) (int, []byte, error) {
	if url == "" {
		return 0, nil, errors.New("url is empty")
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return doHttp(req, timeout)
}

func DoHttpGet(url string, headers map[string]string, timeout time.Duration) (int, []byte, error) {
	if url == "" {
		return 0, nil, errors.New("url is empty")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return doHttp(req, timeout)
}

func DoHttpsGet(url string, headers map[string]string, timeout time.Duration) (int, []byte, error) {
	if url == "" {
		return 0, nil, errors.New("url is empty")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return doHttps(req, timeout)
}

func DoHttpsPost(url string, headers map[string]string, data string, timeout time.Duration) (int, []byte, error) {
	if url == "" {
		return 0, nil, errors.New("url is empty")
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return doHttps(req, timeout)
}

func doHttp(req *http.Request, timeout time.Duration) (int, []byte, error) {
	c := &http.Client{Timeout: timeout}
	resp, err := c.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}

func doHttps(req *http.Request, timeout time.Duration) (int, []byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{tr, nil, nil, timeout}
	resp, err := c.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}
