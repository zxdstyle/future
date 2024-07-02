package model

type FileData struct {
	Ext     string
	Name    string
	Data    []byte
	Headers map[string]string
}
