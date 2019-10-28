package handlers

import (
	"coco"
	"encoding/json"
	. "polo/common"
	"polo/models"
	"polo/utils"
)

type UploadContentParam struct {
	Content string
}

// UploadTextContentHandler 上传文本内容，你应该使用 json 格式传递参数
func UploadTextContentHandler(c *coco.Coco) coco.Result {
	decoder := json.NewDecoder(c.Request.Body)
	var param UploadContentParam
	err := decoder.Decode(&param)
	if err != nil {
		panic(err)
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
	meta, err := models.AddFileMeta(fh.Filename, fileHash, block.ID, block.Size-int64(length), fh.Size)
	if err != nil {
		Logger.Error(err)
		return coco.ErrorResponse(100)
	}
}
