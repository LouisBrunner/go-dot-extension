// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused imports
var _ = fmt.Sprintf("")

type ptrsForRect2List struct {
  ctrFn gdc.PtrConstructor
  ctrFromRect2Fn gdc.PtrConstructor
  ctrFromRect2IFn gdc.PtrConstructor
  ctrFromVector2Vector2Fn gdc.PtrConstructor
  ctrFromFloat32Float32Float32Float32Fn gdc.PtrConstructor
  methodGetCenterFn gdc.PtrBuiltInMethod
  methodGetAreaFn gdc.PtrBuiltInMethod
  methodHasAreaFn gdc.PtrBuiltInMethod
  methodHasPointFn gdc.PtrBuiltInMethod
  methodIsEqualApproxFn gdc.PtrBuiltInMethod
  methodIsFiniteFn gdc.PtrBuiltInMethod
  methodIntersectsFn gdc.PtrBuiltInMethod
  methodEnclosesFn gdc.PtrBuiltInMethod
  methodIntersectionFn gdc.PtrBuiltInMethod
  methodMergeFn gdc.PtrBuiltInMethod
  methodExpandFn gdc.PtrBuiltInMethod
  methodGrowFn gdc.PtrBuiltInMethod
  methodGrowSideFn gdc.PtrBuiltInMethod
  methodGrowIndividualFn gdc.PtrBuiltInMethod
  methodAbsFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorEqualRect2Fn gdc.PtrOperatorEvaluator
  operatorNotEqualRect2Fn gdc.PtrOperatorEvaluator
  operatorMultiplyTransform2DFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  memberpositionGetterFn gdc.PtrGetter
  memberpositionSetterFn gdc.PtrSetter
  membersizeGetterFn gdc.PtrGetter
  membersizeSetterFn gdc.PtrSetter
  memberendGetterFn gdc.PtrGetter
  memberendSetterFn gdc.PtrSetter
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForRect2 ptrsForRect2List

func initRect2Ptrs(iface gdc.Interface) {
  ptrsForRect2.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 0))
  ptrsForRect2.ctrFromRect2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 1))
  ptrsForRect2.ctrFromRect2IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 2))
  ptrsForRect2.ctrFromVector2Vector2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 3))
  ptrsForRect2.ctrFromFloat32Float32Float32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 4))
  {
    methodName := StringNameFromStr("get_center")
    defer methodName.Destroy()
    ptrsForRect2.methodGetCenterFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 2428350749))
  }
  {
    methodName := StringNameFromStr("get_area")
    defer methodName.Destroy()
    ptrsForRect2.methodGetAreaFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("has_area")
    defer methodName.Destroy()
    ptrsForRect2.methodHasAreaFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("has_point")
    defer methodName.Destroy()
    ptrsForRect2.methodHasPointFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 3190634762))
  }
  {
    methodName := StringNameFromStr("is_equal_approx")
    defer methodName.Destroy()
    ptrsForRect2.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 1908192260))
  }
  {
    methodName := StringNameFromStr("is_finite")
    defer methodName.Destroy()
    ptrsForRect2.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("intersects")
    defer methodName.Destroy()
    ptrsForRect2.methodIntersectsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 819294880))
  }
  {
    methodName := StringNameFromStr("encloses")
    defer methodName.Destroy()
    ptrsForRect2.methodEnclosesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 1908192260))
  }
  {
    methodName := StringNameFromStr("intersection")
    defer methodName.Destroy()
    ptrsForRect2.methodIntersectionFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 2282977743))
  }
  {
    methodName := StringNameFromStr("merge")
    defer methodName.Destroy()
    ptrsForRect2.methodMergeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 2282977743))
  }
  {
    methodName := StringNameFromStr("expand")
    defer methodName.Destroy()
    ptrsForRect2.methodExpandFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 293272265))
  }
  {
    methodName := StringNameFromStr("grow")
    defer methodName.Destroy()
    ptrsForRect2.methodGrowFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 39664498))
  }
  {
    methodName := StringNameFromStr("grow_side")
    defer methodName.Destroy()
    ptrsForRect2.methodGrowSideFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 4177736158))
  }
  {
    methodName := StringNameFromStr("grow_individual")
    defer methodName.Destroy()
    ptrsForRect2.methodGrowIndividualFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 3203390369))
  }
  {
    methodName := StringNameFromStr("abs")
    defer methodName.Destroy()
    ptrsForRect2.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2, methodName.AsCPtr(), 3107653634))
  }
  ptrsForRect2.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeRect2, gdc.VariantTypeNil))
  ptrsForRect2.operatorEqualRect2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeRect2, gdc.VariantTypeRect2))
  ptrsForRect2.operatorNotEqualRect2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeRect2, gdc.VariantTypeRect2))
  ptrsForRect2.operatorMultiplyTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeRect2, gdc.VariantTypeTransform2D))
  ptrsForRect2.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeRect2, gdc.VariantTypeDictionary))
  ptrsForRect2.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeRect2, gdc.VariantTypeArray))
  {
    memberName := StringNameFromStr("position")
    defer memberName.Destroy()
    ptrsForRect2.memberpositionGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeRect2, memberName.AsCPtr()))
    ptrsForRect2.memberpositionSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeRect2, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("size")
    defer memberName.Destroy()
    ptrsForRect2.membersizeGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeRect2, memberName.AsCPtr()))
    ptrsForRect2.membersizeSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeRect2, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("end")
    defer memberName.Destroy()
    ptrsForRect2.memberendGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeRect2, memberName.AsCPtr()))
    ptrsForRect2.memberendSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeRect2, memberName.AsCPtr()))
  }
  ptrsForRect2.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeRect2))
  ptrsForRect2.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeRect2))
}

