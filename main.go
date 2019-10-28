package main

import (
	"github.com/olivetree123/coco"
	"github.com/sirupsen/logrus"
	"polo/handlers"
)

func main() {
	c := coco.NewCoco()
	c.AddRouter("POST", "/polo/upload", handlers.UploadHandler)
	c.AddRouter("GET", "/polo/object/:hash", handlers.DownloadHandler)
	err := c.Run("0.0.0.0", 8300)
	if err != nil {
		logrus.Error(err)
	}
}