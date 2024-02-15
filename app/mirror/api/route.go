package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ismdeep/mirror-data/app/mirror/conf"
	"github.com/ismdeep/mirror-data/app/mirror/internal/mirror"
)

var eng *gin.Engine

func init() {
	gin.SetMode(conf.ROOT.Server.Mode)
	eng = gin.New()

	eng.GET("/*path", func(c *gin.Context) {
		path := c.Param("path")

		// if is a file
		if fileInfo := mirror.M.GetFileInfo(path); fileInfo != nil {
			ProxyFile(c, fileInfo)
			return
		}

		// if is folder
		ShowFolder(c, path)
	})

	eng.HEAD("/*path", func(c *gin.Context) {
		path := c.Param("path")
		fileInfo := mirror.M.GetFileInfo(path)
		if fileInfo == nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.Header("Content-Length", fmt.Sprintf("%v", fileInfo.ContentLength))
		c.Header("Content-Type", fileInfo.ContentType)
		c.Header("Accept-Ranges", "bytes")
		c.Header("Last-Modified", fileInfo.LastModified.Format(time.RFC1123))
	})
}

// Run http server
func Run() error {
	fmt.Printf(`

  #     #   ###   ######    ######    #######   ######  
  ##   ##    #    #     #   #     #   #     #   #     # 
  # # # #    #    #     #   #     #   #     #   #     # 
  #  #  #    #    ######    ######    #     #   ######  
  #     #    #    #   #     #   #     #     #   #   #   
  #     #    #    #    #    #    #    #     #   #    #  
  #     #   ###   #     #   #     #   #######   #     #


`)
	fmt.Println("API SERVER STARTED:", conf.ROOT.Server.Bind)
	return eng.Run(conf.ROOT.Server.Bind)
}
