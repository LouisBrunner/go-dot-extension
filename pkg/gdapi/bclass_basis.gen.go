// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Basis struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  BasisIdentity = "Basis(1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipX = "Basis(-1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipY = "Basis(1, 0, 0, 0, -1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipZ = "Basis(1, 0, 0, 0, 1, 0, 0, 0, -1)" // TODO: construct correctly
)

// Enums

// Constructors

func NewBasis() Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromBasis(from Basis, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromQuaternion(from Quaternion, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromVector3Float32(axis Vector3, angle float32, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromVector3Vector3Vector3(x_axis Vector3, y_axis Vector3, z_axis Vector3, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Basis) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Basis) Type() gdc.VariantType {
  return gdc.VariantTypeBasis
}

func (me *Basis) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Basis) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Basis) Inverse() Basis {
  name := StringNameFromStr("inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 594669093) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Transposed() Basis {
  name := StringNameFromStr("transposed")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 594669093) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Orthonormalized() Basis {
  name := StringNameFromStr("orthonormalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 594669093) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Determinant() float32 {
  name := StringNameFromStr("determinant")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Rotated(axis Vector3, angle float32, ) Basis {
  name := StringNameFromStr("rotated")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 1998708965) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Scaled(scale Vector3, ) Basis {
  name := StringNameFromStr("scaled")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 3934786792) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) GetScale() Vector3 {
  name := StringNameFromStr("get_scale")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 1776574132) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) GetEuler(order int, ) Vector3 {
  name := StringNameFromStr("get_euler")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 1394941017) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Tdotx(with Vector3, ) float32 {
  name := StringNameFromStr("tdotx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 1047977935) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Tdoty(with Vector3, ) float32 {
  name := StringNameFromStr("tdoty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 1047977935) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Tdotz(with Vector3, ) float32 {
  name := StringNameFromStr("tdotz")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 1047977935) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) Slerp(to Basis, weight float32, ) Basis {
  name := StringNameFromStr("slerp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 3118673011) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), gdc.ConstTypePtr(&weight), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) IsConformal() bool {
  name := StringNameFromStr("is_conformal")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) IsEqualApprox(b Basis, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 3165333982) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Basis) GetRotationQuaternion() Quaternion {
  name := StringNameFromStr("get_rotation_quaternion")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 4274879941) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func BasisLookingAt(target Vector3, up Vector3, use_model_front bool, ) Basis {
  name := StringNameFromStr("looking_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 3728732505) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{target.AsCTypePtr(), up.AsCTypePtr(), gdc.ConstTypePtr(&use_model_front), }

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func BasisFromScale(scale Vector3, ) Basis {
  name := StringNameFromStr("from_scale")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 3703240166) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func BasisFromEuler(euler Vector3, order int, ) Basis {
  name := StringNameFromStr("from_euler")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, name.AsCPtr(), 2802321791) // FIXME: should cache?

  var ret Basis
  args := []gdc.ConstTypePtr{euler.AsCTypePtr(), gdc.ConstTypePtr(&order), }

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Basis) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) MultiplyInt(rightArg int) Basis {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Basis
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) MultiplyFloat32(rightArg float32) Basis {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Basis
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) MultiplyVector3(right Vector3) Vector3 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) EqualBasis(right Basis) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) NotEqualBasis(right Basis) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) MultiplyBasis(right Basis) Basis {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Basis
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Basis) X() Vector3 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) SetX(value Vector3) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Basis) Y() Vector3 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) SetY(value Vector3) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Basis) Z() Vector3 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Basis) SetZ(value Vector3) {
  name := StringNameFromStr("z")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
