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

type ptrsForGradientList struct {
  fnAddPoint gdc.MethodBindPtr
  fnRemovePoint gdc.MethodBindPtr
  fnSetOffset gdc.MethodBindPtr
  fnGetOffset gdc.MethodBindPtr
  fnReverse gdc.MethodBindPtr
  fnSetColor gdc.MethodBindPtr
  fnGetColor gdc.MethodBindPtr
  fnSample gdc.MethodBindPtr
  fnGetPointCount gdc.MethodBindPtr
  fnSetOffsets gdc.MethodBindPtr
  fnGetOffsets gdc.MethodBindPtr
  fnSetColors gdc.MethodBindPtr
  fnGetColors gdc.MethodBindPtr
  fnSetInterpolationMode gdc.MethodBindPtr
  fnGetInterpolationMode gdc.MethodBindPtr
  fnSetInterpolationColorSpace gdc.MethodBindPtr
  fnGetInterpolationColorSpace gdc.MethodBindPtr
}

var ptrsForGradient ptrsForGradientList

func initGradientPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Gradient")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_point")
    defer methodName.Destroy()
    ptrsForGradient.fnAddPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3629403827))
  }
  {
    methodName := StringNameFromStr("remove_point")
    defer methodName.Destroy()
    ptrsForGradient.fnRemovePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_offset")
    defer methodName.Destroy()
    ptrsForGradient.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_offset")
    defer methodName.Destroy()
    ptrsForGradient.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4025615559))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForGradient.fnReverse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForGradient.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
  }
  {
    methodName := StringNameFromStr("get_color")
    defer methodName.Destroy()
    ptrsForGradient.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2624840992))
  }
  {
    methodName := StringNameFromStr("sample")
    defer methodName.Destroy()
    ptrsForGradient.fnSample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1250405064))
  }
  {
    methodName := StringNameFromStr("get_point_count")
    defer methodName.Destroy()
    ptrsForGradient.fnGetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_offsets")
    defer methodName.Destroy()
    ptrsForGradient.fnSetOffsets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
  }
  {
    methodName := StringNameFromStr("get_offsets")
    defer methodName.Destroy()
    ptrsForGradient.fnGetOffsets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675695659))
  }
  {
    methodName := StringNameFromStr("set_colors")
    defer methodName.Destroy()
    ptrsForGradient.fnSetColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3546319833))
  }
  {
    methodName := StringNameFromStr("get_colors")
    defer methodName.Destroy()
    ptrsForGradient.fnGetColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1392750486))
  }
  {
    methodName := StringNameFromStr("set_interpolation_mode")
    defer methodName.Destroy()
    ptrsForGradient.fnSetInterpolationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1971444490))
  }
  {
    methodName := StringNameFromStr("get_interpolation_mode")
    defer methodName.Destroy()
    ptrsForGradient.fnGetInterpolationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3674172981))
  }
  {
    methodName := StringNameFromStr("set_interpolation_color_space")
    defer methodName.Destroy()
    ptrsForGradient.fnSetInterpolationColorSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3685995981))
  }
  {
    methodName := StringNameFromStr("get_interpolation_color_space")
    defer methodName.Destroy()
    ptrsForGradient.fnGetInterpolationColorSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1538296000))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnAddPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) RemovePoint(point int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnRemovePoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) SetOffset(point int64, offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetOffset(point int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Gradient) Reverse()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnReverse), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) SetColor(point int64, color Color, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetColor(point int64, ) Color {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) Sample(offset float64, ) Color {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&offset)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnSample), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) GetPointCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnGetPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Gradient) SetOffsets(offsets PackedFloat32Array, )  {
  cargs := []gdc.ConstTypePtr{offsets.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnSetOffsets), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetOffsets() PackedFloat32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnGetOffsets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) SetColors(colors PackedColorArray, )  {
  cargs := []gdc.ConstTypePtr{colors.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnSetColors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetColors() PackedColorArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedColorArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnGetColors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Gradient) SetInterpolationMode(interpolation_mode GradientInterpolationMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interpolation_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnSetInterpolationMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetInterpolationMode() GradientInterpolationMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GradientInterpolationMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnGetInterpolationMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Gradient) SetInterpolationColorSpace(interpolation_color_space GradientColorSpace, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interpolation_color_space) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnSetInterpolationColorSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Gradient) GetInterpolationColorSpace() GradientColorSpace {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GradientColorSpace

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradient.fnGetInterpolationColorSpace), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
