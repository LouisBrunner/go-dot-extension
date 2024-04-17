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

type ptrsForRenderDataRDList struct {
}

var ptrsForRenderDataRD ptrsForRenderDataRDList

func initRenderDataRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderDataRD")
	defer className.Destroy()

}

type RenderDataRD struct {
	RenderData
}

func (me *RenderDataRD) BaseClass() string {
	return "RenderDataRD"
}

func NewRenderDataRD() *RenderDataRD {
	str := StringNameFromStr("RenderDataRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderDataRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderDataRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderDataRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderDataRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
