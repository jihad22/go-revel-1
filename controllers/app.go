package controllers

import (
	"fmt"
	r "github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"web-sid/app"
	"web-sid/app/models"
	"web-sid/app/routes"
)

type App struct {
	ClientSide
}

func (c App) checkUser() r.Result {
	if user := c.connected(); &user == nil {
		c.Flash.Error("Mohon Untuk Login")
		return c.Redirect(routes.ClientSide.LogOut())
	}
	return nil
}

func (c App) AdminPage() r.Result {
	user := c.connected()
	if user == nil {

		c.Flash.Error("Mohon Untuk Login")
		return c.Redirect(routes.ClientSide.LogOut())
	}
	admin := strings.ToTitle(c.connected().Username)
	dashboard := "dashboard"
	c.ViewArgs["Nama"] = user.Nama
	c.ViewArgs["PageName"] = "Dashboard"
	return c.Render(admin, dashboard)
}

func (c App) DataAdmin() r.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Mohon Untuk Login")
		return c.Redirect(c.ClientSide.LogOut())
	}
	var list []models.User
	lists := app.GORM.Debug().Find(&list)
	if lists.Error != nil {
		panic(lists.Error)
	}
	admin := strings.ToTitle(c.connected().Username)
	dataAdmin := "dataadmin"
	c.ViewArgs["Nama"] = user.Nama
	c.ViewArgs["PageName"] = "Data Admin"
	return c.Render(list, admin, dataAdmin)
}

//tambah admin
func (c App) AddNewAdmin() r.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Mohon Untuk Login")
		return c.Redirect(routes.ClientSide.LogOut())
	}
	admin := strings.ToTitle(c.connected().Username)
	c.ViewArgs["Nama"] = user.Nama
	return c.Render(admin)
}

func (c App) AddingNewAdmin(usr models.User, password, confirmPassword string) r.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Mohon Untuk Login")
		return c.Redirect(routes.ClientSide.LogOut())
	}
	c.Validation.Required(password).Message("Kolom Password Harus Diisi")
	c.Validation.Required(confirmPassword).Message("Kolom Konfirmasi Harus Diisi")
	c.Validation.MinSize(password, 7)
	c.Validation.Required(confirmPassword == password).Message("Password tidak sama")
	usr.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		fmt.Println("PROSES GAGAL !!!")
		return c.Redirect(routes.App.AddNewAdmin())
	}
	fmt.Println("No Error")
	usr.Password, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := app.GORM.Save(&usr)
	if err.Error != nil {
		panic(err.Error)
	}
	return c.Redirect(routes.ClientSide.LoginPage())
}
