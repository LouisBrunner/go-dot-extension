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

type ptrsForCurveList struct {
  fnGetPointCount gdc.MethodBindPtr
  fnSetPointCount gdc.MethodBindPtr
  fnAddPoint gdc.MethodBindPtr
  fnRemovePoint gdc.MethodBindPtr
  fnClearPoints gdc.MethodBindPtr
  fnGetPointPosition gdc.MethodBindPtr
  fnSetPointValue gdc.MethodBindPtr
  fnSetPointOffset gdc.MethodBindPtr
  fnSample gdc.MethodBindPtr
  fnSampleBaked gdc.MethodBindPtr
  fnGetPointLeftTangent gdc.MethodBindPtr
  fnGetPointRightTangent gdc.MethodBindPtr
  fnGetPointLeftMode gdc.MethodBindPtr
  fnGetPointRightMode gdc.MethodBindPtr
  fnSetPointLeftTangent gdc.MethodBindPtr
  fnSetPointRightTangent gdc.MethodBindPtr
  fnSetPointLeftMode gdc.MethodBindPtr
  fnSetPointRightMode gdc.MethodBindPtr
  fnGetMinValue gdc.MethodBindPtr
  fnSetMinValue gdc.MethodBindPtr
  fnGetMaxValue gdc.MethodBindPtr
  fnSetMaxValue gdc.MethodBindPtr
  fnCleanDupes gdc.MethodBindPtr
  fnBake gdc.MethodBindPtr
  fnGetBakeResolution gdc.MethodBindPtr
  fnSetBakeResolution gdc.MethodBindPtr
}

var ptrsForCurve ptrsForCurveList

func initCurvePtrs(iface gdc.Interface) {

  className := StringNameFromStr("Curve")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_point_count")
    defer methodName.Destroy()
    ptrsForCurve.fnGetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_point_count")
    defer methodName.Destroy()
    ptrsForCurve.fnSetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("add_point")
    defer methodName.Destroy()
    ptrsForCurve.fnAddPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 434072736))
  }
  {
    methodName := StringNameFromStr("remove_point")
    defer methodName.Destroy()
    ptrsForCurve.fnRemovePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("clear_points")
    defer methodName.Destroy()
    ptrsForCurve.fnClearPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_point_position")
    defer methodName.Destroy()
    ptrsForCurve.fnGetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
  }
  {
    methodName := StringNameFromStr("set_point_value")
    defer methodName.Destroy()
    ptrsForCurve.fnSetPointValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("set_point_offset")
    defer methodName.Destroy()
    ptrsForCurve.fnSetPointOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780573764))
  }
  {
    methodName := StringNameFromStr("sample")
    defer methodName.Destroy()
    ptrsForCurve.fnSample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3919130443))
  }
  {
    methodName := StringNameFromStr("sample_baked")
    defer methodName.Destroy()
    ptrsForCurve.fnSampleBaked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3919130443))
  }
  {
    methodName := StringNameFromStr("get_point_left_tangent")
    defer methodName.Destroy()
    ptrsForCurve.fnGetPointLeftTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("get_point_right_tangent")
    defer methodName.Destroy()
    ptrsForCurve.fnGetPointRightTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("get_point_left_mode")
    defer methodName.Destroy()
    ptrsForCurve.fnGetPointLeftMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 426950354))
  }
  {
    methodName := StringNameFromStr("get_point_right_mode")
    defer methodName.Destroy()
    ptrsForCurve.fnGetPointRightMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 426950354))
  }
  {
    methodName := StringNameFromStr("set_point_left_tangent")
    defer methodName.Destroy()
    ptrsForCurve.fnSetPointLeftTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("set_point_right_tangent")
    defer methodName.Destroy()
    ptrsForCurve.fnSetPointRightTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("set_point_left_mode")
    defer methodName.Destroy()
    ptrsForCurve.fnSetPointLeftMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1217242874))
  }
  {
    methodName := StringNameFromStr("set_point_right_mode")
    defer methodName.Destroy()
    ptrsForCurve.fnSetPointRightMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1217242874))
  }
  {
    methodName := StringNameFromStr("get_min_value")
    defer methodName.Destroy()
    ptrsForCurve.fnGetMinValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_min_value")
    defer methodName.Destroy()
    ptrsForCurve.fnSetMinValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_value")
    defer methodName.Destroy()
    ptrsForCurve.fnGetMaxValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_value")
    defer methodName.Destroy()
    ptrsForCurve.fnSetMaxValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("clean_dupes")
    defer methodName.Destroy()
    ptrsForCurve.fnCleanDupes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bake")
    defer methodName.Destroy()
    ptrsForCurve.fnBake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_bake_resolution")
    defer methodName.Destroy()
    ptrsForCurve.fnGetBakeResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_bake_resolution")
    defer methodName.Destroy()
    ptrsForCurve.fnSetBakeResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
}

type Curve struct {
  Resource
}

func (me *Curve) BaseClass() string {
  return "Curve"
}

func NewCurve() *Curve {
  str := StringNameFromStr("Curve") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Curve{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type CurveTangentMode int
const (
  CurveTangentModeTangentFree CurveTangentMode = 0
  CurveTangentModeTangentLinear CurveTangentMode = 1
  CurveTangentModeTangentModeCount CurveTangentMode = 2
)

func (me *Curve) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Curve) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Curve) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Curve) GetPointCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) SetPointCount(count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetPointCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) AddPoint(position Vector2, left_tangent float64, right_tangent float64, left_mode CurveTangentMode, right_mode CurveTangentMode, ) int64 {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&left_tangent) , gdc.ConstTypePtr(&right_tangent) , gdc.ConstTypePtr(&left_mode) , gdc.ConstTypePtr(&right_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&left_tangent)
  pinner.Pin(&right_tangent)
  pinner.Pin(&left_mode)
  pinner.Pin(&right_mode)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnAddPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) RemovePoint(index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnRemovePoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) ClearPoints()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnClearPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) GetPointPosition(index int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve) SetPointValue(index int64, y float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&y) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetPointValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) SetPointOffset(index int64, offset float64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&index)
  pinner.Pin(&offset)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetPointOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) Sample(offset float64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&offset)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSample), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) SampleBaked(offset float64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&offset)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSampleBaked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) GetPointLeftTangent(index int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetPointLeftTangent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) GetPointRightTangent(index int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetPointRightTangent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) GetPointLeftMode(index int64, ) CurveTangentMode {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CurveTangentMode
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetPointLeftMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Curve) GetPointRightMode(index int64, ) CurveTangentMode {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CurveTangentMode
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetPointRightMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Curve) SetPointLeftTangent(index int64, tangent float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&tangent) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetPointLeftTangent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) SetPointRightTangent(index int64, tangent float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&tangent) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetPointRightTangent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) SetPointLeftMode(index int64, mode CurveTangentMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetPointLeftMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) SetPointRightMode(index int64, mode CurveTangentMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetPointRightMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) GetMinValue() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetMinValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) SetMinValue(min float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetMinValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) GetMaxValue() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetMaxValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) SetMaxValue(max float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetMaxValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) CleanDupes()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnCleanDupes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) Bake()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnBake), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve) GetBakeResolution() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnGetBakeResolution), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve) SetBakeResolution(resolution int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resolution) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve.fnSetBakeResolution), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CurveRangeChangedSignalFn func()

func (me *Curve) ConnectRangeChanged(subs SignalSubscribers, fn CurveRangeChangedSignalFn) {
  sig := StringNameFromStr("range_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Curve) DisconnectRangeChanged(subs SignalSubscribers, fn CurveRangeChangedSignalFn) {
  sig := StringNameFromStr("range_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
