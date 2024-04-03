// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type FogMaterial struct {
  Material
}

func (me *FogMaterial) BaseClass() string {
  return "FogMaterial"
}

func NewFogMaterial() *FogMaterial {
  str := StringNameFromStr("FogMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &FogMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *FogMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FogMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FogMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *FogMaterial) SetDensity(density float64, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogMaterial) GetDensity() float64 {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FogMaterial) SetAlbedo(albedo Color, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(albedo.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogMaterial) GetAlbedo() Color {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FogMaterial) SetEmission(emission Color, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(emission.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogMaterial) GetEmission() Color {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FogMaterial) SetHeightFalloff(height_falloff float64, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height_falloff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height_falloff), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogMaterial) GetHeightFalloff() float64 {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height_falloff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FogMaterial) SetEdgeFade(edge_fade float64, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edge_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge_fade), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogMaterial) GetEdgeFade() float64 {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FogMaterial) SetDensityTexture(density_texture Texture3D, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_density_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1188404210) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(density_texture.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogMaterial) GetDensityTexture() Texture3D {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_density_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373985333) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
