package controllers

import (
	r "github.com/revel/revel"
	"time"
)

const (
	SQL_DATE  = "2006-01-02T15:04:05Z07:00"
	SHOW_DATE = "02-01-2006"
)

func init() {

	r.TemplateFuncs["indonesiandate"] = func(value string) string {
		t, _ := time.Parse(SQL_DATE, value)
		return t.Format(SHOW_DATE)
	}
	r.InterceptMethod(ClientSide.SetUser, r.BEFORE)
	r.InterceptMethod(App.checkUser, r.BEFORE)
}
