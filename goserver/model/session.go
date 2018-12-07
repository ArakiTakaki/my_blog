package model

// SessionInfo セッション
type SessionInfo struct {
	UserID         interface{} //ログインしているユーザのID
	IsSessionAlive bool        //セッションが生きているかどうか
}
