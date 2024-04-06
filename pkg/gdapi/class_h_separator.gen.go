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

type ptrsForHSeparatorList struct {
}

var ptrsForHSeparator ptrsForHSeparatorList

func initHSeparatorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("HSeparator")
	defer className.Destroy()
}

type HSeparator struct {
	Separator
}

func (me *HSeparator) BaseClass() string {
	return "HSeparator"
}

func NewHSeparator() *HSeparator {
	str := StringNameFromStr("HSeparator") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &HSeparator{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *HSeparator) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *HSeparator) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *HSeparator) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
