// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CubemapArray struct {
  obj gdc.ObjectPtr
}

func (me *CubemapArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CubemapArray) BaseClass() string {
  return "CubemapArray"
}



// Enums

func (me *CubemapArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CubemapArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CubemapArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CubemapArray) CreatePlaceholder() Resource {
  classNameV := StringNameFromStr("CubemapArray")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121922552) // FIXME: should cache?
  var ret Resource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties