// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector2 struct {
  data   *[classSizeVector2]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  Vector2Zero = "Vector2(0, 0)" // TODO: construct correctly
  Vector2One = "Vector2(1, 1)" // TODO: construct correctly
  Vector2Inf = "Vector2(inf, inf)" // TODO: construct correctly
  Vector2Left = "Vector2(-1, 0)" // TODO: construct correctly
  Vector2Right = "Vector2(1, 0)" // TODO: construct correctly
  Vector2Up = "Vector2(0, -1)" // TODO: construct correctly
  Vector2Down = "Vector2(0, 1)" // TODO: construct correctly
)

// Enums

type Vector2Axis int
const (
  Vector2AxisAxisX Vector2Axis = 0
  Vector2AxisAxisY Vector2Axis = 1
)

// Constructors
func newVector2() *Vector2 {
  me := &Vector2{
    data:   new([classSizeVector2]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewVector2() *Vector2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector2()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewVector2FromVector2(from Vector2, ) *Vector2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector2()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector2FromVector2I(from Vector2i, ) *Vector2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector2()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector2FromFloat32Float32(x float64, y float64, ) *Vector2 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&x)
  pinner.Pin(&y)
  me := newVector2()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 3) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), }))
  return me
}

// Destructor
func (me *Vector2) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector2() (*Vector2, error) {
	if me.Type() != gdc.VariantTypeVector2 {
		return nil, fmt.Errorf("variant is not a Vector2")
	}
  bclass := newVector2()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector2) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Vector2FromPtr(ptr gdc.ConstTypePtr) *Vector2 {
  me := newVector2()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Vector2) Type() gdc.VariantType {
  return gdc.VariantTypeVector2
}

func (me *Vector2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector2) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector2) Angle() float32 {
  name := StringNameFromStr("angle")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) AngleTo(to Vector2, ) float32 {
  name := StringNameFromStr("angle_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3819070308) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) AngleToPoint(to Vector2, ) float32 {
  name := StringNameFromStr("angle_to_point")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3819070308) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) DirectionTo(to Vector2, ) Vector2 {
  name := StringNameFromStr("direction_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2026743667) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) DistanceTo(to Vector2, ) float32 {
  name := StringNameFromStr("distance_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3819070308) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) DistanceSquaredTo(to Vector2, ) float32 {
  name := StringNameFromStr("distance_squared_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3819070308) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Length() float32 {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) LengthSquared() float32 {
  name := StringNameFromStr("length_squared")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) LimitLength(length float32, ) Vector2 {
  name := StringNameFromStr("limit_length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2544004089) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&length)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Normalized() Vector2 {
  name := StringNameFromStr("normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) IsNormalized() bool {
  name := StringNameFromStr("is_normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) IsEqualApprox(to Vector2, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3190634762) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) IsZeroApprox() bool {
  name := StringNameFromStr("is_zero_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Posmod(mod float32, ) Vector2 {
  name := StringNameFromStr("posmod")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2544004089) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&mod)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Posmodv(modv Vector2, ) Vector2 {
  name := StringNameFromStr("posmodv")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2026743667) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{modv.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Project(b Vector2, ) Vector2 {
  name := StringNameFromStr("project")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2026743667) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Lerp(to Vector2, weight float32, ) Vector2 {
  name := StringNameFromStr("lerp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 4250033116) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&weight)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Slerp(to Vector2, weight float32, ) Vector2 {
  name := StringNameFromStr("slerp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 4250033116) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&weight)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) CubicInterpolate(b Vector2, pre_a Vector2, post_b Vector2, weight float32, ) Vector2 {
  name := StringNameFromStr("cubic_interpolate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 193522989) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&weight)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) CubicInterpolateInTime(b Vector2, pre_a Vector2, post_b Vector2, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) Vector2 {
  name := StringNameFromStr("cubic_interpolate_in_time")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 1957055074) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&weight)), gdc.ConstTypePtr(unsafe.Pointer(&b_t)), gdc.ConstTypePtr(unsafe.Pointer(&pre_a_t)), gdc.ConstTypePtr(unsafe.Pointer(&post_b_t)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) BezierInterpolate(control_1 Vector2, control_2 Vector2, end Vector2, t float32, ) Vector2 {
  name := StringNameFromStr("bezier_interpolate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 193522989) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{control_1.AsCTypePtr(), control_2.AsCTypePtr(), end.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&t)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) BezierDerivative(control_1 Vector2, control_2 Vector2, end Vector2, t float32, ) Vector2 {
  name := StringNameFromStr("bezier_derivative")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 193522989) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{control_1.AsCTypePtr(), control_2.AsCTypePtr(), end.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&t)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) MaxAxisIndex() int {
  name := StringNameFromStr("max_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) MinAxisIndex() int {
  name := StringNameFromStr("min_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) MoveToward(to Vector2, delta float32, ) Vector2 {
  name := StringNameFromStr("move_toward")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 4250033116) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&delta)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Rotated(angle float32, ) Vector2 {
  name := StringNameFromStr("rotated")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2544004089) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&angle)), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Orthogonal() Vector2 {
  name := StringNameFromStr("orthogonal")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Floor() Vector2 {
  name := StringNameFromStr("floor")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Ceil() Vector2 {
  name := StringNameFromStr("ceil")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Round() Vector2 {
  name := StringNameFromStr("round")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Aspect() float32 {
  name := StringNameFromStr("aspect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Dot(with Vector2, ) float32 {
  name := StringNameFromStr("dot")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3819070308) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Slide(n Vector2, ) Vector2 {
  name := StringNameFromStr("slide")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2026743667) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{n.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Bounce(n Vector2, ) Vector2 {
  name := StringNameFromStr("bounce")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2026743667) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{n.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Reflect(n Vector2, ) Vector2 {
  name := StringNameFromStr("reflect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2026743667) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{n.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Cross(with Vector2, ) float32 {
  name := StringNameFromStr("cross")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 3819070308) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Abs() Vector2 {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Sign() Vector2 {
  name := StringNameFromStr("sign")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2428350749) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Clamp(min Vector2, max Vector2, ) Vector2 {
  name := StringNameFromStr("clamp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 318031021) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector2) Snapped(step Vector2, ) Vector2 {
  name := StringNameFromStr("snapped")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 2026743667) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{step.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func Vector2FromAngle(angle float32, ) Vector2 {
  name := StringNameFromStr("from_angle")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, name.AsCPtr(), 889263119) // FIXME: should cache?

  var ret Vector2
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&angle)), }

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Vector2) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) Negate() Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) Positive() Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) MultiplyInt(rightArg int64) Vector2 {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) DivideInt(rightArg int64) Vector2 {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) MultiplyFloat32(rightArg float64) Vector2 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) DivideFloat32(rightArg float64) Vector2 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) EqualVector2(right Vector2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) NotEqualVector2(right Vector2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) LessVector2(right Vector2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) LessEqualVector2(right Vector2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) GreaterVector2(right Vector2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) GreaterEqualVector2(right Vector2) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) AddVector2(right Vector2) Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) SubtractVector2(right Vector2) Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) MultiplyVector2(right Vector2) Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) DivideVector2(right Vector2) Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) MultiplyTransform2D(right Transform2D) Vector2 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector2
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) InPackedVector2Array(right PackedVector2Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Vector2) X() float32 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) SetX(value float32) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&value)))
}

func (me *Vector2) Y() float32 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector2) SetY(value float32) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&value)))
}
