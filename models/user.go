package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"regexp"
	"time"
)

type User struct {
	gorm.Model
	Nama        string    `gorm:"column:nama;type:varchar(50)"`
	TempatLahir string    `gorm:"column:tempatLahir;type:varchar(50)"`
	TglLahir    time.Time `gorm:"column:tgl_lahir"`
	Alamat      string    `gorm:"column:alamat;type:varchar(255)"`
	Username    string    `gorm:"column:username;type:varchar(50);unique"`
	Password    []byte    `gorm:"column:password"`
	Level       uint8     `gorm:"column:level;type:int(1)"`
}

var userRegex = regexp.MustCompile("^\\w*$")

func (u User) Validate(rv *revel.Validation) {
	rv.Check(
		u.Nama,
		revel.Required{},
	)

	rv.Check(
		u.Username,
		revel.Required{},
		revel.MinSize{6},
		revel.Match{userRegex},
	)

	rv.Check(
		u.TglLahir,
		revel.Required{},
	)

	rv.Check(
		u.Alamat,
		revel.Required{},
	)
}