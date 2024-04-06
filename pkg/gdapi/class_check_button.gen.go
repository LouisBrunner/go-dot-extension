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

type ptrsForCheckButtonList struct {
}

var ptrsForCheckButton ptrsForCheckButtonList

func initCheckButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CheckButton")
	defer className.Destroy()

}

type CheckButton struct {
	Button
}

func (me *CheckButton) BaseClass() string {
	return "CheckButton"
}

func NewCheckButton() *CheckButton {
	str := StringNameFromStr("CheckButton") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CheckButton{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CheckButton) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CheckButton) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CheckButton) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
