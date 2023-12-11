// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AABB struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewAABB() AABB {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeAABB))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return AABB{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewAABBFromAABB(from AABB, ) AABB {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeAABB))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return AABB{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewAABBFromVector3Vector3(position Vector3, size Vector3, ) AABB {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeAABB))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr(), }))
  return AABB{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *AABB) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *AABB) Type() gdc.VariantType {
  return gdc.VariantTypeAABB
}

func (me *AABB) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *AABB) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *AABB) Abs() AABB {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1576868580) // FIXME: should cache?

  var ret AABB
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetCenter() Vector3 {
  name := StringNameFromStr("get_center")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1776574132) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetVolume() float32 {
  name := StringNameFromStr("get_volume")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) HasVolume() bool {
  name := StringNameFromStr("has_volume")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) HasSurface() bool {
  name := StringNameFromStr("has_surface")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) HasPoint(point Vector3, ) bool {
  name := StringNameFromStr("has_point")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1749054343) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) IsEqualApprox(aabb AABB, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 299946684) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) Intersects(with AABB, ) bool {
  name := StringNameFromStr("intersects")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 299946684) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) Encloses(with AABB, ) bool {
  name := StringNameFromStr("encloses")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 299946684) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) IntersectsPlane(plane Plane, ) bool {
  name := StringNameFromStr("intersects_plane")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1150170233) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{plane.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) Intersection(with AABB, ) AABB {
  name := StringNameFromStr("intersection")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1271470306) // FIXME: should cache?

  var ret AABB
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) Merge(with AABB, ) AABB {
  name := StringNameFromStr("merge")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1271470306) // FIXME: should cache?

  var ret AABB
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) Expand(to_point Vector3, ) AABB {
  name := StringNameFromStr("expand")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 2851643018) // FIXME: should cache?

  var ret AABB
  args := []gdc.ConstTypePtr{to_point.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) Grow(by float32, ) AABB {
  name := StringNameFromStr("grow")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 239217291) // FIXME: should cache?

  var ret AABB
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&by), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetSupport(dir Vector3, ) Vector3 {
  name := StringNameFromStr("get_support")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 2923479887) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{dir.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetLongestAxis() Vector3 {
  name := StringNameFromStr("get_longest_axis")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1776574132) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetLongestAxisIndex() int {
  name := StringNameFromStr("get_longest_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetLongestAxisSize() float32 {
  name := StringNameFromStr("get_longest_axis_size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetShortestAxis() Vector3 {
  name := StringNameFromStr("get_shortest_axis")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1776574132) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetShortestAxisIndex() int {
  name := StringNameFromStr("get_shortest_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetShortestAxisSize() float32 {
  name := StringNameFromStr("get_shortest_axis_size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) GetEndpoint(idx int, ) Vector3 {
  name := StringNameFromStr("get_endpoint")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 1394941017) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) IntersectsSegment(from Vector3, to Vector3, ) Variant {
  name := StringNameFromStr("intersects_segment")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 2048133369) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *AABB) IntersectsRay(from Vector3, dir Vector3, ) Variant {
  name := StringNameFromStr("intersects_ray")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, name.AsCPtr(), 2048133369) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{from.AsCTypePtr(), dir.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *AABB) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) EqualAABB(right AABB) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) NotEqualAABB(right AABB) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) MultiplyTransform3D(right Transform3D) AABB {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret AABB
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *AABB) Position() Vector3 {
  name := StringNameFromStr("position")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) SetPosition(value Vector3) {
  name := StringNameFromStr("position")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *AABB) Size() Vector3 {
  name := StringNameFromStr("size")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) SetSize(value Vector3) {
  name := StringNameFromStr("size")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *AABB) End() Vector3 {
  name := StringNameFromStr("end")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *AABB) SetEnd(value Vector3) {
  name := StringNameFromStr("end")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
