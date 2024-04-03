// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Gradient struct {
  Resource
}

func (me *Gradient) BaseClass() string {
  return "Gradient"
}

func NewGradient() *Gradient {
  str := StringNameFromStr("Gradient") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Gradient{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type GradientInterpolationMode int
const (
  GradientInterpolationModeGradientInterpolateLinear GradientInterpolationMode = 0
  GradientInterpolationModeGradientInterpolateConstant GradientInterpolationMode = 1
  GradientInterpolationModeGradientInterpolateCubic GradientInterpolationMode = 2
)

type GradientColorSpace int
const (
  GradientColorSpaceGradientColorSpaceSrgb GradientColorSpace = 0
  GradientColorSpaceGradientColorSpaceLinearSrgb GradientColorSpace = 1
  GradientColorSpaceGradientColorSpaceOklab GradientColorSpace = 2
)

func (me *Gradient) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Gradient) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Gradient) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Gradient) AddPoint(offset float64, color Color, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3629403827) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) RemovePoint(point int64, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) SetOffset(point int64, offset float64, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), gdc.ConstTypePtr(&offset), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetOffset(point int64, ) float64 {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4025615559) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Gradient) Reverse()  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reverse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) SetColor(point int64, color Color, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetColor(point int64, ) Color {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2624840992) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) Sample(offset float64, ) Color {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1250405064) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) GetPointCount() int64 {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Gradient) SetOffsets(offsets PackedFloat32Array, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offsets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offsets.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetOffsets() PackedFloat32Array {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offsets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) SetColors(colors PackedColorArray, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3546319833) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(colors.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetColors() PackedColorArray {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedColorArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) SetInterpolationMode(interpolation_mode GradientInterpolationMode, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interpolation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1971444490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interpolation_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetInterpolationMode() GradientInterpolationMode {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interpolation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3674172981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret GradientInterpolationMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Gradient) SetInterpolationColorSpace(interpolation_color_space GradientColorSpace, )  {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interpolation_color_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3685995981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interpolation_color_space), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetInterpolationColorSpace() GradientColorSpace {
  classNameV := StringNameFromStr("Gradient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interpolation_color_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1538296000) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret GradientColorSpace

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
