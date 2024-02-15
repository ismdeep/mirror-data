package mirror

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// parse mirror text file data
// @parameter s string content text
// @return bucketName string bucket name
// @return file       File   file data
func parseText(s string) (string, File, error) {
	items := strings.Split(s, "|")
	if len(items) != 6 {
		return "", File{}, errors.New("invalid text")
	}

	contentLength, err := strconv.ParseInt(items[3], 10, 64)
	if err != nil {
		return "", File{}, errors.New("invalid content-length")
	}

	lastModifiedUnix, err := strconv.ParseInt(items[5], 10, 64)
	if err != nil {
		return "", File{}, errors.New("invalid last-modified")
	}

	file := File{
		Path:          items[1],
		OriginURL:     items[2],
		ContentLength: contentLength,
		ContentType:   items[4],
		LastModified:  time.Unix(lastModifiedUnix, 0),
	}

	return items[0], file, nil
}
