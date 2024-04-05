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

type ptrsForCurve3DList struct {
  fnGetPointCount gdc.MethodBindPtr
  fnSetPointCount gdc.MethodBindPtr
  fnAddPoint gdc.MethodBindPtr
  fnSetPointPosition gdc.MethodBindPtr
  fnGetPointPosition gdc.MethodBindPtr
  fnSetPointTilt gdc.MethodBindPtr
  fnGetPointTilt gdc.MethodBindPtr
  fnSetPointIn gdc.MethodBindPtr
  fnGetPointIn gdc.MethodBindPtr
  fnSetPointOut gdc.MethodBindPtr
  fnGetPointOut gdc.MethodBindPtr
  fnRemovePoint gdc.MethodBindPtr
  fnClearPoints gdc.MethodBindPtr
  fnSample gdc.MethodBindPtr
  fnSamplef gdc.MethodBindPtr
  fnSetBakeInterval gdc.MethodBindPtr
  fnGetBakeInterval gdc.MethodBindPtr
  fnSetUpVectorEnabled gdc.MethodBindPtr
  fnIsUpVectorEnabled gdc.MethodBindPtr
  fnGetBakedLength gdc.MethodBindPtr
  fnSampleBaked gdc.MethodBindPtr
  fnSampleBakedWithRotation gdc.MethodBindPtr
  fnSampleBakedUpVector gdc.MethodBindPtr
  fnGetBakedPoints gdc.MethodBindPtr
  fnGetBakedTilts gdc.MethodBindPtr
  fnGetBakedUpVectors gdc.MethodBindPtr
  fnGetClosestPoint gdc.MethodBindPtr
  fnGetClosestOffset gdc.MethodBindPtr
  fnTessellate gdc.MethodBindPtr
  fnTessellateEvenLength gdc.MethodBindPtr
}

var ptrsForCurve3D ptrsForCurve3DList

func initCurve3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Curve3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_point_count")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_point_count")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("add_point")
    defer methodName.Destroy()
    ptrsForCurve3D.fnAddPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2931053748))
  }
  {
    methodName := StringNameFromStr("set_point_position")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
  }
  {
    methodName := StringNameFromStr("get_point_position")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
  }
  {
    methodName := StringNameFromStr("set_point_tilt")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSetPointTilt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_point_tilt")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetPointTilt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("set_point_in")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSetPointIn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
  }
  {
    methodName := StringNameFromStr("get_point_in")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetPointIn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
  }
  {
    methodName := StringNameFromStr("set_point_out")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSetPointOut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
  }
  {
    methodName := StringNameFromStr("get_point_out")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetPointOut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
  }
  {
    methodName := StringNameFromStr("remove_point")
    defer methodName.Destroy()
    ptrsForCurve3D.fnRemovePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("clear_points")
    defer methodName.Destroy()
    ptrsForCurve3D.fnClearPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("sample")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3285246857))
  }
  {
    methodName := StringNameFromStr("samplef")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSamplef = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2553580215))
  }
  {
    methodName := StringNameFromStr("set_bake_interval")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSetBakeInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_bake_interval")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetBakeInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_up_vector_enabled")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSetUpVectorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_up_vector_enabled")
    defer methodName.Destroy()
    ptrsForCurve3D.fnIsUpVectorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_baked_length")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetBakedLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("sample_baked")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSampleBaked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1350085894))
  }
  {
    methodName := StringNameFromStr("sample_baked_with_rotation")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSampleBakedWithRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1939359131))
  }
  {
    methodName := StringNameFromStr("sample_baked_up_vector")
    defer methodName.Destroy()
    ptrsForCurve3D.fnSampleBakedUpVector = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1362627031))
  }
  {
    methodName := StringNameFromStr("get_baked_points")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetBakedPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
  }
  {
    methodName := StringNameFromStr("get_baked_tilts")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetBakedTilts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675695659))
  }
  {
    methodName := StringNameFromStr("get_baked_up_vectors")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetBakedUpVectors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
  }
  {
    methodName := StringNameFromStr("get_closest_point")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetClosestPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 192990374))
  }
  {
    methodName := StringNameFromStr("get_closest_offset")
    defer methodName.Destroy()
    ptrsForCurve3D.fnGetClosestOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1109078154))
  }
  {
    methodName := StringNameFromStr("tessellate")
    defer methodName.Destroy()
    ptrsForCurve3D.fnTessellate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1519759391))
  }
  {
    methodName := StringNameFromStr("tessellate_even_length")
    defer methodName.Destroy()
    ptrsForCurve3D.fnTessellateEvenLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 133237049))
  }
}

type Curve3D struct {
  Resource
}

func (me *Curve3D) BaseClass() string {
  return "Curve3D"
}

func NewCurve3D() *Curve3D {
  str := StringNameFromStr("Curve3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Curve3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Curve3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Curve3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Curve3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Curve3D) GetPointCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SetPointCount(count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSetPointCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) AddPoint(position Vector3, in Vector3, out Vector3, index int64, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), in.AsCTypePtr(), out.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnAddPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) SetPointPosition(idx int64, position Vector3, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSetPointPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointPosition(idx int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SetPointTilt(idx int64, tilt float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&tilt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSetPointTilt), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointTilt(idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetPointTilt), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SetPointIn(idx int64, position Vector3, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSetPointIn), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointIn(idx int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetPointIn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SetPointOut(idx int64, position Vector3, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSetPointOut), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointOut(idx int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetPointOut), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) RemovePoint(idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnRemovePoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) ClearPoints()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnClearPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) Sample(idx int64, t float64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&t) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)
  pinner.Pin(&t)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSample), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) Samplef(fofs float64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fofs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&fofs)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSamplef), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SetBakeInterval(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSetBakeInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetBakeInterval() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetBakeInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SetUpVectorEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSetUpVectorEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) IsUpVectorEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnIsUpVectorEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) GetBakedLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetBakedLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SampleBaked(offset float64, cubic bool, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&cubic) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&offset)
  pinner.Pin(&cubic)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSampleBaked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SampleBakedWithRotation(offset float64, cubic bool, apply_tilt bool, ) Transform3D {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&cubic) , gdc.ConstTypePtr(&apply_tilt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&offset)
  pinner.Pin(&cubic)
  pinner.Pin(&apply_tilt)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSampleBakedWithRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SampleBakedUpVector(offset float64, apply_tilt bool, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&apply_tilt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&offset)
  pinner.Pin(&apply_tilt)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnSampleBakedUpVector), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetBakedPoints() PackedVector3Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetBakedPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetBakedTilts() PackedFloat32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetBakedTilts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetBakedUpVectors() PackedVector3Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetBakedUpVectors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetClosestPoint(to_point Vector3, ) Vector3 {
  cargs := []gdc.ConstTypePtr{to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetClosestPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetClosestOffset(to_point Vector3, ) float64 {
  cargs := []gdc.ConstTypePtr{to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnGetClosestOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) Tessellate(max_stages int64, tolerance_degrees float64, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages) , gdc.ConstTypePtr(&tolerance_degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&max_stages)
  pinner.Pin(&tolerance_degrees)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnTessellate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) TessellateEvenLength(max_stages int64, tolerance_length float64, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages) , gdc.ConstTypePtr(&tolerance_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&max_stages)
  pinner.Pin(&tolerance_length)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve3D.fnTessellateEvenLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
