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

type ptrsForRenderSceneBuffersExtensionList struct {
  fnXConfigure gdc.MethodBindPtr
  fnXSetFsrSharpness gdc.MethodBindPtr
  fnXSetTextureMipmapBias gdc.MethodBindPtr
  fnXSetUseDebanding gdc.MethodBindPtr
}

var ptrsForRenderSceneBuffersExtension ptrsForRenderSceneBuffersExtensionList

func initRenderSceneBuffersExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RenderSceneBuffersExtension")
  defer className.Destroy()
}

type RenderSceneBuffersExtension struct {
  RenderSceneBuffers
}

func (me *RenderSceneBuffersExtension) BaseClass() string {
  return "RenderSceneBuffersExtension"
}

func NewRenderSceneBuffersExtension() *RenderSceneBuffersExtension {
  str := StringNameFromStr("RenderSceneBuffersExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RenderSceneBuffersExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RenderSceneBuffersExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RenderSceneBuffersExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RenderSceneBuffersExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
