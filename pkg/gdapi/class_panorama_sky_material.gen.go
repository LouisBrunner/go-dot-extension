// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PanoramaSkyMaterial struct {
  obj gdc.ObjectPtr
}

func (me *PanoramaSkyMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PanoramaSkyMaterial) BaseClass() string {
  return "PanoramaSkyMaterial"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PanoramaSkyMaterial) GetPanorama() Texture2D {
  classNameV := StringNameFromStr("PanoramaSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_panorama")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PanoramaSkyMaterial) SetFilteringEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PanoramaSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filtering_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PanoramaSkyMaterial) IsFilteringEnabled() bool {
  classNameV := StringNameFromStr("PanoramaSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_filtering_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
