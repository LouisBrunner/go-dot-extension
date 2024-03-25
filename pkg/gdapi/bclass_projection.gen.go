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

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&flip_y)), }

  pinner.Pin(&flip_y)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreateLightAtlasRect(rect Rect2, ) Projection {
  name := StringNameFromStr("create_light_atlas_rect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2654950662) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{rect.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreatePerspective(fovy float32, aspect float32, z_near float32, z_far float32, flip_fov bool, ) Projection {
  name := StringNameFromStr("create_perspective")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 390915442) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&fovy)), gdc.ConstTypePtr(unsafe.Pointer(&aspect)), gdc.ConstTypePtr(unsafe.Pointer(&z_near)), gdc.ConstTypePtr(unsafe.Pointer(&z_far)), gdc.ConstTypePtr(unsafe.Pointer(&flip_fov)), }

  pinner.Pin(&fovy)
  pinner.Pin(&aspect)
  pinner.Pin(&z_near)
  pinner.Pin(&z_far)
  pinner.Pin(&flip_fov)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreatePerspectiveHmd(fovy float32, aspect float32, z_near float32, z_far float32, flip_fov bool, eye int, intraocular_dist float32, convergence_dist float32, ) Projection {
  name := StringNameFromStr("create_perspective_hmd")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2857674800) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&fovy)), gdc.ConstTypePtr(unsafe.Pointer(&aspect)), gdc.ConstTypePtr(unsafe.Pointer(&z_near)), gdc.ConstTypePtr(unsafe.Pointer(&z_far)), gdc.ConstTypePtr(unsafe.Pointer(&flip_fov)), gdc.ConstTypePtr(unsafe.Pointer(&eye)), gdc.ConstTypePtr(unsafe.Pointer(&intraocular_dist)), gdc.ConstTypePtr(unsafe.Pointer(&convergence_dist)), }

  pinner.Pin(&fovy)
  pinner.Pin(&aspect)
  pinner.Pin(&z_near)
  pinner.Pin(&z_far)
  pinner.Pin(&flip_fov)
  pinner.Pin(&eye)
  pinner.Pin(&intraocular_dist)
  pinner.Pin(&convergence_dist)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreateForHmd(eye int, aspect float32, intraocular_dist float32, display_width float32, display_to_lens float32, oversample float32, z_near float32, z_far float32, ) Projection {
  name := StringNameFromStr("create_for_hmd")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4184144994) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&eye)), gdc.ConstTypePtr(unsafe.Pointer(&aspect)), gdc.ConstTypePtr(unsafe.Pointer(&intraocular_dist)), gdc.ConstTypePtr(unsafe.Pointer(&display_width)), gdc.ConstTypePtr(unsafe.Pointer(&display_to_lens)), gdc.ConstTypePtr(unsafe.Pointer(&oversample)), gdc.ConstTypePtr(unsafe.Pointer(&z_near)), gdc.ConstTypePtr(unsafe.Pointer(&z_far)), }

  pinner.Pin(&eye)
  pinner.Pin(&aspect)
  pinner.Pin(&intraocular_dist)
  pinner.Pin(&display_width)
  pinner.Pin(&display_to_lens)
  pinner.Pin(&oversample)
  pinner.Pin(&z_near)
  pinner.Pin(&z_far)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreateOrthogonal(left float32, right float32, bottom float32, top float32, z_near float32, z_far float32, ) Projection {
  name := StringNameFromStr("create_orthogonal")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3707929169) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&left)), gdc.ConstTypePtr(unsafe.Pointer(&right)), gdc.ConstTypePtr(unsafe.Pointer(&bottom)), gdc.ConstTypePtr(unsafe.Pointer(&top)), gdc.ConstTypePtr(unsafe.Pointer(&z_near)), gdc.ConstTypePtr(unsafe.Pointer(&z_far)), }

  pinner.Pin(&left)
  pinner.Pin(&right)
  pinner.Pin(&bottom)
  pinner.Pin(&top)
  pinner.Pin(&z_near)
  pinner.Pin(&z_far)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreateOrthogonalAspect(size float32, aspect float32, z_near float32, z_far float32, flip_fov bool, ) Projection {
  name := StringNameFromStr("create_orthogonal_aspect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 390915442) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&size)), gdc.ConstTypePtr(unsafe.Pointer(&aspect)), gdc.ConstTypePtr(unsafe.Pointer(&z_near)), gdc.ConstTypePtr(unsafe.Pointer(&z_far)), gdc.ConstTypePtr(unsafe.Pointer(&flip_fov)), }

  pinner.Pin(&size)
  pinner.Pin(&aspect)
  pinner.Pin(&z_near)
  pinner.Pin(&z_far)
  pinner.Pin(&flip_fov)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreateFrustum(left float32, right float32, bottom float32, top float32, z_near float32, z_far float32, ) Projection {
  name := StringNameFromStr("create_frustum")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3707929169) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&left)), gdc.ConstTypePtr(unsafe.Pointer(&right)), gdc.ConstTypePtr(unsafe.Pointer(&bottom)), gdc.ConstTypePtr(unsafe.Pointer(&top)), gdc.ConstTypePtr(unsafe.Pointer(&z_near)), gdc.ConstTypePtr(unsafe.Pointer(&z_far)), }

  pinner.Pin(&left)
  pinner.Pin(&right)
  pinner.Pin(&bottom)
  pinner.Pin(&top)
  pinner.Pin(&z_near)
  pinner.Pin(&z_far)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreateFrustumAspect(size float32, aspect float32, offset Vector2, z_near float32, z_far float32, flip_fov bool, ) Projection {
  name := StringNameFromStr("create_frustum_aspect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 1535076251) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&size)), gdc.ConstTypePtr(unsafe.Pointer(&aspect)), offset.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&z_near)), gdc.ConstTypePtr(unsafe.Pointer(&z_far)), gdc.ConstTypePtr(unsafe.Pointer(&flip_fov)), }

  pinner.Pin(&size)
  pinner.Pin(&aspect)
  pinner.Pin(&z_near)
  pinner.Pin(&z_far)
  pinner.Pin(&flip_fov)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionCreateFitAabb(aabb AABB, ) Projection {
  name := StringNameFromStr("create_fit_aabb")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2264694907) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) Determinant() float32 {
  name := StringNameFromStr("determinant")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) PerspectiveZnearAdjusted(new_znear float32, ) Projection {
  name := StringNameFromStr("perspective_znear_adjusted")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3584785443) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&new_znear)), }

  pinner.Pin(&new_znear)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetProjectionPlane(plane int, ) Plane {
  name := StringNameFromStr("get_projection_plane")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 1551184160) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Plane
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&plane)), }

  pinner.Pin(&plane)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) FlippedY() Projection {
  name := StringNameFromStr("flipped_y")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4212530932) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) JitterOffseted(offset Vector2, ) Projection {
  name := StringNameFromStr("jitter_offseted")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2448438599) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func ProjectionGetFovy(fovx float32, aspect float32, ) float32 {
  name := StringNameFromStr("get_fovy")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3514207532) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&fovx)), gdc.ConstTypePtr(unsafe.Pointer(&aspect)), }

  pinner.Pin(&fovx)
  pinner.Pin(&aspect)

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetZFar() float32 {
  name := StringNameFromStr("get_z_far")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetZNear() float32 {
  name := StringNameFromStr("get_z_near")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetAspect() float32 {
  name := StringNameFromStr("get_aspect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetFov() float32 {
  name := StringNameFromStr("get_fov")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) IsOrthogonal() bool {
  name := StringNameFromStr("is_orthogonal")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetViewportHalfExtents() Vector2 {
  name := StringNameFromStr("get_viewport_half_extents")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2428350749) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector2
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetFarPlaneHalfExtents() Vector2 {
  name := StringNameFromStr("get_far_plane_half_extents")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 2428350749) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector2
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) Inverse() Projection {
  name := StringNameFromStr("inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4212530932) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Projection
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetPixelsPerMeter(for_pixel_width int, ) int {
  name := StringNameFromStr("get_pixels_per_meter")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 4103005248) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&for_pixel_width)), }

  pinner.Pin(&for_pixel_width)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Projection) GetLodMultiplier() float32 {
  name := StringNameFromStr("get_lod_multiplier")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, name.AsCPtr(), 466405837) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Projection) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) MultiplyVector4(right Vector4) Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) EqualProjection(right Projection) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) NotEqualProjection(right Projection) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) MultiplyProjection(right Projection) Projection {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Projection
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Projection) X() Vector4 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) SetX(value Vector4) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Projection) Y() Vector4 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) SetY(value Vector4) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Projection) Z() Vector4 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) SetZ(value Vector4) {
  name := StringNameFromStr("z")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Projection) W() Vector4 {
  name := StringNameFromStr("w")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Projection) SetW(value Vector4) {
  name := StringNameFromStr("w")
  defer name.Destroy()

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
