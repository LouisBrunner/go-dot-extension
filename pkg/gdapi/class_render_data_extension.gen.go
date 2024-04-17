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

type ptrsForRenderDataExtensionList struct {
	fnXGetRenderSceneBuffers gdc.MethodBindPtr
	fnXGetRenderSceneData    gdc.MethodBindPtr
	fnXGetEnvironment        gdc.MethodBindPtr
	fnXGetCameraAttributes   gdc.MethodBindPtr
}

var ptrsForRenderDataExtension ptrsForRenderDataExtensionList

func initRenderDataExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderDataExtension")
	defer className.Destroy()

}

type RenderDataExtension struct {
	RenderData
}

func (me *RenderDataExtension) BaseClass() string {
	return "RenderDataExtension"
}

func NewRenderDataExtension() *RenderDataExtension {
	str := StringNameFromStr("RenderDataExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderDataExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderDataExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderDataExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderDataExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
