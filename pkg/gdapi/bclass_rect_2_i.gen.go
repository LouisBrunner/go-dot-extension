// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Rect2i struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewRect2i() Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromRect2I(from Rect2i, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromRect2(from Rect2, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromVector2IVector2I(position Vector2i, size Vector2i, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr(), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromIntIntIntInt(x int, y int, width int, height int, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Rect2i) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Rect2i) Type() gdc.VariantType {
  return gdc.VariantTypeRect2I
}

func (me *Rect2i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Rect2i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Rect2i) GetCenter() Vector2i {
  name := StringNameFromStr("get_center")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 3444277866) // FIXME: should cache?

  var ret Vector2i
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) GetArea() int {
  name := StringNameFromStr("get_area")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) HasArea() bool {
  name := StringNameFromStr("has_area")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) HasPoint(point Vector2i, ) bool {
  name := StringNameFromStr("has_point")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 328189994) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) Intersects(b Rect2i, ) bool {
  name := StringNameFromStr("intersects")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 3434691493) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) Encloses(b Rect2i, ) bool {
  name := StringNameFromStr("encloses")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 3434691493) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) Intersection(b Rect2i, ) Rect2i {
  name := StringNameFromStr("intersection")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 717431873) // FIXME: should cache?

  var ret Rect2i
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) Merge(b Rect2i, ) Rect2i {
  name := StringNameFromStr("merge")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 717431873) // FIXME: should cache?

  var ret Rect2i
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) Expand(to Vector2i, ) Rect2i {
  name := StringNameFromStr("expand")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 1355196872) // FIXME: should cache?

  var ret Rect2i
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) Grow(amount int, ) Rect2i {
  name := StringNameFromStr("grow")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 1578070074) // FIXME: should cache?

  var ret Rect2i
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) GrowSide(side int, amount int, ) Rect2i {
  name := StringNameFromStr("grow_side")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 3191154199) // FIXME: should cache?

  var ret Rect2i
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&amount), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) GrowIndividual(left int, top int, right int, bottom int, ) Rect2i {
  name := StringNameFromStr("grow_individual")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 1893743416) // FIXME: should cache?

  var ret Rect2i
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&left), gdc.ConstTypePtr(&top), gdc.ConstTypePtr(&right), gdc.ConstTypePtr(&bottom), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Rect2i) Abs() Rect2i {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, name.AsCPtr(), 1469025700) // FIXME: should cache?

  var ret Rect2i
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Rect2i) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) EqualRect2I(right Rect2i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) NotEqualRect2I(right Rect2i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Rect2i) Position() Vector2i {
  name := StringNameFromStr("position")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2i
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) SetPosition(value Vector2i) {
  name := StringNameFromStr("position")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Rect2i) Size() Vector2i {
  name := StringNameFromStr("size")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2i
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) SetSize(value Vector2i) {
  name := StringNameFromStr("size")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Rect2i) End() Vector2i {
  name := StringNameFromStr("end")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2i
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Rect2i) SetEnd(value Vector2i) {
  name := StringNameFromStr("end")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
