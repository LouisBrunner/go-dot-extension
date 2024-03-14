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

func  (me *FogMaterial) SetDensity(density float32, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FogMaterial) GetDensity() float32 {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FogMaterial) SetHeightFalloff(height_falloff float32, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height_falloff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height_falloff), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FogMaterial) GetHeightFalloff() float32 {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height_falloff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FogMaterial) SetEdgeFade(edge_fade float32, )  {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edge_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge_fade), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FogMaterial) GetEdgeFade() float32 {
  classNameV := StringNameFromStr("FogMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Texture3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
