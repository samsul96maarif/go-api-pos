package request

import "mime/multipart"

type Im struct {
	File       multipart.File        `json:"file"`
	Compare    multipart.File        `json:"compare"`
	FileHeader *multipart.FileHeader `json:"file_header"`
}

type D struct {
	Nib  uint
	Desc string
}
