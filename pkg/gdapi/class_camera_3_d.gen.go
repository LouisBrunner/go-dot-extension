// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera3D struct {
  obj gdc.ObjectPtr
}

func (me *Camera3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Camera3D) BaseClass() string {
  return "Camera3D"
}



// Enums

type Camera3DProjectionType int
const (
  Camera3DProjectionTypeProjectionPerspective Camera3DProjectionType = 0
  Camera3DProjectionTypeProjectionOrthogonal Camera3DProjectionType = 1
  Camera3DProjectionTypeProjectionFrustum Camera3DProjectionType = 2
)

type Camera3DKeepAspect int
const (
  Camera3DKeepAspectKeepWidth Camera3DKeepAspect = 0
  Camera3DKeepAspectKeepHeight Camera3DKeepAspect = 1
)

type Camera3DDopplerTracking int
const (
  Camera3DDopplerTrackingDopplerTrackingDisabled Camera3DDopplerTracking = 0
  Camera3DDopplerTrackingDopplerTrackingIdleStep Camera3DDopplerTracking = 1
  Camera3DDopplerTrackingDopplerTrackingPhysicsStep Camera3DDopplerTracking = 2
)

func (me *Camera3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Camera3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Camera3D) ProjectRayNormal(screen_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) ProjectLocalRayNormal(screen_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) ProjectRayOrigin(screen_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) UnprojectPosition(world_point Vector3, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) IsPositionBehind(world_point Vector3, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) ProjectPosition(screen_point Vector2, z_depth float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetPerspective(fov float32, z_near float32, z_far float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetOrthogonal(size float32, z_near float32, z_far float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetFrustum(size float32, offset Vector2, z_near float32, z_far float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) MakeCurrent()  {
  panic("TODO: implement")
}

func  (me *Camera3D) ClearCurrent(enable_next bool, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetCurrent(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) IsCurrent()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetCameraTransform()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetCameraProjection()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetFov()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetFrustumOffset()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetSize()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetFar()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetNear()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetFov(fov float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetFrustumOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetSize(size float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetFar(far float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetNear(near float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetProjection()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetProjection(mode Camera3DProjectionType, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetHOffset(offset float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetHOffset()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetVOffset(offset float32, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetVOffset()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetCullMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetCullMask()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetEnvironment(env Environment, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetEnvironment()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetAttributes(env CameraAttributes, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetAttributes()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetKeepAspectMode(mode Camera3DKeepAspect, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetKeepAspectMode()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetDopplerTracking(mode Camera3DDopplerTracking, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetDopplerTracking()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetFrustum()  {
  panic("TODO: implement")
}

func  (me *Camera3D) IsPositionInFrustum(world_point Vector3, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetCameraRid()  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetPyramidShapeRid()  {
  panic("TODO: implement")
}

func  (me *Camera3D) SetCullMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *Camera3D) GetCullMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
