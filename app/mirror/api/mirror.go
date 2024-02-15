package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ismdeep/mirror-data/app/mirror/conf"
	"github.com/ismdeep/mirror-data/app/mirror/internal/mirror"
	"github.com/ismdeep/mirror-data/app/mirror/internal/util"
)

// ProxyFile proxy file
func ProxyFile(c *gin.Context, fileInfo *mirror.File) {
	// https://dl.google.com/go/go1.20.5.darwin-amd64.tar.gz
	originURL, err := util.RealOriginURL(fileInfo.OriginURL)
	if err != nil {
		c.Status(http.StatusBadGateway)
		_, _ = c.Writer.WriteString("ERROR")
		return
	}

	// redirect 307
	if c.GetHeader("X-DOWNLOAD-TOKEN") != conf.ROOT.System.DownloadToken {
		c.Redirect(http.StatusTemporaryRedirect, originURL)
		return
	}

	site, reqPath, query := util.DivideURL(originURL)
	remote, err := url.Parse(site)
	if err != nil {
		c.Status(http.StatusBadGateway)
		_, _ = c.Writer.WriteString("ERROR")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = reqPath
		req.URL.RawQuery = query
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

// ShowFolder show folder page
func ShowFolder(c *gin.Context, path string) {
	webFiles := mirror.M.GetFolderInfo(path)
	type link struct {
		Name            string
		F               *mirror.File
		IsDir           bool
		DirLastModified time.Time
	}

	var links []link

	for name, file := range webFiles {
		links = append(links, link{
			Name:            name,
			F:               file.F,
			IsDir:           file.F == nil,
			DirLastModified: file.LastModified,
		})
	}

	sort.SliceStable(links, func(i, j int) bool {
		return strings.Compare(links[i].Name, links[j].Name) <= 0
	})

	var s strings.Builder

	if path != "/" {
		s.WriteString(fmt.Sprintf(`<a href="../">../</a>%v`, "\n"))
	}

	for _, l := range links {
		if l.IsDir {
			s.WriteString(fmt.Sprintf(`<a href="%v/">%v/</a>    %v%v     -%v`, l.Name, l.Name, util.Space(49-len(l.Name)), l.DirLastModified.Format("2006-01-02 15:04"), "\n"))
		} else {
			s.WriteString(fmt.Sprintf(`<a href="%v">%v</a>    %v%v    %v%v`, l.Name, l.Name, util.Space(50-len(l.Name)), l.F.LastModified.Format("2006-01-02 15:04"), l.F.ContentLength, "\n"))
		}
	}

	c.Header("Content-Type", "text/html")
	_, _ = c.Writer.WriteString(fmt.Sprintf(`<html><head><title>Index of %v</title></head><body><h1>Index of %v</h1><hr><pre>%v</pre><hr></body></html>`, path, path, s.String()))
}
