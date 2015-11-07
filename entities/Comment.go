package entities

import (
	"time"
)

// コメントの内容を保持する
type Comment struct {
	Id        int       // コメントID
	Body      string    // コメント本文
	UserName  string    // コメントをしたユーザー名
	FilePath  string    // コメントがついたファイル名。issueの場合は空
	CreatedAt time.Time // コメントが最初に投稿された時間
	UpdatedAt time.Time // コメントが最後に更新された時間
}
