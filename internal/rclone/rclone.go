package rclone

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
)

// JSONObj json object
type JSONObj struct {
	Path     string `json:"Path"`
	Name     string `json:"Name"`
	Size     int64  `json:"Size"`
	MimeType string `json:"MimeType"`
	ModTime  string `json:"ModTime"`
	IsDir    bool   `json:"IsDir"`
}

// JSON get json
func JSON(args ...string) ([]JSONObj, error) {
	var buf bytes.Buffer
	cmd := exec.Command("rclone", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var lst []JSONObj
	if err := json.Unmarshal(buf.Bytes(), &lst); err != nil {
		return nil, err
	}

	return lst, nil
}
