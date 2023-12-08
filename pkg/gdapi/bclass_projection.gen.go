// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Projection struct {
  obj gdc.ObjectPtr
}

func (me *Projection) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Projection) BaseClass() string {
  return "Projection"
}

// TODO: needed?
// const (
//   ProjectionPlaneNear = 0
//   ProjectionPlaneFar = 1
//   ProjectionPlaneLeft = 2
//   ProjectionPlaneTop = 3
//   ProjectionPlaneRight = 4
//   ProjectionPlaneBottom = 5
// // )

var (
  ProjectionIdentity = "Projection(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)" // TODO: construct correctly
  ProjectionZero = "Projection(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)" // TODO: construct correctly
)

type ProjectionPlanes int
const (
  ProjectionPlanesPlaneNear ProjectionPlanes = 0
  ProjectionPlanesPlaneFar ProjectionPlanes = 1
  ProjectionPlanesPlaneLeft ProjectionPlanes = 2
  ProjectionPlanesPlaneTop ProjectionPlanes = 3
  ProjectionPlanesPlaneRight ProjectionPlanes = 4
  ProjectionPlanesPlaneBottom ProjectionPlanes = 5
)

func  ProjectionCreateDepthCorrection(flip_y bool, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreateLightAtlasRect(rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreatePerspective(fovy float32, aspect float32, z_near float32, z_far float32, flip_fov bool, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreatePerspectiveHmd(fovy float32, aspect float32, z_near float32, z_far float32, flip_fov bool, eye int, intraocular_dist float32, convergence_dist float32, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreateForHmd(eye int, aspect float32, intraocular_dist float32, display_width float32, display_to_lens float32, oversample float32, z_near float32, z_far float32, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreateOrthogonal(left float32, right float32, bottom float32, top float32, z_near float32, z_far float32, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreateOrthogonalAspect(size float32, aspect float32, z_near float32, z_far float32, flip_fov bool, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreateFrustum(left float32, right float32, bottom float32, top float32, z_near float32, z_far float32, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreateFrustumAspect(size float32, aspect float32, offset Vector2, z_near float32, z_far float32, flip_fov bool, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionCreateFitAabb(aabb AABB, ) { // TODO: return value
  // TODO: implement
}

func  (me *Projection) Determinant() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) PerspectiveZnearAdjusted(new_znear float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetProjectionPlane(plane int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Projection) FlippedY() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) JitterOffseted(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  ProjectionGetFovy(fovx float32, aspect float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetZFar() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetZNear() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetAspect() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetFov() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) IsOrthogonal() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetViewportHalfExtents() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetFarPlaneHalfExtents() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) Inverse() { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetPixelsPerMeter(for_pixel_width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Projection) GetLodMultiplier() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
