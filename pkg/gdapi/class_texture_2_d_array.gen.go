// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Texture2DArray struct {
  obj gdc.ObjectPtr
}

func (me *Texture2DArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture2DArray) BaseClass() string {
  return "Texture2DArray"
}



// Enums

func (me *Texture2DArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture2DArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture2DArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Texture2DArray) CreatePlaceholder() Resource {
  classNameV := StringNameFromStr("Texture2DArray")
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