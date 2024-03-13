// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Texture2DRD struct {
  obj gdc.ObjectPtr
}

func (me *Texture2DRD) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture2DRD) BaseClass() string {
  return "Texture2DRD"
}



// Enums

func (me *Texture2DRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture2DRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture2DRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Texture2DRD) SetTextureRdRid(texture_rd_rid RID, )  {
  classNameV := StringNameFromStr("Texture2DRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_rd_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture_rd_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Texture2DRD) GetTextureRdRid() RID {
  classNameV := StringNameFromStr("Texture2DRD")
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
