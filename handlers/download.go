package handlers

import (
	"github.com/olivetree123/coco"
	. "polo/common"
	"polo/models"
)

// DownloadHandler 下载
func DownloadHandler(c *coco.Coco) coco.Result {
	hash := c.Params.ByName("hash")
	meta, err := models.GetFileMeta(hash[:32])
	if err != nil {
		Logger.Error(err, ", fileHash = ", hash[:32])
		return coco.ErrorResponse(100)
	}
	block, err := models.GetBlock(meta.BlockID)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	content, err := block.Read(meta.Offset, meta.Length)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	err = meta.AddDownloadCount()
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	return coco.FileResponse(meta.FileName, content)
}
