package uploader

type Respond struct {
	Status bool         `json:"status"`
	Error  ErrorRespond `json:"error,omitempty"`
	Data   Data         `json:"data,omitempty"`
}

type ErrorRespond struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Data struct {
	File File   `json:"file"`
	Type string `json:"type"`
}

type File struct {
	Url      url      `json:"url"`
	MetaData MetaData `json:"metadata"`
}

type MetaData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type url struct {
	Full  string `json:"full"`
	Short string `json:"short"`
}

const (
	anonfiles = "https://api.anonfiles.com/upload"
	bayfiles  = "https://api.bayfiles.com/upload"
)
