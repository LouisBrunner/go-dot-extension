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

type ptrsForTextServerDummyList struct {
}

var ptrsForTextServerDummy ptrsForTextServerDummyList

func initTextServerDummyPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextServerDummy")
	defer className.Destroy()
}

type TextServerDummy struct {
	TextServerExtension
}

func (me *TextServerDummy) BaseClass() string {
	return "TextServerDummy"
}

func NewTextServerDummy() *TextServerDummy {
	str := StringNameFromStr("TextServerDummy") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextServerDummy{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TextServerDummy) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextServerDummy) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextServerDummy) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
