package content

// MODE production or development
const MODE = "development"

// const MODE = "production"

// SessionMode true: セッションを使用する  false: セッションにmockを適用する
const SessionMode = false

const (
	// SessionAlive ログインしている状態か否かを判定
	SessionAlive = iota
	// SessionUserID ログインしているユーザのIDを取得
	SessionUserID
)
