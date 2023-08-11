package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type LinkPair struct {
	Link       string
	OriginLink string
}

type Storage struct {
	BucketName    string
	ExistsMap     map[string]bool
	Fp            *os.File
	FpMutex       sync.Mutex
	C             chan LinkPair
	WG            sync.WaitGroup
	CoroutineSize int
}

func NewStorage(bucketName string, coroutineSize int) (*Storage, error) {

	var storage Storage
	storage.BucketName = bucketName
	storage.CoroutineSize = coroutineSize
	storage.ExistsMap = make(map[string]bool)
	storage.C = make(chan LinkPair, 1024)

	fp, err := os.OpenFile(fmt.Sprintf("./data/%v.txt", bucketName), os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fp.Close()
	}()

	r := bufio.NewReaderSize(fp, 65535)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		s = strings.TrimSpace(s)
		items := strings.Split(s, "|")
		storage.ExistsMap[items[1]] = true
	}

	fpWrite, err := os.OpenFile(fmt.Sprintf("./data/%v.txt", bucketName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	storage.Fp = fpWrite

	return &storage, nil
}

func (receiver *Storage) Add(link string, originLink string) {
	receiver.C <- LinkPair{
		Link:       link,
		OriginLink: originLink,
	}
}

func (receiver *Storage) write(content string) error {
	receiver.FpMutex.Lock()
	defer func() {
		receiver.FpMutex.Unlock()
	}()

	if _, err := receiver.Fp.WriteString(content); err != nil {
		return err
	}

	return nil
}

func (receiver *Storage) StartConsumer() {
	for i := 0; i < receiver.CoroutineSize; i++ {
		receiver.WG.Add(1)
		go func() {
			for item := range receiver.C {
				if _, ok := receiver.ExistsMap[item.Link]; ok {
					fmt.Println(item.Link, "already exists")
					continue
				}

				req, err := http.NewRequest(http.MethodHead, item.OriginLink, nil)
				if err != nil {
					continue
				}
				resp, err := (&http.Client{}).Do(req)
				if err != nil {
					continue
				}

				lastModified, err := time.Parse(time.RFC1123, resp.Header.Get("Last-Modified"))
				if err != nil {
					continue
				}
				contentLength := resp.Header.Get("Content-Length")
				contentType := resp.Header.Get("Content-Type")

				_ = receiver.write(fmt.Sprintf("%v|%v|%v|%v|%v|%v\n",
					receiver.BucketName,
					item.Link,
					item.OriginLink,
					contentLength,
					contentType,
					lastModified.Unix()))
				fmt.Println(item.Link, item.OriginLink)
			}
			receiver.WG.Done()
		}()
	}
}
