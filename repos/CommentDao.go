package repos

import (
	"../conf"
	"../entities"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// データベースにコメントをInsertする。
// エラーが発生した場合は戻り値としてエラーを返す
func Create(comment entities.Comment) error {
	db, err := gorm.Open("mysql", conf.DbUri)
	if err != nil {
		return err
	}
	db.DB()
	db.Create(comment)
	return nil
}
