// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PlaceholderTexture2D struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderTexture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderTexture2D) BaseClass() string {
  return "PlaceholderTexture2D"
}



// Enums

func (me *PlaceholderTexture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PlaceholderTexture2D) SetSize(size Vector2, )  {
  classNameV := StringNameFromStr("PlaceholderTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *PlaceholderTexture2D) GetPropSize() Vector2 {
  panic("TODO: implement")
}

func (me *PlaceholderTexture2D) SetPropSize(value Vector2) {
  panic("TODO: implement")
}