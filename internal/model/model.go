package model

type User struct {
	IdUser   int    `gorm:"primaryKey;column:id_user;autoIncrement" json:"-"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Npm      string `gorm:"column:npm" json:"npm"`
	Kelas    string `gorm:"column:kelas" json:"kelas"`
	AsalKota string `gorm:"column:asal_kota" json:"asal_kota"`
}
