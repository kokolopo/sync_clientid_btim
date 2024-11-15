package utils

type UserBlink struct {
	KycID      int    `gorm:"column:KycID"`
	ClientID   string `gorm:"column:ClientID"`
	Email      string `gorm:"column:Email"`
	IdentityNo string `gorm:"column:IdentityNo"`
}

type ClientIDBTIM struct {
	ClientID int `gorm:"column:ClientID"`
}
