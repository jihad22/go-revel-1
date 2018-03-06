package controllers

import (
	"fmt"
	r "github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"time"
	"web-sid/app"
	"web-sid/app/models"
	"web-sid/app/routes"
)

type ClientSide struct {
	*r.Controller
}

func (c ClientSide) SetUser() r.Result {
	user := c.connected()
	if user != nil {
		c.ViewArgs["user"] = user
	}
	return nil
}

func (c ClientSide) Index() r.Result {
	t := time.Now()
	c.ViewArgs["myTime"] = t.Format("_2, January, 2006")
	index := "index"
	return c.Render(index)
}

func (c ClientSide) LoginPage() r.Result {
	loginPage := "loginPage"
	c.ViewArgs["logTitle"] = "KORLEKO UTARA"
	return c.Render(loginPage)
}

func (c ClientSide) getUser(username string) *models.User {
	var user models.User
	users := app.GORM.Where("username=?", username).Find(&user)
	if users.Error != nil {
		fmt.Println("RECORD NOT FOUND !")
	}
	return &user
}

func (c ClientSide) connected() *models.User {
	if username, ok := c.Session["USER_SESS"]; ok {
		return c.getUser(username)
	}
	if c.ViewArgs["USER_SESS"] != nil {
		return c.ViewArgs["USER_SESS"].(*models.User)
	}
	return nil
}

func (c ClientSide) Login(username, password string, remember bool) r.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err == nil {
			c.Session["USER_SESS"] = user.Username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Wellcome %s", username)
			return c.Redirect(routes.App.AdminPage())
		}
	}
	c.Flash.Out["username"] = user.Username
	c.Flash.Error("Username Atau Password Salah")
	return c.Redirect(routes.ClientSide.LoginPage())
}

// LOGOUT
func (c ClientSide) LogOut() r.Result {
	for s := range c.Session {
		delete(c.Session, s)
	}
	return c.Redirect(routes.ClientSide.LoginPage())
}
