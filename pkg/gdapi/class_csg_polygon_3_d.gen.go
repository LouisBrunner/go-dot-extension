// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type CSGPolygon3D struct {
  CSGPrimitive3D
}

func (me *CSGPolygon3D) BaseClass() string {
  return "CSGPolygon3D"
}

func NewCSGPolygon3D() *CSGPolygon3D {
  str := StringNameFromStr("CSGPolygon3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CSGPolygon3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type CSGPolygon3DMode int
const (
  CSGPolygon3DModeModeDepth CSGPolygon3DMode = 0
  CSGPolygon3DModeModeSpin CSGPolygon3DMode = 1
  CSGPolygon3DModeModePath CSGPolygon3DMode = 2
)

type CSGPolygon3DPathRotation int
const (
  CSGPolygon3DPathRotationPathRotationPolygon CSGPolygon3DPathRotation = 0
  CSGPolygon3DPathRotationPathRotationPath CSGPolygon3DPathRotation = 1
  CSGPolygon3DPathRotationPathRotationPathFollow CSGPolygon3DPathRotation = 2
)

type CSGPolygon3DPathIntervalType int
const (
  CSGPolygon3DPathIntervalTypePathIntervalDistance CSGPolygon3DPathIntervalType = 0
  CSGPolygon3DPathIntervalTypePathIntervalSubdivide CSGPolygon3DPathIntervalType = 1
)

func (me *CSGPolygon3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGPolygon3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGPolygon3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGPolygon3D) SetPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CSGPolygon3D) SetMode(mode CSGPolygon3DMode, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3158377035) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetMode() CSGPolygon3DMode {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1201612222) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CSGPolygon3DMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CSGPolygon3D) SetDepth(depth float64, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetDepth() float64 {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetSpinDegrees(degrees float64, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spin_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetSpinDegrees() float64 {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spin_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetSpinSides(spin_sides int64, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spin_sides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spin_sides) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetSpinSides() int64 {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spin_sides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetPathNode(path NodePath, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetPathNode() NodePath {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CSGPolygon3D) SetPathIntervalType(interval_type CSGPolygon3DPathIntervalType, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_interval_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744240707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetPathIntervalType() CSGPolygon3DPathIntervalType {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_interval_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3434618397) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CSGPolygon3DPathIntervalType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CSGPolygon3D) SetPathInterval(interval float64, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetPathInterval() float64 {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetPathSimplifyAngle(degrees float64, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_simplify_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetPathSimplifyAngle() float64 {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_simplify_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetPathRotation(path_rotation CSGPolygon3DPathRotation, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1412947288) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetPathRotation() CSGPolygon3DPathRotation {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 647219346) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CSGPolygon3DPathRotation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CSGPolygon3D) SetPathLocal(enable bool, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) IsPathLocal() bool {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_path_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetPathContinuousU(enable bool, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_continuous_u")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) IsPathContinuousU() bool {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_path_continuous_u")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetPathUDistance(distance float64, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_u_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetPathUDistance() float64 {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_u_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetPathJoined(enable bool, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_joined")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) IsPathJoined() bool {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_path_joined")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGPolygon3D) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetMaterial() Material {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CSGPolygon3D) SetSmoothFaces(smooth_faces bool, )  {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_smooth_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth_faces) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGPolygon3D) GetSmoothFaces() bool {
  classNameV := StringNameFromStr("CSGPolygon3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_smooth_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
