package sessionContent

type SESSION int

// SESSION
const (
	Alive = iota
	UserID
)

func (s SESSION) String() string {
	switch s {
	case Alive:
		return "Alive"
	case UserID:
		return "UserID"
	default:
		return "Unknown"
	}
}
