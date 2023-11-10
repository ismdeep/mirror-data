package store

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

	"github.com/ismdeep/mirror-data/pkg/log"
	"go.uber.org/zap"
)

// LinkPair link pair
type LinkPair struct {
	Link       string
	OriginLink string
}

// Storage model
type Storage struct {
	BucketName    string
	ExistsMap     map[string]bool
	Fp            *os.File
	FpMutex       sync.Mutex
	C             chan LinkPair
	WG            sync.WaitGroup
	CoroutineSize int
}

// New create a storage
func New(bucketName string, coroutineSize int) *Storage {
	var storage Storage
	storage.BucketName = bucketName
	storage.CoroutineSize = coroutineSize
	storage.ExistsMap = make(map[string]bool)
	storage.C = make(chan LinkPair, 1024)

	fp, err := os.OpenFile(fmt.Sprintf("./data/%v.txt", bucketName), os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	storage.Fp = fpWrite

	storage.startConsumer()

	return &storage
}

// Add append link pair
func (receiver *Storage) Add(link string, originLink string) {
	receiver.C <- LinkPair{
		Link:       link,
		OriginLink: originLink,
	}
}

func (receiver *Storage) write(link string, content string) error {
	receiver.FpMutex.Lock()
	defer func() {
		receiver.FpMutex.Unlock()
	}()

	if _, ok := receiver.ExistsMap[link]; ok {
		return nil
	}

	if _, err := receiver.Fp.WriteString(content); err != nil {
		return err
	}

	receiver.ExistsMap[link] = true

	return nil
}

func (receiver *Storage) startConsumer() {
	for i := 0; i < receiver.CoroutineSize; i++ {
		receiver.WG.Add(1)
		go func() {
			for item := range receiver.C {
				if _, ok := receiver.ExistsMap[item.Link]; ok {
					log.WithName(receiver.BucketName).Debug("already exists", zap.String("link", item.Link))
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

				_ = receiver.write(item.Link, fmt.Sprintf("%v|%v|%v|%v|%v|%v\n",
					receiver.BucketName,
					item.Link,
					item.OriginLink,
					contentLength,
					contentType,
					lastModified.Unix()))
				log.WithName(receiver.BucketName).Info("saved", zap.String("link", item.Link))
			}
			receiver.WG.Done()
		}()
	}
}

// CloseAndWait close C and wait
func (receiver *Storage) CloseAndWait() {
	close(receiver.C)
	receiver.WG.Wait()
}
