package mirror

import (
	"fmt"
	"strings"
	"time"
)

// WebFile struct
type WebFile struct {
	LastModified time.Time
	F            *File
}

// Mirror struct
type Mirror struct {
	data           map[string][]*File
	filePtr        map[string]*File
	bucketNameList []string

	files   map[string]*File
	folders map[string]map[string]WebFile
}

// NewMirror new a mirror
func NewMirror() *Mirror {
	var m Mirror
	m.data = make(map[string][]*File)
	m.filePtr = make(map[string]*File)
	m.files = make(map[string]*File)
	m.folders = make(map[string]map[string]WebFile)
	return &m
}

// AddFile add file to bucket
// @return ok bool. false: nothing changed, true: file data have been written
func (receiver *Mirror) AddFile(bucketName string, file File) bool {
	localFile, ok := receiver.filePtr[fmt.Sprintf("%v|%v", bucketName, file.Path)]

	// check if exists and already exists the same value
	if ok &&
		localFile.OriginURL == file.OriginURL &&
		localFile.ContentLength == file.ContentLength &&
		localFile.ContentType == file.ContentType &&
		localFile.LastModified.Unix() == file.LastModified.Unix() {
		return false
	}

	receiver.data[bucketName] = append(receiver.data[bucketName], &file)
	receiver.filePtr[fmt.Sprintf("%v|%v", bucketName, file.Path)] = &file

	path := fmt.Sprintf("/%v/%v", bucketName, file.Path)
	receiver.files[path] = &file

	folderPath := path[:strings.LastIndex(path, "/")+1]
	fileName := path[strings.LastIndex(path, "/")+1:]

	// check nil
	if _, ok := receiver.folders[folderPath]; !ok {
		receiver.folders[folderPath] = make(map[string]WebFile)
	}

	// add folder
	receiver.AddFolder(folderPath, file.LastModified)

	// add file
	receiver.folders[folderPath][fileName] = WebFile{
		F: &file,
	}

	return true
}

// AddFolder add a folder
func (receiver *Mirror) AddFolder(path string, lastModified time.Time) {
	if path == "/" {
		return
	}

	path = path[:len(path)-1]

	parentPath := path[:strings.LastIndex(path, "/")+1]
	name := path[strings.LastIndex(path, "/")+1:]

	// check nil
	if _, ok := receiver.folders[parentPath]; !ok {
		receiver.folders[parentPath] = make(map[string]WebFile)
	}

	// add folder
	if t, ok := receiver.folders[parentPath][name]; ok {
		lm := receiver.folders[parentPath][name].LastModified
		if lm.UnixNano() < lastModified.UnixNano() {
			t.LastModified = lastModified
			receiver.folders[parentPath][name] = t
		}
	} else {
		receiver.folders[parentPath][name] = WebFile{
			LastModified: lastModified,
			F:            nil,
		}
	}

	receiver.AddFolder(parentPath, lastModified)
}

// FileExists check file is exists
func (receiver *Mirror) FileExists(bucketName string, path string) bool {
	_, ok := receiver.filePtr[fmt.Sprintf("%v|%v", bucketName, path)]
	return ok
}

// ListBuckets get bucket name list
func (receiver *Mirror) ListBuckets() []string {
	return receiver.bucketNameList
}

// ListFiles get file name list of bucket
func (receiver *Mirror) ListFiles(bucketName string) []*File {
	return receiver.data[bucketName]
}

// GetFolderInfo get folder info
func (receiver *Mirror) GetFolderInfo(parentPath string) map[string]WebFile {
	d, ok := receiver.folders[parentPath]
	if !ok {
		return nil
	}

	return d
}

// GetFileInfo get file info
func (receiver *Mirror) GetFileInfo(path string) *File {
	f, ok := receiver.files[path]
	if !ok {
		return nil
	}

	return f
}
