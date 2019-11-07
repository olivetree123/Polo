package models

import (
	. "polo/common"
)

// 内容元数据
type ContentMeta struct {
	BaseModel
	HashValue string `json:"hashValue"`
	BlockID   uint   `json:"blockID"`
	Offset    int64  `json:"offset"`
	Length    int64  `json:"length"`
}

// 增加元数据
func AddContentMeta(hashValue string, blockID uint, offset int64, length int64) (*ContentMeta, error) {
	meta := ContentMeta{
		HashValue: hashValue,
		BlockID:   blockID,
		Offset:    offset,
		Length:    length,
	}
	err := DB.Create(&meta).Error
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	return &meta, nil
}
