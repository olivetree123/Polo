package models

import (
	"os"
	"path/filepath"
	. "polo/common"
	"strconv"
)

const BlockMaxSize int64 = 1024 * 1024 * 1024 * 5 // 5GB

// Block 存储块
type Block struct {
	BaseModel
	Path    string `gorm:"unique"`
	MaxSize int64  // 最大 size
	Size    int64  // 当前 size
}

// NewBlock 新建 Block
func NewBlock(path string, maxSize int64) (*Block, error) {
	block := Block{
		Path:    path,
		MaxSize: maxSize,
		Size:    0,
	}
	err := DB.Create(&block).Error
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	Logger.Info("block = ", block)
	return &block, nil
}

func GetBlock(blockID uint) (*Block, error) {
	var block Block
	err := DB.First(&block, "id = ?", blockID).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// ListBlock 列出所有 block
func ListBlock() ([]Block, error) {
	var blocks []Block
	err := DB.Find(&blocks).Error
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	return blocks, nil
}

func GetBlockByPath(path string) (*Block, error) {
	var block Block
	err := DB.First(&block, "path = ?", path).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// FreeSize 计算 block 剩余空间大小
func (block *Block) FreeSize() int64 {
	return block.MaxSize - block.Size
}

// GetAvailableBlock 获取足够大小的 block
func GetAvailableBlock(size int64) (*Block, error) {
	blocks, err := ListBlock()
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	for _, block := range blocks {
		if block.FreeSize() > size {
			return &block, nil
		}
	}
	return nil, nil
}

func (block *Block) Write(content []byte) error {
	f, err := os.OpenFile(block.GetPath(), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	_, err = f.WriteAt(content, block.Size)
	if err != nil {
		return err
	}
	block.Size += int64(len(content))
	err = DB.Save(block).Error
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func (block *Block) Read(offset int64, length int64) ([]byte, error) {
	f, err := os.OpenFile(block.GetPath(), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	content := make([]byte, length)
	_, err = f.ReadAt(content, offset)
	if err != nil {
		return nil, err
	}
	f.Close()
	return content, nil
}

func (block *Block) GetPath() string {
	return filepath.Join(block.Path, strconv.Itoa(int(block.ID))+".db")
}
