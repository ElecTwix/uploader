package uploader

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	anonfiles string = "https://api.anonfiles.com/upload"
	bayfiles  string = "https://api.bayfiles.com/upload"
)

func Upload(url string, file *os.File) (Respond, error) {
	var resjson Respond
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("file", filepath.Base(file.Name()))
		if err != nil {
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	req, err := http.NewRequest(http.MethodPost, url, r)
	req.Header.Add("Content-Type", m.FormDataContentType())
	if err != nil {
		return resjson, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return resjson, err
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resjson, err
	}

	err = json.Unmarshal(content, &resjson)
	if err != nil {
		return resjson, err
	}
	return resjson, err
}
