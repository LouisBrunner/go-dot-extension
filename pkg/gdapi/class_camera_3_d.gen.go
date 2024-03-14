// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Camera3D struct {
  Node3D
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

func (me *Camera3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Camera3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Camera3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Camera3D) ProjectRayNormal(screen_point Vector2, ) Vector3 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("project_ray_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1718073306) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(screen_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) ProjectLocalRayNormal(screen_point Vector2, ) Vector3 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("project_local_ray_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1718073306) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(screen_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) ProjectRayOrigin(screen_point Vector2, ) Vector3 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("project_ray_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1718073306) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(screen_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) UnprojectPosition(world_point Vector3, ) Vector2 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unproject_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3758901831) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(world_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) IsPositionBehind(world_point Vector3, ) bool {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_position_behind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3108956480) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(world_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) ProjectPosition(screen_point Vector2, z_depth float32, ) Vector3 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("project_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2171975744) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(screen_point.AsCTypePtr()), gdc.ConstTypePtr(&z_depth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetPerspective(fov float32, z_near float32, z_far float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2385087082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fov), gdc.ConstTypePtr(&z_near), gdc.ConstTypePtr(&z_far), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetOrthogonal(size float32, z_near float32, z_far float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_orthogonal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2385087082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&z_near), gdc.ConstTypePtr(&z_far), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetFrustum(size float32, offset Vector2, z_near float32, z_far float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frustum")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 354890663) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(offset.AsCTypePtr()), gdc.ConstTypePtr(&z_near), gdc.ConstTypePtr(&z_far), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) MakeCurrent()  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) ClearCurrent(enable_next bool, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable_next), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetCurrent(enabled bool, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) IsCurrent() bool {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetCameraTransform() Transform3D {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetCameraProjection() Projection {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_projection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2910717950) // FIXME: should cache?
  var ret Projection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetFov() float32 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetFrustumOffset() Vector2 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frustum_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetSize() float32 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetFar() float32 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetNear() float32 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetFov(fov float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fov), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetFrustumOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frustum_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetSize(size float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetFar(far float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&far), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetNear(near float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&near), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetProjection() Camera3DProjectionType {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_projection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2624185235) // FIXME: should cache?
  var ret Camera3DProjectionType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetProjection(mode Camera3DProjectionType, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_projection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4218540108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) SetHOffset(offset float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetHOffset() float32 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetVOffset(offset float32, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetVOffset() float32 {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetCullMask(mask int, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetCullMask() int {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetEnvironment(env Environment, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4143518816) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(env.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetEnvironment() Environment {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3082064660) // FIXME: should cache?
  var ret Environment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetAttributes(env CameraAttributes, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2817810567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(env.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetAttributes() CameraAttributes {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3921283215) // FIXME: should cache?
  var ret CameraAttributes
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetKeepAspectMode(mode Camera3DKeepAspect, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_aspect_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740651252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetKeepAspectMode() Camera3DKeepAspect {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keep_aspect_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2790278316) // FIXME: should cache?
  var ret Camera3DKeepAspect
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetDopplerTracking(mode Camera3DDopplerTracking, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_doppler_tracking")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3109431270) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetDopplerTracking() Camera3DDopplerTracking {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_doppler_tracking")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1584483649) // FIXME: should cache?
  var ret Camera3DDopplerTracking
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetFrustum() Plane {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frustum")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Plane
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) IsPositionInFrustum(world_point Vector3, ) bool {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_position_in_frustum")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3108956480) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(world_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetCameraRid() RID {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) GetPyramidShapeRid() RID {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pyramid_shape_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Camera3D) SetCullMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Camera3D) GetCullMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("Camera3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