type Rect2 struct {
  data   *[classSizeRect2]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newRect2() *Rect2 {
  me := &Rect2{
    data:   new([classSizeRect2]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewRect2() *Rect2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newRect2()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewRect2FromRect2(from Rect2, ) *Rect2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newRect2()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2.ctrFromRect2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewRect2FromRect2I(from Rect2i, ) *Rect2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newRect2()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2.ctrFromRect2IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewRect2FromVector2Vector2(position Vector2, size Vector2, ) *Rect2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newRect2()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2.ctrFromVector2Vector2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr(), }))
  return me
}

func NewRect2FromFloat32Float32Float32Float32(x float64, y float64, width float64, height float64, ) *Rect2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&x)
  pinner.Pin(&y)
  pinner.Pin(&width)
  pinner.Pin(&height)
  me := newRect2()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2.ctrFromFloat32Float32Float32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), }))
  return me
}

// Destructor
func (me *Rect2) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsRect2() (*Rect2, error) {
	if me.Type() != gdc.VariantTypeRect2 {
		return nil, fmt.Errorf("variant is not a Rect2")
	}
  bclass := newRect2()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForRect2.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Rect2) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForRect2.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Rect2FromPtr(ptr gdc.ConstTypePtr) *Rect2 {
  me := newRect2()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Rect2) Type() gdc.VariantType {
  return gdc.VariantTypeRect2
}

func (me *Rect2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Rect2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Rect2) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Rect2) GetCenter() Vector2 {
  ret := NewVector2()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodGetCenterFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Rect2) GetArea() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodGetAreaFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Rect2) HasArea() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodHasAreaFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Rect2) HasPoint(point Vector2, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodHasPointFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Rect2) IsEqualApprox(rect Rect2, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{rect.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Rect2) IsFinite() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Rect2) Intersects(b Rect2, include_borders bool, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  varg1 := NewBoolFromBool(include_borders)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodIntersectsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Rect2) Encloses(b Rect2, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodEnclosesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Rect2) Intersection(b Rect2, ) Rect2 {
  ret := NewRect2()


  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodIntersectionFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Rect2) Merge(b Rect2, ) Rect2 {
  ret := NewRect2()


  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodMergeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Rect2) Expand(to Vector2, ) Rect2 {
  ret := NewRect2()


  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodExpandFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Rect2) Grow(amount float64, ) Rect2 {
  ret := NewRect2()

  varg0 := NewFloatFromFloat32(amount)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodGrowFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Rect2) GrowSide(side int64, amount float64, ) Rect2 {
  ret := NewRect2()

  varg0 := NewIntFromInt(side)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(amount)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodGrowSideFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Rect2) GrowIndividual(left float64, top float64, right float64, bottom float64, ) Rect2 {
  ret := NewRect2()

  varg0 := NewFloatFromFloat32(left)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(top)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(right)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(bottom)
  defer varg3.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodGrowIndividualFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Rect2) Abs() Rect2 {
  ret := NewRect2()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

// Operators

func (me *Rect2) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Rect2) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Rect2) Not() bool {
  opPtr := ptrsForRect2.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Rect2) EqualRect2(right Rect2) bool {
  opPtr := ptrsForRect2.operatorEqualRect2Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Rect2) NotEqualRect2(right Rect2) bool {
  opPtr := ptrsForRect2.operatorNotEqualRect2Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Rect2) MultiplyTransform2D(right Transform2D) Rect2 {
  opPtr := ptrsForRect2.operatorMultiplyTransform2DFn
  ret := NewRect2()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Rect2) InDictionary(right Dictionary) bool {
  opPtr := ptrsForRect2.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Rect2) InArray(right Array) bool {
  opPtr := ptrsForRect2.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Rect2) Position() Vector2 {
  ret := NewVector2()
  me.iface.CallPtrGetter(ensurePtr(ptrsForRect2.memberpositionGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Rect2) SetPosition(value Vector2) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForRect2.memberpositionSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Rect2) Size() Vector2 {
  ret := NewVector2()
  me.iface.CallPtrGetter(ensurePtr(ptrsForRect2.membersizeGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Rect2) SetSize(value Vector2) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForRect2.membersizeSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Rect2) End() Vector2 {
  ret := NewVector2()
  me.iface.CallPtrGetter(ensurePtr(ptrsForRect2.memberendGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Rect2) SetEnd(value Vector2) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForRect2.memberendSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
