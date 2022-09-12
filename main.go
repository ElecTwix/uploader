package uploader

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Site string

const (
	AnonFiles Site = "https://api.anonfiles.com/upload"
	BayFiles  Site = "https://api.bayfiles.com/upload"
)

func Upload(url Site, file *os.File) (*Respond, error) {

	var resjson Respond
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	err = writer.Close()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, string(url), body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &resjson)
	if err != nil {
		return nil, err
	}
	return &resjson, err
}
