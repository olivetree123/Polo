package handlers

import (
	"github.com/olivetree123/coco"
	. "polo/common"
	"polo/models"
	"polo/utils"
)

// UploadFileHandler 上传文件
func UploadFileHandler(c *coco.Coco) coco.Result {
	f, fh, err := c.Request.FormFile("file")
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	content := make([]byte, fh.Size)
	length, err := f.Read(content)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	block, err := models.GetAvailableBlock(int64(length))
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	err = block.Write(content)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	fileHash := utils.ContentMD5(content)
	meta, err := models.AddFileMeta(fh.Filename, fileHash, block.ID, block.Size-int64(length), fh.Size)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	return coco.APIResponse(meta)
}

// FileInfoHandler 获取文件信息，文件大小、下载次数
func FileInfoHandler(c *coco.Coco) coco.Result {
	hash := c.Params.ByName("hash")
	meta, err := models.GetFileMeta(hash)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	return coco.APIResponse(meta)
}
