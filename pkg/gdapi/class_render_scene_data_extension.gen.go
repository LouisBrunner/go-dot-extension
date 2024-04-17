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

type ptrsForRenderSceneDataExtensionList struct {
	fnXGetCamTransform   gdc.MethodBindPtr
	fnXGetCamProjection  gdc.MethodBindPtr
	fnXGetViewCount      gdc.MethodBindPtr
	fnXGetViewEyeOffset  gdc.MethodBindPtr
	fnXGetViewProjection gdc.MethodBindPtr
	fnXGetUniformBuffer  gdc.MethodBindPtr
}

var ptrsForRenderSceneDataExtension ptrsForRenderSceneDataExtensionList

func initRenderSceneDataExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderSceneDataExtension")
	defer className.Destroy()

}

type RenderSceneDataExtension struct {
	RenderSceneData
}

func (me *RenderSceneDataExtension) BaseClass() string {
	return "RenderSceneDataExtension"
}

func NewRenderSceneDataExtension() *RenderSceneDataExtension {
	str := StringNameFromStr("RenderSceneDataExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderSceneDataExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderSceneDataExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderSceneDataExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderSceneDataExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
