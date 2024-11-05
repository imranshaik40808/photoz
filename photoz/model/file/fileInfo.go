package model

import "time"

type FileInfo struct {
	Name        string    `json:"name,omitempty"`
	Path        string    `json:"path,omitempty"`
	Extension   string    `json:"extension,omitempty"`
	Size        int64     `json:"size,omitempty"`
	Attribute   []byte    `json:"attribute,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	Hash        string    `json:"-"`
}
