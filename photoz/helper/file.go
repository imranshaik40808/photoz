package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	model "photoz/model/file"
)

func HashFile(filename string) (string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a new SHA256 hash
	hash := sha256.New()

	// Copy the file's content into the hash
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Get the final hash result as a byte slice
	hashInBytes := hash.Sum(nil)

	// Convert bytes to a hexadecimal string
	return hex.EncodeToString(hashInBytes), nil
}

func GetFileInfo(fileName string) (*model.FileInfo, error) {
	exists := FileExists(fileName)
	if exists {
		info, _ := os.Stat(fileName)
		fileInfo := &model.FileInfo{
			Name:        info.Name(),
			Path:        fileName,
			Extension:   filepath.Ext(fileName),
			Size:        info.Size(),
			CreatedDate: info.ModTime(),
		}
		return fileInfo, nil
	}
	return nil, errors.New("empty name")
}

func GetSummary(dir string, excludePath []string, excludeExtension []string) *model.DirSummary {

	excludePathMap := getExcludePathMap(excludePath)
	excludeExtensionMap := getPathExtension(excludeExtension)
	var files []model.FileInfo
	var deniedFiles []string
	var summary model.DirSummary
	fileCount := 0
	dirCount := 0
	accessDeniedFileAndFolder := 0
	totalSize := 0
	excludedFileCount := 0
	excludedFiles := make([]string, 0)
	excludedDirectoryCount := 0
	excludedDirectories := make([]string, 0)
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				deniedFiles = append(deniedFiles, path)
				accessDeniedFileAndFolder += 1
				return nil
			}
			if !info.IsDir() {
				ext := filepath.Ext(path)
				if _, ok := excludeExtensionMap[ext]; ok {
					excludedFiles = append(excludedFiles, path)
					excludedFileCount += 1
					return nil
				}
				fileCount += 1
				hash, err := HashFile(path)
				if err != nil {
					println("Failed to read file %s", path)
				}
				fileInfo := model.FileInfo{
					Name:        info.Name(),
					Path:        path,
					Extension:   ext,
					Size:        info.Size(),
					CreatedDate: info.ModTime(),
					Hash:        hash,
				}
				files = append(files, fileInfo)
				totalSize += int(info.Size())
			} else {
				if _, ok := excludePathMap[path]; ok {
					excludedDirectories = append(excludedDirectories, path)
					excludedDirectoryCount += 1
					return filepath.SkipDir
				} else {
					dirCount += 1
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	summary = model.DirSummary{
		DeniedFileAndFolderCount: int64(accessDeniedFileAndFolder),
		DeniedFilesAndFolders:    deniedFiles,
		TotalSize:                int64(totalSize),
		FileCount:                int64(fileCount),
		DirCount:                 int64(dirCount),
		Files:                    files,
		ExcludedFileCount:        int64(excludedFileCount),
		ExcludedFiles:            excludedFiles,
		ExcludedDirectoryCount:   int64(excludedDirectoryCount),
		ExcludedDirectories:      excludedDirectories,
	}
	return &summary
}

func getPathExtension(excludeExtension []string) map[string]string {
	var excludeExtensionMap map[string]string = make(map[string]string)
	if len(excludeExtension) > 0 {
		for _, v := range excludeExtension {
			excludeExtensionMap[v] = ""
		}
	}
	return excludeExtensionMap
}

func getExcludePathMap(excludePath []string) map[string]string {
	var excludePathMap map[string]string = make(map[string]string)
	if len(excludePath) > 0 {
		for _, v := range excludePath {
			excludePathMap[v] = ""
		}
	}
	return excludePathMap
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func MoveFile(src, dest string) error {
	// Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	// Create the destination file
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	// Copy the content from source to destination
	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	// Remove the original source file
	if err := os.Remove(src); err != nil {
		return fmt.Errorf("failed to remove source file: %w", err)
	}

	return nil
}
