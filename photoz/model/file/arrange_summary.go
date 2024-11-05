package model

type ArrangeSummary struct {
	FailedFiles         []string `json:"failed_files,omitempty"`
	FailedFileSize      int      `json:"failed_file_size,omitempty"`
	FailedFileCount     int      `json:"failed_file_count,omitempty"`
	SuccessfulFiles     []string `json:"successful_files,omitempty"`
	SuccessfulFileSize  int      `json:"successful_file_size,omitempty"`
	SuccessfulFileCount int      `json:"successful_file_count,omitempty"`
	DuplicateFiles      []string `json:"duplicate_files,omitempty"`
	DuplicateFileSize   int      `json:"duplicate_file_size,omitempty"`
	DuplicateFileCount  int      `json:"duplicate_file_count,omitempty"`
}
