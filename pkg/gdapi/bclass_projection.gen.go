// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Projection struct {
  iface gdc.Interface
  ptr gdc.TypePtr
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

func NewProjection() Projection {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeProjection))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Projection{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewProjectionFromProjection(from Projection, ) Projection {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeProjection))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Projection{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewProjectionFromTransform3D(from Transform3D, ) Projection {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeProjection))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Projection{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewProjectionFromVector4Vector4Vector4Vector4(x_axis Vector4, y_axis Vector4, z_axis Vector4, w_axis Vector4, ) Projection {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeProjection))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeProjection, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), w_axis.AsCTypePtr(), }))
  return Projection{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Projection) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Projection) Type() gdc.VariantType {
  return gdc.VariantTypeProjection
}

func (me *Projection) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Projection) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  ProjectionCreateDepthCorrection(flip_y bool, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreateLightAtlasRect(rect Rect2, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreatePerspective(fovy float32, aspect float32, z_near float32, z_far float32, flip_fov bool, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreatePerspectiveHmd(fovy float32, aspect float32, z_near float32, z_far float32, flip_fov bool, eye int, intraocular_dist float32, convergence_dist float32, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreateForHmd(eye int, aspect float32, intraocular_dist float32, display_width float32, display_to_lens float32, oversample float32, z_near float32, z_far float32, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreateOrthogonal(left float32, right float32, bottom float32, top float32, z_near float32, z_far float32, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreateOrthogonalAspect(size float32, aspect float32, z_near float32, z_far float32, flip_fov bool, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreateFrustum(left float32, right float32, bottom float32, top float32, z_near float32, z_far float32, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreateFrustumAspect(size float32, aspect float32, offset Vector2, z_near float32, z_far float32, flip_fov bool, ) Projection {
  panic("TODO: implement")
}

func  ProjectionCreateFitAabb(aabb AABB, ) Projection {
  panic("TODO: implement")
}

func  (me *Projection) Determinant() float32 {
  panic("TODO: implement")
}

func  (me *Projection) PerspectiveZnearAdjusted(new_znear float32, ) Projection {
  panic("TODO: implement")
}

func  (me *Projection) GetProjectionPlane(plane int, ) Plane {
  panic("TODO: implement")
}

func  (me *Projection) FlippedY() Projection {
  panic("TODO: implement")
}

func  (me *Projection) JitterOffseted(offset Vector2, ) Projection {
  panic("TODO: implement")
}

func  ProjectionGetFovy(fovx float32, aspect float32, ) float32 {
  panic("TODO: implement")
}

func  (me *Projection) GetZFar() float32 {
  panic("TODO: implement")
}

func  (me *Projection) GetZNear() float32 {
  panic("TODO: implement")
}

func  (me *Projection) GetAspect() float32 {
  panic("TODO: implement")
}

func  (me *Projection) GetFov() float32 {
  panic("TODO: implement")
}

func  (me *Projection) IsOrthogonal() bool {
  panic("TODO: implement")
}

func  (me *Projection) GetViewportHalfExtents() Vector2 {
  panic("TODO: implement")
}

func  (me *Projection) GetFarPlaneHalfExtents() Vector2 {
  panic("TODO: implement")
}

func  (me *Projection) Inverse() Projection {
  panic("TODO: implement")
}

func  (me *Projection) GetPixelsPerMeter(for_pixel_width int, ) int {
  panic("TODO: implement")
}

func  (me *Projection) GetLodMultiplier() float32 {
  panic("TODO: implement")
}

// Operators

func (me *Projection) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Projection) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Projection) Not() bool {
  panic("TODO: implement")
}

func (me *Projection) MultiplyVector4(right Vector4) Vector4 {
  panic("TODO: implement")
}

func (me *Projection) EqualsProjection(right Projection) bool {
  panic("TODO: implement")
}

func (me *Projection) NotEqualsProjection(right Projection) bool {
  panic("TODO: implement")
}

func (me *Projection) MultiplyProjection(right Projection) Projection {
  panic("TODO: implement")
}

func (me *Projection) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Projection) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
