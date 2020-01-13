package main

import (
	"github.com/olivetree123/coco"
	"github.com/sirupsen/logrus"
	"polo/handlers"
)

func main() {
	c := coco.NewCoco()
	c.Header["Access-Control-Allow-Origin"] = "*"
	c.Header["Access-Control-Allow-Methods"] = "POST, GET, OPTIONS, PUT, DELETE"
	c.Header["Access-Control-Allow-Headers"] = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	c.AddRouter("POST", "/polo/upload/content", handlers.UploadContentHandler)
	c.AddRouter("POST", "/polo/upload/file", handlers.UploadFileHandler)
	c.AddRouter("GET", "/polo/object/:hash", handlers.DownloadHandler)
	c.AddRouter("GET", "/polo/content/:hash", handlers.GetContentHandler)
	c.AddRouter("OPTIONS", "/polo/object/:hash", handlers.OptionsHandler)
	c.AddRouter("GET", "/polo/info/:hash", handlers.FileInfoHandler)
	err := c.Run("0.0.0.0", 8300)
	if err != nil {
		logrus.Error(err)
	}
}
