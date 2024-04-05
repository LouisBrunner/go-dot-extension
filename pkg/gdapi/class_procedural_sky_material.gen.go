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

type ptrsForProceduralSkyMaterialList struct {
  fnSetSkyTopColor gdc.MethodBindPtr
  fnGetSkyTopColor gdc.MethodBindPtr
  fnSetSkyHorizonColor gdc.MethodBindPtr
  fnGetSkyHorizonColor gdc.MethodBindPtr
  fnSetSkyCurve gdc.MethodBindPtr
  fnGetSkyCurve gdc.MethodBindPtr
  fnSetSkyEnergyMultiplier gdc.MethodBindPtr
  fnGetSkyEnergyMultiplier gdc.MethodBindPtr
  fnSetSkyCover gdc.MethodBindPtr
  fnGetSkyCover gdc.MethodBindPtr
  fnSetSkyCoverModulate gdc.MethodBindPtr
  fnGetSkyCoverModulate gdc.MethodBindPtr
  fnSetGroundBottomColor gdc.MethodBindPtr
  fnGetGroundBottomColor gdc.MethodBindPtr
  fnSetGroundHorizonColor gdc.MethodBindPtr
  fnGetGroundHorizonColor gdc.MethodBindPtr
  fnSetGroundCurve gdc.MethodBindPtr
  fnGetGroundCurve gdc.MethodBindPtr
  fnSetGroundEnergyMultiplier gdc.MethodBindPtr
  fnGetGroundEnergyMultiplier gdc.MethodBindPtr
  fnSetSunAngleMax gdc.MethodBindPtr
  fnGetSunAngleMax gdc.MethodBindPtr
  fnSetSunCurve gdc.MethodBindPtr
  fnGetSunCurve gdc.MethodBindPtr
  fnSetUseDebanding gdc.MethodBindPtr
  fnGetUseDebanding gdc.MethodBindPtr
}

var ptrsForProceduralSkyMaterial ptrsForProceduralSkyMaterialList

func initProceduralSkyMaterialPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ProceduralSkyMaterial")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_sky_top_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSkyTopColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_sky_top_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSkyTopColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_sky_horizon_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSkyHorizonColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_sky_horizon_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSkyHorizonColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_sky_curve")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSkyCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sky_curve")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSkyCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sky_energy_multiplier")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSkyEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sky_energy_multiplier")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSkyEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sky_cover")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSkyCover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_sky_cover")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSkyCover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_sky_cover_modulate")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSkyCoverModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_sky_cover_modulate")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSkyCoverModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_ground_bottom_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetGroundBottomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_ground_bottom_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetGroundBottomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_ground_horizon_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetGroundHorizonColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_ground_horizon_color")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetGroundHorizonColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_ground_curve")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetGroundCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ground_curve")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetGroundCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ground_energy_multiplier")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetGroundEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ground_energy_multiplier")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetGroundEnergyMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sun_angle_max")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSunAngleMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sun_angle_max")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSunAngleMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sun_curve")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetSunCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sun_curve")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetSunCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_use_debanding")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnSetUseDebanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_use_debanding")
    defer methodName.Destroy()
    ptrsForProceduralSkyMaterial.fnGetUseDebanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSkyTopColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyTopColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSkyTopColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetSkyHorizonColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSkyHorizonColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyHorizonColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSkyHorizonColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetSkyCurve(curve float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSkyCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyCurve() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSkyCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSkyEnergyMultiplier(multiplier float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSkyEnergyMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyEnergyMultiplier() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSkyEnergyMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSkyCover(sky_cover Texture2D, )  {
  cargs := []gdc.ConstTypePtr{sky_cover.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSkyCover), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyCover() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSkyCover), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetSkyCoverModulate(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSkyCoverModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSkyCoverModulate() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSkyCoverModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetGroundBottomColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetGroundBottomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundBottomColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetGroundBottomColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetGroundHorizonColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetGroundHorizonColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundHorizonColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetGroundHorizonColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProceduralSkyMaterial) SetGroundCurve(curve float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetGroundCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundCurve() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetGroundCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetGroundEnergyMultiplier(energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetGroundEnergyMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetGroundEnergyMultiplier() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetGroundEnergyMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSunAngleMax(degrees float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSunAngleMax), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSunAngleMax() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSunAngleMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetSunCurve(curve float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetSunCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetSunCurve() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetSunCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProceduralSkyMaterial) SetUseDebanding(use_debanding bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_debanding) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnSetUseDebanding), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProceduralSkyMaterial) GetUseDebanding() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForProceduralSkyMaterial.fnGetUseDebanding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
