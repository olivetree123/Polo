package handlers

import (
	"github.com/olivetree123/coco"
	. "polo/common"
	"polo/models"
)

func DownloadHandler(c *coco.Coco) coco.Result {
	hash := c.Params.ByName("hash")
	meta, err := models.GetFileMeta(hash)
	if err != nil {
		Logger.Error(err)
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
	return coco.FileResponse(meta.FileName, content)
}
