package solution

type Solution struct {
	ID     uint64 `gorm:"primaryKey"`
	UserID uint64
	TaskID uint64
	Code   string `json:"description"`
	IsDone bool
}
