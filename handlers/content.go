package handlers

import (
	"encoding/json"
	"github.com/olivetree123/coco"
	. "polo/common"
	"polo/models"
	"polo/utils"
)

type UploadContentParam struct {
	Content string
}

// UploadContentHandler 上传文本内容，你应该使用 json 格式传递参数
func UploadContentHandler(c *coco.Coco) coco.Result {
	decoder := json.NewDecoder(c.Request.Body)
	var param UploadContentParam
	err := decoder.Decode(&param)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	content := []byte(param.Content)
	block, err := models.GetAvailableBlock(int64(len(content)))
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
	meta, err := models.AddContentMeta(fileHash, block.ID, block.Size-int64(len(content)), int64(len(content)))
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
	return coco.APIResponse(meta)
}

// GetContentHandler 获取内容
func GetContentHandler(c *coco.Coco) coco.Result {
	hash := c.Params.ByName("hash")
	meta, err := models.GetContentMeta(hash)
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
	return coco.APIResponse(string(content))
}
