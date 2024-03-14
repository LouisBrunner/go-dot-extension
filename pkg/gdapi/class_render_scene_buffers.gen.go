// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RenderSceneBuffers struct {
  RefCounted
}

func (me *RenderSceneBuffers) BaseClass() string {
  return "RenderSceneBuffers"
}



// Enums

func (me *RenderSceneBuffers) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RenderSceneBuffers) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RenderSceneBuffers) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RenderSceneBuffers) Configure(config RenderSceneBuffersConfiguration, )  {
  classNameV := StringNameFromStr("RenderSceneBuffers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("configure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3072623270) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(config.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
