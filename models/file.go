package models

// FileMeta 文件元数据
type FileMeta struct {
	BaseModel
	FileName string `json:"fileName"`
	FileHash string `json:"fileHash"`
	BlockID  uint   `json:"blockID"`
	Offset   int64  `json:"offset"`
	Length   int64  `json:"length"`
	Count1   int    `json:"count1"` // 下载次数
}

// AddFileMeta 添加元数据
func AddFileMeta(fileName string, fileHash string, blockID uint, offset int64, length int64) (*FileMeta, error) {
	meta := FileMeta{
		FileName: fileName,
		FileHash: fileHash,
		BlockID:  blockID,
		Offset:   offset,
		Length:   length,
	}
	if err := DB.Create(&meta).Error; err != nil {
		return nil, err
	}
	return &meta, nil
}

// GetFileMeta 根据哈希值获取元数据
func GetFileMeta(fileHash string) (*FileMeta, error) {
	var meta FileMeta
	if err := DB.First(&meta, "file_hash = ?", fileHash).Error; err != nil {
		return nil, err
	}
	return &meta, nil
}

func (meta *FileMeta) AddDownloadCount() error {
	meta.Count1 += 1
	err := DB.Save(&meta).Error
	return err
}
