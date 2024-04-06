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

type ptrsForMarginContainerList struct {
}

var ptrsForMarginContainer ptrsForMarginContainerList

func initMarginContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MarginContainer")
	defer className.Destroy()

}

type MarginContainer struct {
	Container
}

func (me *MarginContainer) BaseClass() string {
	return "MarginContainer"
}

func NewMarginContainer() *MarginContainer {
	str := StringNameFromStr("MarginContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MarginContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MarginContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MarginContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MarginContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
