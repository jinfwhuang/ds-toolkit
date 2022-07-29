package arweave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

const (
	APIkey = "Bearer ESTe15ac716-261e-41f3-a62e-8ead4cfdbd7bARY"
	url    = "https://api.estuary.tech/content/add"
)

// Does not work!
// Getting error from estuary - {"error":{"code":404,"reason":"Not Found","details":"Not Found"}}.
// Getting the same error via curl and estuary's swagger (https://docs.estuary.tech/swagger-ui-page#/content/post_content_add).
// curl reference:
// curl \
// -X POST https://api.estuary.tech/content/add \
// > -H "Authorization: Bearer ESTe15ac716-261e-41f3-a62e-8ead4cfdbd7bARY" \
// > -H "Accept: application/json" \
// > -H "Content-Type: multipart/form-data" \
// > -F "data=@{FULL_FILE_PATH_DO_NOT_REMOVE_@}"
func Write(data []byte) (string, error) {
	client := &http.Client{}
	payloadDir := fileInRuntimeDir("/test_blob.json")
	b, _, err := createMultipartFormData("data", payloadDir)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", APIkey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "multipart/form-data;")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	println(resp.Status)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.String(), nil
}

// Not implemented as writing does not work.
func Read(id string) ([]byte, error) {
	return nil, nil
}

func fileInRuntimeDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename) + file
}

func createMultipartFormData(fieldName, fileName string) (bytes.Buffer, *multipart.Writer, error) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file := mustOpen(fileName)
	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		return bytes.Buffer{}, nil, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return bytes.Buffer{}, nil, err
	}
	w.Close()
	return b, w, nil
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println("PWD: ", pwd)
		panic(err)
	}
	return r
}
