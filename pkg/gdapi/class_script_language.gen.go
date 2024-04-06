// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForScriptLanguageList struct {
}

var ptrsForScriptLanguage ptrsForScriptLanguageList

func initScriptLanguagePtrs(iface gdc.Interface) {

	className := StringNameFromStr("ScriptLanguage")
	defer className.Destroy()
}

type ScriptLanguage struct {
	Object
}

func (me *ScriptLanguage) BaseClass() string {
	return "ScriptLanguage"
}

func NewScriptLanguage() *ScriptLanguage {
	str := StringNameFromStr("ScriptLanguage") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ScriptLanguage{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ScriptLanguage) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ScriptLanguage) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ScriptLanguage) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
