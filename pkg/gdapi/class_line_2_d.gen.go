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

type ptrsForLine2DList struct {
  fnSetPoints gdc.MethodBindPtr
  fnGetPoints gdc.MethodBindPtr
  fnSetPointPosition gdc.MethodBindPtr
  fnGetPointPosition gdc.MethodBindPtr
  fnGetPointCount gdc.MethodBindPtr
  fnAddPoint gdc.MethodBindPtr
  fnRemovePoint gdc.MethodBindPtr
  fnClearPoints gdc.MethodBindPtr
  fnSetClosed gdc.MethodBindPtr
  fnIsClosed gdc.MethodBindPtr
  fnSetWidth gdc.MethodBindPtr
  fnGetWidth gdc.MethodBindPtr
  fnSetCurve gdc.MethodBindPtr
  fnGetCurve gdc.MethodBindPtr
  fnSetDefaultColor gdc.MethodBindPtr
  fnGetDefaultColor gdc.MethodBindPtr
  fnSetGradient gdc.MethodBindPtr
  fnGetGradient gdc.MethodBindPtr
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
  fnSetTextureMode gdc.MethodBindPtr
  fnGetTextureMode gdc.MethodBindPtr
  fnSetJointMode gdc.MethodBindPtr
  fnGetJointMode gdc.MethodBindPtr
  fnSetBeginCapMode gdc.MethodBindPtr
  fnGetBeginCapMode gdc.MethodBindPtr
  fnSetEndCapMode gdc.MethodBindPtr
  fnGetEndCapMode gdc.MethodBindPtr
  fnSetSharpLimit gdc.MethodBindPtr
  fnGetSharpLimit gdc.MethodBindPtr
  fnSetRoundPrecision gdc.MethodBindPtr
  fnGetRoundPrecision gdc.MethodBindPtr
  fnSetAntialiased gdc.MethodBindPtr
  fnGetAntialiased gdc.MethodBindPtr
}

var ptrsForLine2D ptrsForLine2DList

func initLine2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Line2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_points")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
  }
  {
    methodName := StringNameFromStr("get_points")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
  }
  {
    methodName := StringNameFromStr("set_point_position")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
  }
  {
    methodName := StringNameFromStr("get_point_position")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
  }
  {
    methodName := StringNameFromStr("get_point_count")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("add_point")
    defer methodName.Destroy()
    ptrsForLine2D.fnAddPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2654014372))
  }
  {
    methodName := StringNameFromStr("remove_point")
    defer methodName.Destroy()
    ptrsForLine2D.fnRemovePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("clear_points")
    defer methodName.Destroy()
    ptrsForLine2D.fnClearPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_closed")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetClosed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_closed")
    defer methodName.Destroy()
    ptrsForLine2D.fnIsClosed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_width")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_width")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_curve")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
  }
  {
    methodName := StringNameFromStr("get_curve")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
  }
  {
    methodName := StringNameFromStr("set_default_color")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetDefaultColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_default_color")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetDefaultColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_gradient")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetGradient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
  }
  {
    methodName := StringNameFromStr("get_gradient")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetGradient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
  }
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_texture_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetTextureMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1952559516))
  }
  {
    methodName := StringNameFromStr("get_texture_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetTextureMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2341040722))
  }
  {
    methodName := StringNameFromStr("set_joint_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetJointMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 604292979))
  }
  {
    methodName := StringNameFromStr("get_joint_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetJointMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2546544037))
  }
  {
    methodName := StringNameFromStr("set_begin_cap_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetBeginCapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1669024546))
  }
  {
    methodName := StringNameFromStr("get_begin_cap_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetBeginCapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1107511441))
  }
  {
    methodName := StringNameFromStr("set_end_cap_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetEndCapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1669024546))
  }
  {
    methodName := StringNameFromStr("get_end_cap_mode")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetEndCapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1107511441))
  }
  {
    methodName := StringNameFromStr("set_sharp_limit")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetSharpLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sharp_limit")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetSharpLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_round_precision")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetRoundPrecision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_round_precision")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetRoundPrecision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_antialiased")
    defer methodName.Destroy()
    ptrsForLine2D.fnSetAntialiased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_antialiased")
    defer methodName.Destroy()
    ptrsForLine2D.fnGetAntialiased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type Line2D struct {
  Node2D
}

func (me *Line2D) BaseClass() string {
  return "Line2D"
}

func NewLine2D() *Line2D {
  str := StringNameFromStr("Line2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Line2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type Line2DLineJointMode int
const (
  Line2DLineJointModeLineJointSharp Line2DLineJointMode = 0
  Line2DLineJointModeLineJointBevel Line2DLineJointMode = 1
  Line2DLineJointModeLineJointRound Line2DLineJointMode = 2
)

type Line2DLineCapMode int
const (
  Line2DLineCapModeLineCapNone Line2DLineCapMode = 0
  Line2DLineCapModeLineCapBox Line2DLineCapMode = 1
  Line2DLineCapModeLineCapRound Line2DLineCapMode = 2
)

type Line2DLineTextureMode int
const (
  Line2DLineTextureModeLineTextureNone Line2DLineTextureMode = 0
  Line2DLineTextureModeLineTextureTile Line2DLineTextureMode = 1
  Line2DLineTextureModeLineTextureStretch Line2DLineTextureMode = 2
)

func (me *Line2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Line2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Line2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Line2D) SetPoints(points PackedVector2Array, )  {
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetPoints() PackedVector2Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Line2D) SetPointPosition(index int64, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetPointPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetPointPosition(index int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Line2D) GetPointCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Line2D) AddPoint(position Vector2, index int64, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnAddPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) RemovePoint(index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnRemovePoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) ClearPoints()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnClearPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) SetClosed(closed bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&closed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetClosed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) IsClosed() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnIsClosed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Line2D) SetWidth(width float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetWidth() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Line2D) SetCurve(curve Curve, )  {
  cargs := []gdc.ConstTypePtr{curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetCurve() Curve {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Line2D) SetDefaultColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetDefaultColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetDefaultColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetDefaultColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Line2D) SetGradient(color Gradient, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetGradient), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetGradient() Gradient {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetGradient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Line2D) SetTexture(texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Line2D) SetTextureMode(mode Line2DLineTextureMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetTextureMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetTextureMode() Line2DLineTextureMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Line2DLineTextureMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetTextureMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Line2D) SetJointMode(mode Line2DLineJointMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetJointMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetJointMode() Line2DLineJointMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Line2DLineJointMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetJointMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Line2D) SetBeginCapMode(mode Line2DLineCapMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetBeginCapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetBeginCapMode() Line2DLineCapMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Line2DLineCapMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetBeginCapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Line2D) SetEndCapMode(mode Line2DLineCapMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetEndCapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetEndCapMode() Line2DLineCapMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Line2DLineCapMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetEndCapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Line2D) SetSharpLimit(limit float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetSharpLimit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetSharpLimit() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetSharpLimit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Line2D) SetRoundPrecision(precision int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&precision) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetRoundPrecision), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetRoundPrecision() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetRoundPrecision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Line2D) SetAntialiased(antialiased bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnSetAntialiased), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Line2D) GetAntialiased() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLine2D.fnGetAntialiased), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
