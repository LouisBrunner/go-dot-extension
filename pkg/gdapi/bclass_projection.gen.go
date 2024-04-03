// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Projection struct {
  data   *[classSizeProjection]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  ProjectionIdentity = "Projection(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)" // TODO: construct correctly
  ProjectionZero = "Projection(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)" // TODO: construct correctly
)

// Enums

type ProjectionPlanes int
const (
  ProjectionPlanesPlaneNear ProjectionPlanes = 0
  ProjectionPlanesPlaneFar ProjectionPlanes = 1
  ProjectionPlanesPlaneLeft ProjectionPlanes = 2
  ProjectionPlanesPlaneTop ProjectionPlanes = 3
  ProjectionPlanesPlaneRight ProjectionPlanes = 4
  ProjectionPlanesPlaneBottom ProjectionPlanes = 5
)

// Constructors
func newProjection() *Projection {
  me := &Projection{
    data:   new([classSizeProjection]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewProjection() *Projection {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newProjection()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewProjectionFromProjection(from Projection, ) *Projection {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newProjection()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewProjectionFromTransform3D(from Transform3D, ) *Projection {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newProjection()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewProjectionFromVector4Vector4Vector4Vector4(x_axis Vector4, y_axis Vector4, z_axis Vector4, w_axis Vector4, ) *Projection {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newProjection()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 3) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), w_axis.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Projection) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsProjection() (*Projection, error) {
	if me.Type() != gdc.VariantTypeProjection {
		return nil, fmt.Errorf("variant is not a Projection")
	}
  bclass := newProjection()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Projection) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func ProjectionFromPtr(ptr gdc.ConstTypePtr) *Projection {
  me := newProjection()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Projection) Type() gdc.VariantType {
  return gdc.VariantTypeProjection
}

func (me *Projection) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Projection) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Projection) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func ProjectionCreateDepthCorrection(flip_y bool, ) Projection {
  name := StringNameFromStr("create_depth_correction")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 1228516048) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewBoolFromBool(flip_y)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateLightAtlasRect(rect Rect2, ) Projection {
  name := StringNameFromStr("create_light_atlas_rect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2654950662) // FIXME: should cache?

  ret := NewProjection()


  args := []gdc.ConstTypePtr{rect.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreatePerspective(fovy float64, aspect float64, z_near float64, z_far float64, flip_fov bool, ) Projection {
  name := StringNameFromStr("create_perspective")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 390915442) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewFloatFromFloat32(fovy)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(aspect)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(z_near)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(z_far)
  defer varg3.Destroy()
  varg4 := NewBoolFromBool(flip_fov)
  defer varg4.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreatePerspectiveHmd(fovy float64, aspect float64, z_near float64, z_far float64, flip_fov bool, eye int64, intraocular_dist float64, convergence_dist float64, ) Projection {
  name := StringNameFromStr("create_perspective_hmd")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2857674800) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewFloatFromFloat32(fovy)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(aspect)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(z_near)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(z_far)
  defer varg3.Destroy()
  varg4 := NewBoolFromBool(flip_fov)
  defer varg4.Destroy()
  varg5 := NewIntFromInt(eye)
  defer varg5.Destroy()
  varg6 := NewFloatFromFloat32(intraocular_dist)
  defer varg6.Destroy()
  varg7 := NewFloatFromFloat32(convergence_dist)
  defer varg7.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr(), varg7.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateForHmd(eye int64, aspect float64, intraocular_dist float64, display_width float64, display_to_lens float64, oversample float64, z_near float64, z_far float64, ) Projection {
  name := StringNameFromStr("create_for_hmd")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4184144994) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewIntFromInt(eye)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(aspect)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(intraocular_dist)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(display_width)
  defer varg3.Destroy()
  varg4 := NewFloatFromFloat32(display_to_lens)
  defer varg4.Destroy()
  varg5 := NewFloatFromFloat32(oversample)
  defer varg5.Destroy()
  varg6 := NewFloatFromFloat32(z_near)
  defer varg6.Destroy()
  varg7 := NewFloatFromFloat32(z_far)
  defer varg7.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr(), varg7.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateOrthogonal(left float64, right float64, bottom float64, top float64, z_near float64, z_far float64, ) Projection {
  name := StringNameFromStr("create_orthogonal")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3707929169) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewFloatFromFloat32(left)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(right)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(bottom)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(top)
  defer varg3.Destroy()
  varg4 := NewFloatFromFloat32(z_near)
  defer varg4.Destroy()
  varg5 := NewFloatFromFloat32(z_far)
  defer varg5.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateOrthogonalAspect(size float64, aspect float64, z_near float64, z_far float64, flip_fov bool, ) Projection {
  name := StringNameFromStr("create_orthogonal_aspect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 390915442) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewFloatFromFloat32(size)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(aspect)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(z_near)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(z_far)
  defer varg3.Destroy()
  varg4 := NewBoolFromBool(flip_fov)
  defer varg4.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateFrustum(left float64, right float64, bottom float64, top float64, z_near float64, z_far float64, ) Projection {
  name := StringNameFromStr("create_frustum")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3707929169) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewFloatFromFloat32(left)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(right)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(bottom)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(top)
  defer varg3.Destroy()
  varg4 := NewFloatFromFloat32(z_near)
  defer varg4.Destroy()
  varg5 := NewFloatFromFloat32(z_far)
  defer varg5.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateFrustumAspect(size float64, aspect float64, offset Vector2, z_near float64, z_far float64, flip_fov bool, ) Projection {
  name := StringNameFromStr("create_frustum_aspect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 1535076251) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewFloatFromFloat32(size)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(aspect)
  defer varg1.Destroy()

  varg3 := NewFloatFromFloat32(z_near)
  defer varg3.Destroy()
  varg4 := NewFloatFromFloat32(z_far)
  defer varg4.Destroy()
  varg5 := NewBoolFromBool(flip_fov)
  defer varg5.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), offset.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateFitAabb(aabb AABB, ) Projection {
  name := StringNameFromStr("create_fit_aabb")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2264694907) // FIXME: should cache?

  ret := NewProjection()


  args := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) Determinant() float64 {
  name := StringNameFromStr("determinant")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) PerspectiveZnearAdjusted(new_znear float64, ) Projection {
  name := StringNameFromStr("perspective_znear_adjusted")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3584785443) // FIXME: should cache?

  ret := NewProjection()

  varg0 := NewFloatFromFloat32(new_znear)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) GetProjectionPlane(plane int64, ) Plane {
  name := StringNameFromStr("get_projection_plane")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 1551184160) // FIXME: should cache?

  ret := NewPlane()

  varg0 := NewIntFromInt(plane)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) FlippedY() Projection {
  name := StringNameFromStr("flipped_y")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4212530932) // FIXME: should cache?

  ret := NewProjection()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) JitterOffseted(offset Vector2, ) Projection {
  name := StringNameFromStr("jitter_offseted")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2448438599) // FIXME: should cache?

  ret := NewProjection()


  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionGetFovy(fovx float64, aspect float64, ) float64 {
  name := StringNameFromStr("get_fovy")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3514207532) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(fovx)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(aspect)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetZFar() float64 {
  name := StringNameFromStr("get_z_far")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetZNear() float64 {
  name := StringNameFromStr("get_z_near")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetAspect() float64 {
  name := StringNameFromStr("get_aspect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetFov() float64 {
  name := StringNameFromStr("get_fov")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) IsOrthogonal() bool {
  name := StringNameFromStr("is_orthogonal")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetViewportHalfExtents() Vector2 {
  name := StringNameFromStr("get_viewport_half_extents")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2428350749) // FIXME: should cache?

  ret := NewVector2()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) GetFarPlaneHalfExtents() Vector2 {
  name := StringNameFromStr("get_far_plane_half_extents")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2428350749) // FIXME: should cache?

  ret := NewVector2()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) Inverse() Projection {
  name := StringNameFromStr("inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4212530932) // FIXME: should cache?

  ret := NewProjection()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) GetPixelsPerMeter(for_pixel_width int64, ) int64 {
  name := StringNameFromStr("get_pixels_per_meter")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(for_pixel_width)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetLodMultiplier() float64 {
  name := StringNameFromStr("get_lod_multiplier")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *Projection) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) MultiplyVector4(right Vector4) Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) EqualProjection(right Projection) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) NotEqualProjection(right Projection) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) MultiplyProjection(right Projection) Projection {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewProjection()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Projection) X() Vector4 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewVector4()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetX(value Vector4) {
  name := StringNameFromStr("x")
  defer name.Destroy()
  valueV := value

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Projection) Y() Vector4 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewVector4()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetY(value Vector4) {
  name := StringNameFromStr("y")
  defer name.Destroy()
  valueV := value

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Projection) Z() Vector4 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewVector4()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetZ(value Vector4) {
  name := StringNameFromStr("z")
  defer name.Destroy()
  valueV := value

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Projection) W() Vector4 {
  name := StringNameFromStr("w")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewVector4()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetW(value Vector4) {
  name := StringNameFromStr("w")
  defer name.Destroy()
  valueV := value

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}
