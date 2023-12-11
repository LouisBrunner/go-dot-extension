// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Rect2 struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewRect2() Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromRect2(from Rect2, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromRect2I(from Rect2i, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromVector2Vector2(position Vector2, size Vector2, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr(), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromFloat32Float32Float32Float32(x float32, y float32, width float32, height float32, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Rect2) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Rect2) Type() gdc.VariantType {
  return gdc.VariantTypeRect2
}

func (me *Rect2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Rect2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Rect2) GetCenter() Vector2 {
  name := StringNameFromStr("get_center")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) GetArea() float32 {
  name := StringNameFromStr("get_area")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) HasArea() bool {
  name := StringNameFromStr("has_area")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) HasPoint(point Vector2, ) bool {
  name := StringNameFromStr("has_point")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 3190634762) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) IsEqualApprox(rect Rect2, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 1908192260) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{rect.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) Intersects(b Rect2, include_borders bool, ) bool {
  name := StringNameFromStr("intersects")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 819294880) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), gdc.ConstTypePtr(&include_borders), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) Encloses(b Rect2, ) bool {
  name := StringNameFromStr("encloses")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 1908192260) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) Intersection(b Rect2, ) Rect2 {
  name := StringNameFromStr("intersection")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 2282977743) // FIXME: should cache?

  var ret Rect2
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) Merge(b Rect2, ) Rect2 {
  name := StringNameFromStr("merge")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 2282977743) // FIXME: should cache?

  var ret Rect2
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) Expand(to Vector2, ) Rect2 {
  name := StringNameFromStr("expand")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 293272265) // FIXME: should cache?

  var ret Rect2
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) Grow(amount float32, ) Rect2 {
  name := StringNameFromStr("grow")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 39664498) // FIXME: should cache?

  var ret Rect2
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) GrowSide(side int, amount float32, ) Rect2 {
  name := StringNameFromStr("grow_side")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 4177736158) // FIXME: should cache?

  var ret Rect2
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&amount), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) GrowIndividual(left float32, top float32, right float32, bottom float32, ) Rect2 {
  name := StringNameFromStr("grow_individual")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 3203390369) // FIXME: should cache?

  var ret Rect2
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&left), gdc.ConstTypePtr(&top), gdc.ConstTypePtr(&right), gdc.ConstTypePtr(&bottom), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2) Abs() Rect2 {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, name.AsCPtr(), 3107653634) // FIXME: should cache?

  var ret Rect2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Rect2) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) EqualRect2(right Rect2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) NotEqualRect2(right Rect2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) MultiplyTransform2D(right Transform2D) Rect2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Rect2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Rect2) Position() Vector2 {
  name := StringNameFromStr("position")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) SetPosition(value Vector2) {
  name := StringNameFromStr("position")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Rect2) Size() Vector2 {
  name := StringNameFromStr("size")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) SetSize(value Vector2) {
  name := StringNameFromStr("size")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Rect2) End() Vector2 {
  name := StringNameFromStr("end")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2) SetEnd(value Vector2) {
  name := StringNameFromStr("end")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
