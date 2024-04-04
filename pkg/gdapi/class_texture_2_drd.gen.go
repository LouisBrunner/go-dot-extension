// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type Texture2DRD struct {
  Texture2D
}

func (me *Texture2DRD) BaseClass() string {
  return "Texture2DRD"
}

func NewTexture2DRD() *Texture2DRD {
  str := StringNameFromStr("Texture2DRD") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Texture2DRD{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{texture_rd_rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture2DRD) GetTextureRdRid() RID {
  classNameV := StringNameFromStr("Texture2DRD")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_rd_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
