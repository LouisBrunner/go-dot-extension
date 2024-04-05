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

type ptrsForProjectionList struct {
  ctrFn gdc.PtrConstructor
  ctrFromProjectionFn gdc.PtrConstructor
  ctrFromTransform3DFn gdc.PtrConstructor
  ctrFromVector4Vector4Vector4Vector4Fn gdc.PtrConstructor
  methodCreateDepthCorrectionFn gdc.PtrBuiltInMethod
  methodCreateLightAtlasRectFn gdc.PtrBuiltInMethod
  methodCreatePerspectiveFn gdc.PtrBuiltInMethod
  methodCreatePerspectiveHmdFn gdc.PtrBuiltInMethod
  methodCreateForHmdFn gdc.PtrBuiltInMethod
  methodCreateOrthogonalFn gdc.PtrBuiltInMethod
  methodCreateOrthogonalAspectFn gdc.PtrBuiltInMethod
  methodCreateFrustumFn gdc.PtrBuiltInMethod
  methodCreateFrustumAspectFn gdc.PtrBuiltInMethod
  methodCreateFitAabbFn gdc.PtrBuiltInMethod
  methodDeterminantFn gdc.PtrBuiltInMethod
  methodPerspectiveZnearAdjustedFn gdc.PtrBuiltInMethod
  methodGetProjectionPlaneFn gdc.PtrBuiltInMethod
  methodFlippedYFn gdc.PtrBuiltInMethod
  methodJitterOffsetedFn gdc.PtrBuiltInMethod
  methodGetFovyFn gdc.PtrBuiltInMethod
  methodGetZFarFn gdc.PtrBuiltInMethod
  methodGetZNearFn gdc.PtrBuiltInMethod
  methodGetAspectFn gdc.PtrBuiltInMethod
  methodGetFovFn gdc.PtrBuiltInMethod
  methodIsOrthogonalFn gdc.PtrBuiltInMethod
  methodGetViewportHalfExtentsFn gdc.PtrBuiltInMethod
  methodGetFarPlaneHalfExtentsFn gdc.PtrBuiltInMethod
  methodInverseFn gdc.PtrBuiltInMethod
  methodGetPixelsPerMeterFn gdc.PtrBuiltInMethod
  methodGetLodMultiplierFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorMultiplyVector4Fn gdc.PtrOperatorEvaluator
  operatorEqualProjectionFn gdc.PtrOperatorEvaluator
  operatorNotEqualProjectionFn gdc.PtrOperatorEvaluator
  operatorMultiplyProjectionFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  memberxGetterFn gdc.PtrGetter
  memberxSetterFn gdc.PtrSetter
  memberyGetterFn gdc.PtrGetter
  memberySetterFn gdc.PtrSetter
  memberzGetterFn gdc.PtrGetter
  memberzSetterFn gdc.PtrSetter
  memberwGetterFn gdc.PtrGetter
  memberwSetterFn gdc.PtrSetter
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForProjection ptrsForProjectionList

func initProjectionPtrs(iface gdc.Interface) {
  ptrsForProjection.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 0))
  ptrsForProjection.ctrFromProjectionFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 1))
  ptrsForProjection.ctrFromTransform3DFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 2))
  ptrsForProjection.ctrFromVector4Vector4Vector4Vector4Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 3))
  {
    methodName := StringNameFromStr("create_depth_correction")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateDepthCorrectionFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 1228516048))
  }
  {
    methodName := StringNameFromStr("create_light_atlas_rect")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateLightAtlasRectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 2654950662))
  }
  {
    methodName := StringNameFromStr("create_perspective")
    defer methodName.Destroy()
    ptrsForProjection.methodCreatePerspectiveFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 390915442))
  }
  {
    methodName := StringNameFromStr("create_perspective_hmd")
    defer methodName.Destroy()
    ptrsForProjection.methodCreatePerspectiveHmdFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 2857674800))
  }
  {
    methodName := StringNameFromStr("create_for_hmd")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateForHmdFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 4184144994))
  }
  {
    methodName := StringNameFromStr("create_orthogonal")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateOrthogonalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 3707929169))
  }
  {
    methodName := StringNameFromStr("create_orthogonal_aspect")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateOrthogonalAspectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 390915442))
  }
  {
    methodName := StringNameFromStr("create_frustum")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateFrustumFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 3707929169))
  }
  {
    methodName := StringNameFromStr("create_frustum_aspect")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateFrustumAspectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 1535076251))
  }
  {
    methodName := StringNameFromStr("create_fit_aabb")
    defer methodName.Destroy()
    ptrsForProjection.methodCreateFitAabbFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 2264694907))
  }
  {
    methodName := StringNameFromStr("determinant")
    defer methodName.Destroy()
    ptrsForProjection.methodDeterminantFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("perspective_znear_adjusted")
    defer methodName.Destroy()
    ptrsForProjection.methodPerspectiveZnearAdjustedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 3584785443))
  }
  {
    methodName := StringNameFromStr("get_projection_plane")
    defer methodName.Destroy()
    ptrsForProjection.methodGetProjectionPlaneFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 1551184160))
  }
  {
    methodName := StringNameFromStr("flipped_y")
    defer methodName.Destroy()
    ptrsForProjection.methodFlippedYFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 4212530932))
  }
  {
    methodName := StringNameFromStr("jitter_offseted")
    defer methodName.Destroy()
    ptrsForProjection.methodJitterOffsetedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 2448438599))
  }
  {
    methodName := StringNameFromStr("get_fovy")
    defer methodName.Destroy()
    ptrsForProjection.methodGetFovyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 3514207532))
  }
  {
    methodName := StringNameFromStr("get_z_far")
    defer methodName.Destroy()
    ptrsForProjection.methodGetZFarFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("get_z_near")
    defer methodName.Destroy()
    ptrsForProjection.methodGetZNearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("get_aspect")
    defer methodName.Destroy()
    ptrsForProjection.methodGetAspectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("get_fov")
    defer methodName.Destroy()
    ptrsForProjection.methodGetFovFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("is_orthogonal")
    defer methodName.Destroy()
    ptrsForProjection.methodIsOrthogonalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("get_viewport_half_extents")
    defer methodName.Destroy()
    ptrsForProjection.methodGetViewportHalfExtentsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 2428350749))
  }
  {
    methodName := StringNameFromStr("get_far_plane_half_extents")
    defer methodName.Destroy()
    ptrsForProjection.methodGetFarPlaneHalfExtentsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 2428350749))
  }
  {
    methodName := StringNameFromStr("inverse")
    defer methodName.Destroy()
    ptrsForProjection.methodInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 4212530932))
  }
  {
    methodName := StringNameFromStr("get_pixels_per_meter")
    defer methodName.Destroy()
    ptrsForProjection.methodGetPixelsPerMeterFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("get_lod_multiplier")
    defer methodName.Destroy()
    ptrsForProjection.methodGetLodMultiplierFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeProjection, methodName.AsCPtr(), 466405837))
  }
  ptrsForProjection.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeProjection, gdc.VariantTypeNil))
  ptrsForProjection.operatorMultiplyVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeProjection, gdc.VariantTypeVector4))
  ptrsForProjection.operatorEqualProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeProjection, gdc.VariantTypeProjection))
  ptrsForProjection.operatorNotEqualProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeProjection, gdc.VariantTypeProjection))
  ptrsForProjection.operatorMultiplyProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeProjection, gdc.VariantTypeProjection))
  ptrsForProjection.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeProjection, gdc.VariantTypeDictionary))
  ptrsForProjection.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeProjection, gdc.VariantTypeArray))
  {
    memberName := StringNameFromStr("x")
    defer memberName.Destroy()
    ptrsForProjection.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
    ptrsForProjection.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("y")
    defer memberName.Destroy()
    ptrsForProjection.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
    ptrsForProjection.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("z")
    defer memberName.Destroy()
    ptrsForProjection.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
    ptrsForProjection.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("w")
    defer memberName.Destroy()
    ptrsForProjection.memberwGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
    ptrsForProjection.memberwSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeProjection, memberName.AsCPtr()))
  }
  ptrsForProjection.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeProjection))
  ptrsForProjection.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeProjection))
}

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
  me.iface.CallPtrConstructor(ensurePtr(ptrsForProjection.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewProjectionFromProjection(from Projection, ) *Projection {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newProjection()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForProjection.ctrFromProjectionFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewProjectionFromTransform3D(from Transform3D, ) *Projection {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newProjection()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForProjection.ctrFromTransform3DFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewProjectionFromVector4Vector4Vector4Vector4(x_axis Vector4, y_axis Vector4, z_axis Vector4, w_axis Vector4, ) *Projection {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newProjection()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForProjection.ctrFromVector4Vector4Vector4Vector4Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), w_axis.AsCTypePtr(), }))
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
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForProjection.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Projection) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForProjection.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
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
  ret := NewProjection()

  varg0 := NewBoolFromBool(flip_y)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateDepthCorrectionFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateLightAtlasRect(rect Rect2, ) Projection {
  ret := NewProjection()


  args := []gdc.ConstTypePtr{rect.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateLightAtlasRectFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreatePerspective(fovy float64, aspect float64, z_near float64, z_far float64, flip_fov bool, ) Projection {
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


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreatePerspectiveFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreatePerspectiveHmd(fovy float64, aspect float64, z_near float64, z_far float64, flip_fov bool, eye int64, intraocular_dist float64, convergence_dist float64, ) Projection {
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


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreatePerspectiveHmdFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateForHmd(eye int64, aspect float64, intraocular_dist float64, display_width float64, display_to_lens float64, oversample float64, z_near float64, z_far float64, ) Projection {
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


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateForHmdFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateOrthogonal(left float64, right float64, bottom float64, top float64, z_near float64, z_far float64, ) Projection {
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


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateOrthogonalFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateOrthogonalAspect(size float64, aspect float64, z_near float64, z_far float64, flip_fov bool, ) Projection {
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


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateOrthogonalAspectFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateFrustum(left float64, right float64, bottom float64, top float64, z_near float64, z_far float64, ) Projection {
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


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateFrustumFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateFrustumAspect(size float64, aspect float64, offset Vector2, z_near float64, z_far float64, flip_fov bool, ) Projection {
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


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateFrustumAspectFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionCreateFitAabb(aabb AABB, ) Projection {
  ret := NewProjection()


  args := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodCreateFitAabbFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) Determinant() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodDeterminantFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) PerspectiveZnearAdjusted(new_znear float64, ) Projection {
  ret := NewProjection()

  varg0 := NewFloatFromFloat32(new_znear)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodPerspectiveZnearAdjustedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) GetProjectionPlane(plane int64, ) Plane {
  ret := NewPlane()

  varg0 := NewIntFromInt(plane)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetProjectionPlaneFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) FlippedY() Projection {
  ret := NewProjection()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodFlippedYFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) JitterOffseted(offset Vector2, ) Projection {
  ret := NewProjection()


  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodJitterOffsetedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ProjectionGetFovy(fovx float64, aspect float64, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(fovx)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(aspect)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetFovyFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetZFar() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetZFarFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetZNear() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetZNearFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetAspect() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetAspectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetFov() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetFovFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) IsOrthogonal() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodIsOrthogonalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetViewportHalfExtents() Vector2 {
  ret := NewVector2()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetViewportHalfExtentsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) GetFarPlaneHalfExtents() Vector2 {
  ret := NewVector2()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetFarPlaneHalfExtentsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) Inverse() Projection {
  ret := NewProjection()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Projection) GetPixelsPerMeter(for_pixel_width int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(for_pixel_width)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetPixelsPerMeterFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Projection) GetLodMultiplier() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForProjection.methodGetLodMultiplierFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *Projection) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) Not() bool {
  opPtr := ptrsForProjection.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) MultiplyVector4(right Vector4) Vector4 {
  opPtr := ptrsForProjection.operatorMultiplyVector4Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) EqualProjection(right Projection) bool {
  opPtr := ptrsForProjection.operatorEqualProjectionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) NotEqualProjection(right Projection) bool {
  opPtr := ptrsForProjection.operatorNotEqualProjectionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) MultiplyProjection(right Projection) Projection {
  opPtr := ptrsForProjection.operatorMultiplyProjectionFn
  ret := NewProjection()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) InDictionary(right Dictionary) bool {
  opPtr := ptrsForProjection.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Projection) InArray(right Array) bool {
  opPtr := ptrsForProjection.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Projection) X() Vector4 {
  ret := NewVector4()
  me.iface.CallPtrGetter(ensurePtr(ptrsForProjection.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetX(value Vector4) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForProjection.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Projection) Y() Vector4 {
  ret := NewVector4()
  me.iface.CallPtrGetter(ensurePtr(ptrsForProjection.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetY(value Vector4) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForProjection.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Projection) Z() Vector4 {
  ret := NewVector4()
  me.iface.CallPtrGetter(ensurePtr(ptrsForProjection.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetZ(value Vector4) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForProjection.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Projection) W() Vector4 {
  ret := NewVector4()
  me.iface.CallPtrGetter(ensurePtr(ptrsForProjection.memberwGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Projection) SetW(value Vector4) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForProjection.memberwSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
