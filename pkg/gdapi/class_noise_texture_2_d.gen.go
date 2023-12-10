// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NoiseTexture2D struct {
  obj gdc.ObjectPtr
}

func (me *NoiseTexture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NoiseTexture2D) BaseClass() string {
  return "NoiseTexture2D"
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

func  (me *NoiseTexture2D) SetWidth(width int, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture2D) SetHeight(height int, )  {
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NoiseTexture2D) SetSeamlessBlendSkirt(seamless_blend_skirt float32, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seamless_blend_skirt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless_blend_skirt), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture2D) GetSeamlessBlendSkirt() float32 {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless_blend_skirt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NoiseTexture2D) SetBumpStrength(bump_strength float32, )  {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bump_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bump_strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture2D) GetBumpStrength() float32 {
  classNameV := StringNameFromStr("NoiseTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bump_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Noise
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *NoiseTexture2D) GetPropWidth() int {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropWidth(value int) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropHeight() int {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropHeight(value int) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropInvert() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropInvert(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropIn3DSpace() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropIn3DSpace(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropGenerateMipmaps() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropGenerateMipmaps(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropSeamless() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropSeamless(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropSeamlessBlendSkirt() float32 {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropSeamlessBlendSkirt(value float32) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropAsNormalMap() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropAsNormalMap(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropBumpStrength() float32 {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropBumpStrength(value float32) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropNormalize() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropNormalize(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropColorRamp() Gradient {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropColorRamp(value Gradient) {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) GetPropNoise() Noise {
  panic("TODO: implement")
}

func (me *NoiseTexture2D) SetPropNoise(value Noise) {
  panic("TODO: implement")
}