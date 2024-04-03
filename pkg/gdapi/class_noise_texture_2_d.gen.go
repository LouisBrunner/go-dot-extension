// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NoiseTexture2D struct {
  Texture2D
}

func (me *NoiseTexture2D) BaseClass() string {
  return "NoiseTexture2D"
}

func NewNoiseTexture2D() *NoiseTexture2D {
  str := StringNameFromStr("NoiseTexture2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NoiseTexture2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NoiseTexture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NoiseTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NoiseTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NoiseTexture2D) SetWidth(width int64, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) SetHeight(height int64, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) SetInvert(invert bool, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) GetInvert() bool {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetIn3DSpace(enable bool, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_in_3d_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) IsIn3DSpace() bool {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_3d_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetGenerateMipmaps(invert bool, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) IsGeneratingMipmaps() bool {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_generating_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetSeamless(seamless bool, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seamless")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) GetSeamless() bool {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetSeamlessBlendSkirt(seamless_blend_skirt float64, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seamless_blend_skirt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless_blend_skirt), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) GetSeamlessBlendSkirt() float64 {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless_blend_skirt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetAsNormalMap(as_normal_map bool, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_normal_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&as_normal_map), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) IsNormalMap() bool {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_normal_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetBumpStrength(bump_strength float64, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bump_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bump_strength), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) GetBumpStrength() float64 {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bump_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetNormalize(normalize bool, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normalize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) IsNormalized() bool {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NoiseTexture2D) SetColorRamp(gradient Gradient, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(gradient.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) GetColorRamp() Gradient {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NoiseTexture2D) SetNoise(noise Noise, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_noise")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4135492439) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(noise.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NoiseTexture2D) GetNoise() Noise {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 185851837) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNoise()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
