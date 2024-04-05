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

type ptrsForLight3DList struct {
  fnSetEditorOnly gdc.MethodBindPtr
  fnIsEditorOnly gdc.MethodBindPtr
  fnSetParam gdc.MethodBindPtr
  fnGetParam gdc.MethodBindPtr
  fnSetShadow gdc.MethodBindPtr
  fnHasShadow gdc.MethodBindPtr
  fnSetNegative gdc.MethodBindPtr
  fnIsNegative gdc.MethodBindPtr
  fnSetCullMask gdc.MethodBindPtr
  fnGetCullMask gdc.MethodBindPtr
  fnSetEnableDistanceFade gdc.MethodBindPtr
  fnIsDistanceFadeEnabled gdc.MethodBindPtr
  fnSetDistanceFadeBegin gdc.MethodBindPtr
  fnGetDistanceFadeBegin gdc.MethodBindPtr
  fnSetDistanceFadeShadow gdc.MethodBindPtr
  fnGetDistanceFadeShadow gdc.MethodBindPtr
  fnSetDistanceFadeLength gdc.MethodBindPtr
  fnGetDistanceFadeLength gdc.MethodBindPtr
  fnSetColor gdc.MethodBindPtr
  fnGetColor gdc.MethodBindPtr
  fnSetShadowReverseCullFace gdc.MethodBindPtr
  fnGetShadowReverseCullFace gdc.MethodBindPtr
  fnSetBakeMode gdc.MethodBindPtr
  fnGetBakeMode gdc.MethodBindPtr
  fnSetProjector gdc.MethodBindPtr
  fnGetProjector gdc.MethodBindPtr
  fnSetTemperature gdc.MethodBindPtr
  fnGetTemperature gdc.MethodBindPtr
  fnGetCorrelatedColor gdc.MethodBindPtr
}

var ptrsForLight3D ptrsForLight3DList

func initLight3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Light3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_editor_only")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetEditorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_editor_only")
    defer methodName.Destroy()
    ptrsForLight3D.fnIsEditorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_param")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1722734213))
  }
  {
    methodName := StringNameFromStr("get_param")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1844084987))
  }
  {
    methodName := StringNameFromStr("set_shadow")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetShadow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("has_shadow")
    defer methodName.Destroy()
    ptrsForLight3D.fnHasShadow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_negative")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetNegative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_negative")
    defer methodName.Destroy()
    ptrsForLight3D.fnIsNegative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_cull_mask")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_cull_mask")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_enable_distance_fade")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetEnableDistanceFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_distance_fade_enabled")
    defer methodName.Destroy()
    ptrsForLight3D.fnIsDistanceFadeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_distance_fade_begin")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetDistanceFadeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_distance_fade_begin")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetDistanceFadeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_distance_fade_shadow")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetDistanceFadeShadow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_distance_fade_shadow")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetDistanceFadeShadow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_distance_fade_length")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetDistanceFadeLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_distance_fade_length")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetDistanceFadeLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_color")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_shadow_reverse_cull_face")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetShadowReverseCullFace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_shadow_reverse_cull_face")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetShadowReverseCullFace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_bake_mode")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetBakeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 37739303))
  }
  {
    methodName := StringNameFromStr("get_bake_mode")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetBakeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 371737608))
  }
  {
    methodName := StringNameFromStr("set_projector")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetProjector = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_projector")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetProjector = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_temperature")
    defer methodName.Destroy()
    ptrsForLight3D.fnSetTemperature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_temperature")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetTemperature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_correlated_color")
    defer methodName.Destroy()
    ptrsForLight3D.fnGetCorrelatedColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&editor_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetEditorOnly), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) IsEditorOnly() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnIsEditorOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetParam(param Light3DParam, value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetParam(param Light3DParam, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&param)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetShadow(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetShadow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) HasShadow() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnHasShadow), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetNegative(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetNegative), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) IsNegative() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnIsNegative), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetCullMask(cull_mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetCullMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetEnableDistanceFade(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetEnableDistanceFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) IsDistanceFadeEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnIsDistanceFadeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetDistanceFadeBegin(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetDistanceFadeBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetDistanceFadeBegin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetDistanceFadeBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetDistanceFadeShadow(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetDistanceFadeShadow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetDistanceFadeShadow() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetDistanceFadeShadow), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetDistanceFadeLength(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetDistanceFadeLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetDistanceFadeLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetDistanceFadeLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Light3D) SetShadowReverseCullFace(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetShadowReverseCullFace), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetShadowReverseCullFace() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetShadowReverseCullFace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) SetBakeMode(bake_mode Light3DBakeMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bake_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetBakeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetBakeMode() Light3DBakeMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Light3DBakeMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetBakeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Light3D) SetProjector(projector Texture2D, )  {
  cargs := []gdc.ConstTypePtr{projector.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetProjector), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetProjector() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetProjector), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Light3D) SetTemperature(temperature float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&temperature) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnSetTemperature), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light3D) GetTemperature() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetTemperature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light3D) GetCorrelatedColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight3D.fnGetCorrelatedColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
