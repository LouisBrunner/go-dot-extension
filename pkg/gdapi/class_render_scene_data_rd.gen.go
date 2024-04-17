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

type ptrsForRenderSceneDataRDList struct {
}

var ptrsForRenderSceneDataRD ptrsForRenderSceneDataRDList

func initRenderSceneDataRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderSceneDataRD")
	defer className.Destroy()

}

type RenderSceneDataRD struct {
	RenderSceneData
}

func (me *RenderSceneDataRD) BaseClass() string {
	return "RenderSceneDataRD"
}

func NewRenderSceneDataRD() *RenderSceneDataRD {
	str := StringNameFromStr("RenderSceneDataRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderSceneDataRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderSceneDataRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderSceneDataRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderSceneDataRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
