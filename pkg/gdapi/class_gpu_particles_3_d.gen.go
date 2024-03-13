// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GPUParticles3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticles3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticles3D) BaseClass() string {
  return "GPUParticles3D"
}



// Constants

var (
  GPUParticles3DMaxDrawPasses = "4" // TODO: construct correctly
)

// Enums

type GPUParticles3DDrawOrder int
const (
  GPUParticles3DDrawOrderDrawOrderIndex GPUParticles3DDrawOrder = 0
  GPUParticles3DDrawOrderDrawOrderLifetime GPUParticles3DDrawOrder = 1
  GPUParticles3DDrawOrderDrawOrderReverseLifetime GPUParticles3DDrawOrder = 2
  GPUParticles3DDrawOrderDrawOrderViewDepth GPUParticles3DDrawOrder = 3
)

type GPUParticles3DEmitFlags int
const (
  GPUParticles3DEmitFlagsEmitFlagPosition GPUParticles3DEmitFlags = 1
  GPUParticles3DEmitFlagsEmitFlagRotationScale GPUParticles3DEmitFlags = 2
  GPUParticles3DEmitFlagsEmitFlagVelocity GPUParticles3DEmitFlags = 4
  GPUParticles3DEmitFlagsEmitFlagColor GPUParticles3DEmitFlags = 8
  GPUParticles3DEmitFlagsEmitFlagCustom GPUParticles3DEmitFlags = 16
)

type GPUParticles3DTransformAlign int
const (
  GPUParticles3DTransformAlignTransformAlignDisabled GPUParticles3DTransformAlign = 0
  GPUParticles3DTransformAlignTransformAlignZBillboard GPUParticles3DTransformAlign = 1
  GPUParticles3DTransformAlignTransformAlignYToVelocity GPUParticles3DTransformAlign = 2
  GPUParticles3DTransformAlignTransformAlignZBillboardYToVelocity GPUParticles3DTransformAlign = 3
)

func (me *GPUParticles3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticles3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticles3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GPUParticles3D) SetEmitting(emitting bool, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetAmount(amount int, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetLifetime(secs float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetOneShot(enable bool, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetPreProcessTime(secs float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetExplosivenessRatio(ratio float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetRandomnessRatio(ratio float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetVisibilityAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(aabb.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetUseLocalCoordinates(enable bool, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetFixedFps(fps int, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetFractionalDelta(enable bool, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetInterpolate(enable bool, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetProcessMaterial(material Material, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetSpeedScale(scale float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetCollisionBaseSize(size float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_base_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetInterpToEnd(interp float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interp_to_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) IsEmitting() bool {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_emitting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetAmount() int {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetLifetime() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetOneShot() bool {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetPreProcessTime() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pre_process_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetExplosivenessRatio() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_explosiveness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetRandomnessRatio() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_randomness_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetVisibilityAabb() AABB {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetUseLocalCoordinates() bool {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_local_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetFixedFps() int {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_fps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetFractionalDelta() bool {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fractional_delta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetInterpolate() bool {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetProcessMaterial() Material {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetSpeedScale() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetCollisionBaseSize() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_base_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetInterpToEnd() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interp_to_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) SetDrawOrder(order GPUParticles3DDrawOrder, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1208074815) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) GetDrawOrder() GPUParticles3DDrawOrder {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3770381780) // FIXME: should cache?
  var ret GPUParticles3DDrawOrder
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) SetDrawPasses(passes int, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_passes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&passes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetDrawPassMesh(pass int, mesh Mesh, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_pass_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969122797) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pass), gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) GetDrawPasses() int {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_passes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetDrawPassMesh(pass int, ) Mesh {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_pass_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1576363275) // FIXME: should cache?
  var ret Mesh
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pass), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) SetSkin(skin Skin, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3971435618) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) GetSkin() Skin {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2074563878) // FIXME: should cache?
  var ret Skin
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) Restart()  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("restart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) CaptureAabb() AABB {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("capture_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) SetSubEmitter(path NodePath, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sub_emitter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) GetSubEmitter() NodePath {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sub_emitter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) EmitParticle(xform Transform3D, velocity Vector3, color Color, custom Color, flags int, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("emit_particle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 992173727) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(xform.AsCTypePtr()), gdc.ConstTypePtr(velocity.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(custom.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetTrailEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trail_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetTrailLifetime(secs float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trail_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) IsTrailEnabled() bool {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_trail_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) GetTrailLifetime() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_trail_lifetime")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) SetTransformAlign(align GPUParticles3DTransformAlign, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform_align")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3892425954) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&align), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) GetTransformAlign() GPUParticles3DTransformAlign {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform_align")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2100992166) // FIXME: should cache?
  var ret GPUParticles3DTransformAlign
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticles3D) ConvertFromParticles(particles Node, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convert_from_particles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(particles.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) SetAmountRatio(ratio float32, )  {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_amount_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticles3D) GetAmountRatio() float32 {
  classNameV := StringNameFromStr("GPUParticles3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_amount_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GPUParticles3DFinishedSignalFn func()

func (me *GPUParticles3D) ConnectFinished(subs SignalSubscribers, fn GPUParticles3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GPUParticles3D) DisconnectFinished(subs SignalSubscribers, fn GPUParticles3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
