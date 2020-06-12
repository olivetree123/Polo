package models

// ContentMeta 内容元数据
type ContentMeta struct {
	BaseModel
	HashValue string `json:"hashValue"`
	BlockID   uint   `json:"blockID"`
	Offset    int64  `json:"offset"`
	Length    int64  `json:"length"`
}

// AddContentMeta 增加元数据
func AddContentMeta(hashValue string, blockID uint, offset int64, length int64) (*ContentMeta, error) {
	meta := ContentMeta{
		HashValue: hashValue,
		BlockID:   blockID,
		Offset:    offset,
		Length:    length,
	}
	if err := DB.Create(&meta).Error; err != nil {
		return nil, err
	}
	return &meta, nil
}

// GetContentMeta 获取内容元数据
func GetContentMeta(hashValue string) (*ContentMeta, error) {
	var meta ContentMeta
	if err := DB.First(&meta, "hash_value = ?", hashValue).Error; err != nil {
		return nil, err
	}
	return &meta, nil
}
