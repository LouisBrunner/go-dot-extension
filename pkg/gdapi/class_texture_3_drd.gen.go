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

type ptrsForTexture3DRDList struct {
  fnSetTextureRdRid gdc.MethodBindPtr
  fnGetTextureRdRid gdc.MethodBindPtr
}

var ptrsForTexture3DRD ptrsForTexture3DRDList

func initTexture3DRDPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Texture3DRD")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_texture_rd_rid")
    defer methodName.Destroy()
    ptrsForTexture3DRD.fnSetTextureRdRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("get_texture_rd_rid")
    defer methodName.Destroy()
    ptrsForTexture3DRD.fnGetTextureRdRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
}

type Texture3DRD struct {
  Texture3D
}

func (me *Texture3DRD) BaseClass() string {
  return "Texture3DRD"
}

func NewTexture3DRD() *Texture3DRD {
  str := StringNameFromStr("Texture3DRD") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Texture3DRD{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Texture3DRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture3DRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture3DRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Texture3DRD) SetTextureRdRid(texture_rd_rid RID, )  {
  cargs := []gdc.ConstTypePtr{texture_rd_rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3DRD.fnSetTextureRdRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture3DRD) GetTextureRdRid() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3DRD.fnGetTextureRdRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
