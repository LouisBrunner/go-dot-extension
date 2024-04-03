// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Light3D struct {
  VisualInstance3D
}

func (me *Light3D) BaseClass() string {
  return "Light3D"
}

func NewLight3D() *Light3D {
  str := StringNameFromStr("Light3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Light3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type Light3DParam int
const (
  Light3DParamParamEnergy Light3DParam = 0
  Light3DParamParamIndirectEnergy Light3DParam = 1
  Light3DParamParamVolumetricFogEnergy Light3DParam = 2
  Light3DParamParamSpecular Light3DParam = 3
  Light3DParamParamRange Light3DParam = 4
  Light3DParamParamSize Light3DParam = 5
  Light3DParamParamAttenuation Light3DParam = 6
  Light3DParamParamSpotAngle Light3DParam = 7
  Light3DParamParamSpotAttenuation Light3DParam = 8
  Light3DParamParamShadowMaxDistance Light3DParam = 9
  Light3DParamParamShadowSplit1Offset Light3DParam = 10
  Light3DParamParamShadowSplit2Offset Light3DParam = 11
  Light3DParamParamShadowSplit3Offset Light3DParam = 12
  Light3DParamParamShadowFadeStart Light3DParam = 13
  Light3DParamParamShadowNormalBias Light3DParam = 14
  Light3DParamParamShadowBias Light3DParam = 15
  Light3DParamParamShadowPancakeSize Light3DParam = 16
  Light3DParamParamShadowOpacity Light3DParam = 17
  Light3DParamParamShadowBlur Light3DParam = 18
  Light3DParamParamTransmittanceBias Light3DParam = 19
  Light3DParamParamIntensity Light3DParam = 20
  Light3DParamParamMax Light3DParam = 21
)

type Light3DBakeMode int
const (
  Light3DBakeModeBakeDisabled Light3DBakeMode = 0
  Light3DBakeModeBakeStatic Light3DBakeMode = 1
  Light3DBakeModeBakeDynamic Light3DBakeMode = 2
)

func (me *Light3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Light3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Light3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Light3D) SetEditorOnly(editor_only bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&editor_only), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) IsEditorOnly() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetParam(param Light3DParam, value float64, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1722734213) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetParam(param Light3DParam, ) float64 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1844084987) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetShadow(enabled bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) HasShadow() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetNegative(enabled bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_negative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) IsNegative() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_negative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetCullMask(cull_mask int64, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mask), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetCullMask() int64 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetEnableDistanceFade(enable bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_distance_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) IsDistanceFadeEnabled() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_distance_fade_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetDistanceFadeBegin(distance float64, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetDistanceFadeBegin() float64 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetDistanceFadeShadow(distance float64, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetDistanceFadeShadow() float64 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetDistanceFadeLength(distance float64, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetDistanceFadeLength() float64 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetColor() Color {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Light3D) SetShadowReverseCullFace(enable bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_reverse_cull_face")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetShadowReverseCullFace() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_reverse_cull_face")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetBakeMode(bake_mode Light3DBakeMode, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 37739303) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bake_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetBakeMode() Light3DBakeMode {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 371737608) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret Light3DBakeMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Light3D) SetProjector(projector Texture2D, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_projector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(projector.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetProjector() Texture2D {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_projector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Light3D) SetTemperature(temperature float64, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_temperature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&temperature), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetTemperature() float64 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_temperature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) GetCorrelatedColor() Color {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_correlated_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
