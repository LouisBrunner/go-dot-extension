// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Transform2D struct {
  data   *[classSizeTransform2D]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  Transform2DIdentity = "Transform2D(1, 0, 0, 1, 0, 0)" // TODO: construct correctly
  Transform2DFlipX = "Transform2D(-1, 0, 0, 1, 0, 0)" // TODO: construct correctly
  Transform2DFlipY = "Transform2D(1, 0, 0, -1, 0, 0)" // TODO: construct correctly
)

// Enums

// Constructors
func newTransform2D() *Transform2D {
  me := &Transform2D{
    data:   new([classSizeTransform2D]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewTransform2D() *Transform2D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform2D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewTransform2DFromTransform2D(from Transform2D, ) *Transform2D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform2D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewTransform2DFromFloat32Vector2(rotation float64, position Vector2, ) *Transform2D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&rotation)
  me := newTransform2D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation), position.AsCTypePtr(), }))
  return me
}

func NewTransform2DFromFloat32Vector2Float32Vector2(rotation float64, scale Vector2, skew float64, position Vector2, ) *Transform2D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&rotation)
  pinner.Pin(&skew)
  me := newTransform2D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 3) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation), scale.AsCTypePtr(), gdc.ConstTypePtr(&skew), position.AsCTypePtr(), }))
  return me
}

func NewTransform2DFromVector2Vector2Vector2(x_axis Vector2, y_axis Vector2, origin Vector2, ) *Transform2D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform2D()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 4) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Transform2D) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsTransform2D() (*Transform2D, error) {
	if me.Type() != gdc.VariantTypeTransform2D {
		return nil, fmt.Errorf("variant is not a Transform2D")
	}
  bclass := newTransform2D()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Transform2D) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Transform2DFromPtr(ptr gdc.ConstTypePtr) *Transform2D {
  me := newTransform2D()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Transform2D) Type() gdc.VariantType {
  return gdc.VariantTypeTransform2D
}

func (me *Transform2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Transform2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Transform2D) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Transform2D) Inverse() Transform2D {
  name := StringNameFromStr("inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1420440541) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) AffineInverse() Transform2D {
  name := StringNameFromStr("affine_inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1420440541) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) GetRotation() float32 {
  name := StringNameFromStr("get_rotation")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) GetOrigin() Vector2 {
  name := StringNameFromStr("get_origin")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 2428350749) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector2
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) GetScale() Vector2 {
  name := StringNameFromStr("get_scale")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 2428350749) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector2
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) GetSkew() float32 {
  name := StringNameFromStr("get_skew")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) Orthonormalized() Transform2D {
  name := StringNameFromStr("orthonormalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1420440541) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) Rotated(angle float32, ) Transform2D {
  name := StringNameFromStr("rotated")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 729597514) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&angle)), }

  pinner.Pin(&angle)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) RotatedLocal(angle float32, ) Transform2D {
  name := StringNameFromStr("rotated_local")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 729597514) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&angle)), }

  pinner.Pin(&angle)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) Scaled(scale Vector2, ) Transform2D {
  name := StringNameFromStr("scaled")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1446323263) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) ScaledLocal(scale Vector2, ) Transform2D {
  name := StringNameFromStr("scaled_local")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1446323263) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) Translated(offset Vector2, ) Transform2D {
  name := StringNameFromStr("translated")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1446323263) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) TranslatedLocal(offset Vector2, ) Transform2D {
  name := StringNameFromStr("translated_local")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1446323263) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) Determinant() float32 {
  name := StringNameFromStr("determinant")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) BasisXform(v Vector2, ) Vector2 {
  name := StringNameFromStr("basis_xform")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 2026743667) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector2
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{v.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) BasisXformInv(v Vector2, ) Vector2 {
  name := StringNameFromStr("basis_xform_inv")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 2026743667) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector2
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{v.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) InterpolateWith(xform Transform2D, weight float32, ) Transform2D {
  name := StringNameFromStr("interpolate_with")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 359399686) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{xform.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&weight)), }

  pinner.Pin(&weight)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) IsConformal() bool {
  name := StringNameFromStr("is_conformal")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) IsEqualApprox(xform Transform2D, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 3837431929) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{xform.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Transform2D) LookingAt(target Vector2, ) Transform2D {
  name := StringNameFromStr("looking_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, name.AsCPtr(), 1446323263) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Transform2D
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{target.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Transform2D) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) MultiplyInt(rightArg int64) Transform2D {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Transform2D
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) MultiplyFloat32(rightArg float64) Transform2D {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Transform2D
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) MultiplyVector2(right Vector2) Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) MultiplyRect2(right Rect2) Rect2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Rect2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) EqualTransform2D(right Transform2D) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) NotEqualTransform2D(right Transform2D) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) MultiplyTransform2D(right Transform2D) Transform2D {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Transform2D
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) MultiplyPackedVector2Array(right PackedVector2Array) PackedVector2Array {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret PackedVector2Array
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Transform2D) X() Vector2 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) SetX(value Vector2) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Transform2D) Y() Vector2 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) SetY(value Vector2) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Transform2D) Origin() Vector2 {
  name := StringNameFromStr("origin")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Transform2D) SetOrigin(value Vector2) {
  name := StringNameFromStr("origin")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
