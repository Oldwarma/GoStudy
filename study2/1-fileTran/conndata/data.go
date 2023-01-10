package conndata

type Msg struct {
	Head Head
	Data Data
}

type Head struct {
	FunName   string
	Cot       string
	Mid       string
	Timestamp string
}
type Data struct {
	Info    Info
	Content []byte
}
type Info struct {
	File     string `json:"file,omitempty"`
	FileName string `json:"fileName,omitempty"`
	Name     string `json:"name,omitempty"`
	Remark   string `json:"remark,omitempty"`
	Size     string `json:"size,omitempty"`
	Version  string `json:"version,omitempty"`
}
