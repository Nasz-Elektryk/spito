package lua

import (
	"github.com/nasz-elektryk/spito-rules/api"
	"github.com/yuin/gopher-lua"
	"reflect"
)

// Every api needs to be attached here in order to be available:
func attachApi(L *lua.LState) {
	var t = reflect.TypeOf

	setGlobalConstructor(L, "Package", t(api.Package{}))
	setGlobalFunction(L, "GetCurrentDistro", api.GetCurrentDistro)
	setGlobalFunction(L, "GetDaemon", api.GetDaemon)
	setGlobalFunction(L, "PathExists", api.PathExists)
	setGlobalFunction(L, "FileExists", api.FileExists)
	setGlobalFunction(L, "ReadFile", api.ReadFile)
	setGlobalFunction(L, "ReadDir", api.ReadDir)
}
