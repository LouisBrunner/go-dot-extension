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

type ptrsForNoiseTexture3DList struct {
  fnSetWidth gdc.MethodBindPtr
  fnSetHeight gdc.MethodBindPtr
  fnSetDepth gdc.MethodBindPtr
  fnSetInvert gdc.MethodBindPtr
  fnGetInvert gdc.MethodBindPtr
  fnSetSeamless gdc.MethodBindPtr
  fnGetSeamless gdc.MethodBindPtr
  fnSetSeamlessBlendSkirt gdc.MethodBindPtr
  fnGetSeamlessBlendSkirt gdc.MethodBindPtr
  fnSetNormalize gdc.MethodBindPtr
  fnIsNormalized gdc.MethodBindPtr
  fnSetColorRamp gdc.MethodBindPtr
  fnGetColorRamp gdc.MethodBindPtr
  fnSetNoise gdc.MethodBindPtr
  fnGetNoise gdc.MethodBindPtr
}

var ptrsForNoiseTexture3D ptrsForNoiseTexture3DList

func initNoiseTexture3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NoiseTexture3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_width")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_height")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_depth")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_invert")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_invert")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnGetInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_seamless")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetSeamless = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_seamless")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnGetSeamless = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_seamless_blend_skirt")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetSeamlessBlendSkirt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_seamless_blend_skirt")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnGetSeamlessBlendSkirt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_normalize")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetNormalize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_normalized")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnIsNormalized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_color_ramp")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
  }
  {
    methodName := StringNameFromStr("get_color_ramp")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnGetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
  }
  {
    methodName := StringNameFromStr("set_noise")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnSetNoise = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4135492439))
  }
  {
    methodName := StringNameFromStr("get_noise")
    defer methodName.Destroy()
    ptrsForNoiseTexture3D.fnGetNoise = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 185851837))
  }
}

type NoiseTexture3D struct {
  Texture3D
}

func (me *NoiseTexture3D) BaseClass() string {
  return "NoiseTexture3D"
}

func NewNoiseTexture3D() *NoiseTexture3D {
  str := StringNameFromStr("NoiseTexture3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NoiseTexture3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NoiseTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NoiseTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NoiseTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NoiseTexture3D) SetWidth(width int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) SetHeight(height int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) SetDepth(depth int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) SetInvert(invert bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetInvert), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) GetInvert() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnGetInvert), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture3D) SetSeamless(seamless bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetSeamless), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) GetSeamless() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnGetSeamless), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture3D) SetSeamlessBlendSkirt(seamless_blend_skirt float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless_blend_skirt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetSeamlessBlendSkirt), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) GetSeamlessBlendSkirt() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnGetSeamlessBlendSkirt), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture3D) SetNormalize(normalize bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetNormalize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) IsNormalized() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnIsNormalized), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture3D) SetColorRamp(gradient Gradient, )  {
  cargs := []gdc.ConstTypePtr{gradient.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetColorRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) GetColorRamp() Gradient {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnGetColorRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NoiseTexture3D) SetNoise(noise Noise, )  {
  cargs := []gdc.ConstTypePtr{noise.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnSetNoise), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture3D) GetNoise() Noise {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNoise()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture3D.fnGetNoise), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
