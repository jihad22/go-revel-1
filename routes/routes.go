// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) AdminPage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AdminPage", args).URL
}

func (_ tApp) DataAdmin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.DataAdmin", args).URL
}

func (_ tApp) AddNewAdmin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddNewAdmin", args).URL
}

func (_ tApp) AddingNewAdmin(
		usr interface{},
		password string,
		confirmPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "usr", usr)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "confirmPassword", confirmPassword)
	return revel.MainRouter.Reverse("App.AddingNewAdmin", args).URL
}


type tClientSide struct {}
var ClientSide tClientSide


func (_ tClientSide) SetUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ClientSide.SetUser", args).URL
}

func (_ tClientSide) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ClientSide.Index", args).URL
}

func (_ tClientSide) LoginPage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ClientSide.LoginPage", args).URL
}

func (_ tClientSide) Login(
		username string,
		password string,
		remember bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "remember", remember)
	return revel.MainRouter.Reverse("ClientSide.Login", args).URL
}

func (_ tClientSide) LogOut(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ClientSide.LogOut", args).URL
}


