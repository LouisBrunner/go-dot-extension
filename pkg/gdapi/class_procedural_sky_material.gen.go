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

type ProceduralSkyMaterial struct {
  Material
}

func (me *ProceduralSkyMaterial) BaseClass() string {
  return "ProceduralSkyMaterial"
}

func NewProceduralSkyMaterial() *ProceduralSkyMaterial {
  str := StringNameFromStr("ProceduralSkyMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ProceduralSkyMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyTopColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_top_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetSkyHorizonColor(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyHorizonColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetSkyCurve(curve float64, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyCurve() float64 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSkyEnergyMultiplier(multiplier float64, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyEnergyMultiplier() float64 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSkyCover(sky_cover Texture2D, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_cover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{sky_cover.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyCover() Texture2D {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_cover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetSkyCoverModulate(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_cover_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyCoverModulate() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_cover_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetGroundBottomColor(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_bottom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundBottomColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_bottom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetGroundHorizonColor(color Color, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundHorizonColor() Color {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_horizon_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetGroundCurve(curve float64, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundCurve() float64 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetGroundEnergyMultiplier(energy float64, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundEnergyMultiplier() float64 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSunAngleMax(degrees float64, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sun_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSunAngleMax() float64 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sun_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSunCurve(curve float64, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sun_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSunCurve() float64 {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sun_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetUseDebanding(use_debanding bool, )  {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_debanding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_debanding) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetUseDebanding() bool {
  classNameV := StringNameFromStr("ProceduralSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_debanding")
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
