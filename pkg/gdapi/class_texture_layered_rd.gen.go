// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureLayeredRD struct {
  obj gdc.ObjectPtr
}

func (me *TextureLayeredRD) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureLayeredRD) BaseClass() string {
  return "TextureLayeredRD"
}



// Enums

func (me *TextureLayeredRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureLayeredRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureLayeredRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextureLayeredRD) SetTextureRdRid(texture_rd_rid RID, )  {
  classNameV := StringNameFromStr("TextureLayeredRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_rd_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture_rd_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureLayeredRD) GetTextureRdRid() RID {
  classNameV := StringNameFromStr("TextureLayeredRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_rd_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
