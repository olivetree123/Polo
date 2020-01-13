package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mitchellh/mapstructure"
	. "polo/common"
	"polo/config"
	"time"
)

// DB 数据库连接对象
var DB *gorm.DB

// BaseModel 基础模型
type BaseModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Status    bool      `gorm:"default:true" json:"status"` // 0 删除，1 正常
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (model *BaseModel) Remove() bool {
	model.Status = false
	//model.UpdatedAt = time.Now()
	return true
}

// InitDB 数据库初始化
func InitDB() {
	// driver := revel.Config.StringDefault("db.driver", "")
	// connect_string := revel.Config.StringDefault("db.connect", "")
	// DB, err = sql.Open(driver, connect_string)
	var err error
	mysqlConf := config.Config.GetStringMap("mysql")
	mysqlURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConf["user"],
		mysqlConf["password"],
		mysqlConf["host"],
		mysqlConf["port"],
		mysqlConf["db"],
	)
	Logger.Info("mysqlURL = ", mysqlURL)
	DB, err = gorm.Open("mysql", mysqlURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB.SingularTable(true)
	if !DB.HasTable(Block{}) {
		DB.CreateTable(Block{})
	}
	if !DB.HasTable(FileMeta{}) {
		DB.CreateTable(FileMeta{})
	}
	if !DB.HasTable(ContentMeta{}) {
		DB.CreateTable(ContentMeta{})
	}
}

type BlockParam struct {
	Path string
	Size int
}

func InitBlock() {
	blocks := config.Config.GetStringMap("block")
	for _, val := range blocks {
		var param BlockParam
		err := mapstructure.Decode(val, &param)
		if err != nil {
			Logger.Error(err)
			return
		}
		_, err = GetBlockByPath(param.Path)
		if err == nil {
			Logger.Info("block found, continue, path = ", param.Path)
			continue
		}
		if err.Error() == "record not found" {
			_, err = NewBlock(param.Path, int64(param.Size)*1024*1024*1024)
			if err != nil {
				Logger.Error(err)
				return
			}
		}
	}
}

func init() {
	InitDB()
	InitBlock()
}
