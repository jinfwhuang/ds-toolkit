package sia

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiPassword = "dac05fa5b012cdb7a52e57a97199d757"
	siaURL      = "http://localhost:9980/"
)

// Should work, not fully tested though, upload stuck at 30%
func Write(fileDir string, targetDir string) error {
	client := &http.Client{}

	params := url.Values{}
	params.Add("source", fileDir)
	reqBody := strings.NewReader(params.Encode())
	uploadURL := fmt.Sprintf("%vrenter/upload/%v", siaURL, targetDir)
	req, err := http.NewRequest("POST", uploadURL, reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Sia-Agent")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("", apiPassword)

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func Read(siaDir string) (string, error) {
	client := &http.Client{}

	downloadURL := fmt.Sprintf("%vrenter/dir/%v", siaURL, siaDir)
	req, err := http.NewRequest("GET", downloadURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Sia-Agent")
	req.SetBasicAuth("", apiPassword)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
