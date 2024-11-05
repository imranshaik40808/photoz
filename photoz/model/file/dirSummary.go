package model

type DirSummary struct {
	DeniedFileAndFolderCount int64      `json:"denied_file_and_folder_count,omitempty"`
	DeniedFilesAndFolders    []string   `json:"denied_file_and_folder,omitempty"`
	DirCount                 int64      `json:"dir_count,omitempty"`
	FileCount                int64      `json:"file_count,omitempty"`
	Files                    []FileInfo `json:"files,omitempty"`
	TotalSize                int64      `json:"total_size,omitempty"`
	ExcludedFileCount        int64      `json:"excluded_file_count,omitempty"`
	ExcludedFiles            []string   `json:"excluded_files,omitempty"`
	ExcludedDirectoryCount   int64      `json:"excluded_directory_count,omitempty"`
	ExcludedDirectories      []string   `json:"excluded_directories,omitempty"`
}
