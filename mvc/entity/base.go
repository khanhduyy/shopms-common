package entity

type Base struct {
	Id uint `gorm:"column:Id;primaryKey;autoIncrement:true"`
}

func NewBase(id uint) *Base {
	return &Base{Id: id}
}

func NewBaseInt64(id int64) *Base {
	if id > 0 {
		return &Base{Id: uint(id)}
	}
	return nil
}
