// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ProceduralSkyMaterial struct {
  obj gdc.ObjectPtr
}

func (me *ProceduralSkyMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ProceduralSkyMaterial) BaseClass() string {
  return "ProceduralSkyMaterial"
}



// Enums

func (me *ProceduralSkyMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ProceduralSkyMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ProceduralSkyMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ProceduralSkyMaterial) SetSkyTopColor(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_top_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSkyTopColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_top_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetSkyHorizonColor(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSkyHorizonColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetSkyCurve(curve float32, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSkyCurve() float32 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetSkyEnergyMultiplier(multiplier float32, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSkyEnergyMultiplier() float32 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetSkyCover(sky_cover Texture2D, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_cover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(sky_cover.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSkyCover() Texture2D {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_cover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetSkyCoverModulate(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_cover_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSkyCoverModulate() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_cover_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetGroundBottomColor(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_bottom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetGroundBottomColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_bottom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetGroundHorizonColor(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetGroundHorizonColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetGroundCurve(curve float32, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetGroundCurve() float32 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetGroundEnergyMultiplier(energy float32, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetGroundEnergyMultiplier() float32 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetSunAngleMax(degrees float32, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sun_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSunAngleMax() float32 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sun_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetSunCurve(curve float32, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sun_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetSunCurve() float32 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sun_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ProceduralSkyMaterial) SetUseDebanding(use_debanding bool, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_debanding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_debanding), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ProceduralSkyMaterial) GetUseDebanding() bool {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_debanding")
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
