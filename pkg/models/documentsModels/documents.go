package documentsModels

type Document struct {
	Id        int64    `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	File      bool     `json:"file,omitempty"`
	Public    bool     `json:"public,omitempty"`
	Token     string   `json:"token,omitempty"`
	Mime      string   `json:"mime,omitempty"`
	Grant     []string `json:"grant,omitempty"`
	Json      string   `json:"json,omitempty"`
	Directory string   `json:"directory,omitempty"`
	Created   string   `json:"created,omitempty"`
}

type Data struct {
	Json string `json:"json,omitempty"`
	File string `json:"file,omitempty"`
}

type Pass struct {
	Limit int    `json:"limit,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
