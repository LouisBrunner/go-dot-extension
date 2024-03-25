// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Transform3D struct {
  data   *[classSizeTransform3D]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  Transform3DIdentity = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipX = "Transform3D(-1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipY = "Transform3D(1, 0, 0, 0, -1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipZ = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, -1, 0, 0, 0)" // TODO: construct correctly
)

// Enums

// Constructors
func newTransform3D() *Transform3D {
  me := &Transform3D{
    data:   new([classSizeTransform3D]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewTransform3D() *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewTransform3DFromTransform3D(from Transform3D, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewTransform3DFromBasisVector3(basis Basis, origin Vector3, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{basis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return me
}

func NewTransform3DFromVector3Vector3Vector3Vector3(x_axis Vector3, y_axis Vector3, z_axis Vector3, origin Vector3, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 3) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return me
}

func NewTransform3DFromProjection(from Projection, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 4) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Transform3D) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsTransform3D() (*Transform3D, error) {
	if me.Type() != gdc.VariantTypeTransform3D {
		return nil, fmt.Errorf("variant is not a Transform3D")
	}
  bclass := newTransform3D()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Transform3D) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Transform3DFromPtr(ptr gdc.ConstTypePtr) *Transform3D {
  me := newTransform3D()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Transform3D) Type() gdc.VariantType {
  return gdc.VariantTypeTransform3D
}

func (me *Transform3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Transform3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Transform3D) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Transform3D) Inverse() Transform3D {
  name := StringNameFromStr("inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 3816817146) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) AffineInverse() Transform3D {
  name := StringNameFromStr("affine_inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 3816817146) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) Orthonormalized() Transform3D {
  name := StringNameFromStr("orthonormalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 3816817146) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) Rotated(axis Vector3, angle float32, ) Transform3D {
  name := StringNameFromStr("rotated")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 1563203923) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&angle)), }

  pinner.Pin(&angle)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) RotatedLocal(axis Vector3, angle float32, ) Transform3D {
  name := StringNameFromStr("rotated_local")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 1563203923) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&angle)), }

  pinner.Pin(&angle)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) Scaled(scale Vector3, ) Transform3D {
  name := StringNameFromStr("scaled")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 1405596198) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) ScaledLocal(scale Vector3, ) Transform3D {
  name := StringNameFromStr("scaled_local")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 1405596198) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) Translated(offset Vector3, ) Transform3D {
  name := StringNameFromStr("translated")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 1405596198) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) TranslatedLocal(offset Vector3, ) Transform3D {
  name := StringNameFromStr("translated_local")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 1405596198) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) LookingAt(target Vector3, up Vector3, use_model_front bool, ) Transform3D {
  name := StringNameFromStr("looking_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 90889270) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{target.AsCTypePtr(), up.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&use_model_front)), }

  pinner.Pin(&use_model_front)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) InterpolateWith(xform Transform3D, weight float32, ) Transform3D {
  name := StringNameFromStr("interpolate_with")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 1786453358) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform3D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{xform.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&weight)), }

  pinner.Pin(&weight)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) IsEqualApprox(xform Transform3D, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 696001652) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{xform.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform3D) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Transform3D) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) MultiplyInt(rightArg int64) Transform3D {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Transform3D
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) MultiplyFloat32(rightArg float64) Transform3D {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Transform3D
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) MultiplyVector3(right Vector3) Vector3 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) MultiplyPlane(right Plane) Plane {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Plane
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) MultiplyAABB(right AABB) AABB {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret AABB
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) EqualTransform3D(right Transform3D) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) NotEqualTransform3D(right Transform3D) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) MultiplyTransform3D(right Transform3D) Transform3D {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Transform3D
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) MultiplyPackedVector3Array(right PackedVector3Array) PackedVector3Array {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret PackedVector3Array
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Transform3D) Basis() Basis {
  name := StringNameFromStr("basis")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Basis
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) SetBasis(value Basis) {
  name := StringNameFromStr("basis")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Transform3D) Origin() Vector3 {
  name := StringNameFromStr("origin")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform3D) SetOrigin(value Vector3) {
  name := StringNameFromStr("origin")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
