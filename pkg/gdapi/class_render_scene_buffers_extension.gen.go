// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RenderSceneBuffersExtension struct {
  obj gdc.ObjectPtr
}

func (me *RenderSceneBuffersExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RenderSceneBuffersExtension) BaseClass() string {
  return "RenderSceneBuffersExtension"
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
