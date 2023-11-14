package checker

import (
	"github.com/nasz-elektryk/spito/api"
	"github.com/yuin/gopher-lua"
	"reflect"
)

// Every api needs to be attached here in order to be available:
func attachApi(L *lua.LState) {
	var t = reflect.TypeOf

	setGlobalConstructor(L, "Package", t(api.Package{}))
	setGlobalFunction(L, "GetDistro", api.GetDistro)
	setGlobalFunction(L, "GetDaemon", api.GetDaemon)
	setGlobalFunction(L, "PathExists", api.PathExists)
	setGlobalFunction(L, "FileExists", api.FileExists)
	setGlobalFunction(L, "ReadFile", api.ReadFile)
	setGlobalFunction(L, "ReadDir", api.ReadDir)
	setGlobalFunction(L, "FileContains", api.FileContains)
	setGlobalFunction(L, "RemoveComments", api.RemoveComments)
	setGlobalFunction(L, "Find", api.Find)
	setGlobalFunction(L, "FindAll", api.FindAll)
	setGlobalFunction(L, "GetProperLines", api.GetProperLines)
	setGlobalFunction(L, "GetInitSystem", api.GetInitSystem)
}