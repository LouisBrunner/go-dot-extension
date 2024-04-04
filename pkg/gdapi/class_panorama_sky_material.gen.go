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

type PanoramaSkyMaterial struct {
  Material
}

func (me *PanoramaSkyMaterial) BaseClass() string {
  return "PanoramaSkyMaterial"
}

func NewPanoramaSkyMaterial() *PanoramaSkyMaterial {
  str := StringNameFromStr("PanoramaSkyMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PanoramaSkyMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PanoramaSkyMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PanoramaSkyMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PanoramaSkyMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PanoramaSkyMaterial) SetPanorama(texture Texture2D, )  {
  classNameV := StringNameFromStr("PanoramaSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_panorama")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PanoramaSkyMaterial) GetPanorama() Texture2D {
  classNameV := StringNameFromStr("PanoramaSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_panorama")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PanoramaSkyMaterial) SetFilteringEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PanoramaSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filtering_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PanoramaSkyMaterial) IsFilteringEnabled() bool {
  classNameV := StringNameFromStr("PanoramaSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_filtering_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
