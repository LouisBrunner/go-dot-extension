// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForGPUParticles3DList struct {
  fnSetEmitting gdc.MethodBindPtr
  fnSetAmount gdc.MethodBindPtr
  fnSetLifetime gdc.MethodBindPtr
  fnSetOneShot gdc.MethodBindPtr
  fnSetPreProcessTime gdc.MethodBindPtr
  fnSetExplosivenessRatio gdc.MethodBindPtr
  fnSetRandomnessRatio gdc.MethodBindPtr
  fnSetVisibilityAabb gdc.MethodBindPtr
  fnSetUseLocalCoordinates gdc.MethodBindPtr
  fnSetFixedFps gdc.MethodBindPtr
  fnSetFractionalDelta gdc.MethodBindPtr
  fnSetInterpolate gdc.MethodBindPtr
  fnSetProcessMaterial gdc.MethodBindPtr
  fnSetSpeedScale gdc.MethodBindPtr
  fnSetCollisionBaseSize gdc.MethodBindPtr
  fnSetInterpToEnd gdc.MethodBindPtr
  fnIsEmitting gdc.MethodBindPtr
  fnGetAmount gdc.MethodBindPtr
  fnGetLifetime gdc.MethodBindPtr
  fnGetOneShot gdc.MethodBindPtr
  fnGetPreProcessTime gdc.MethodBindPtr
  fnGetExplosivenessRatio gdc.MethodBindPtr
  fnGetRandomnessRatio gdc.MethodBindPtr
  fnGetVisibilityAabb gdc.MethodBindPtr
  fnGetUseLocalCoordinates gdc.MethodBindPtr
  fnGetFixedFps gdc.MethodBindPtr
  fnGetFractionalDelta gdc.MethodBindPtr
  fnGetInterpolate gdc.MethodBindPtr
  fnGetProcessMaterial gdc.MethodBindPtr
  fnGetSpeedScale gdc.MethodBindPtr
  fnGetCollisionBaseSize gdc.MethodBindPtr
  fnGetInterpToEnd gdc.MethodBindPtr
  fnSetDrawOrder gdc.MethodBindPtr
  fnGetDrawOrder gdc.MethodBindPtr
  fnSetDrawPasses gdc.MethodBindPtr
  fnSetDrawPassMesh gdc.MethodBindPtr
  fnGetDrawPasses gdc.MethodBindPtr
  fnGetDrawPassMesh gdc.MethodBindPtr
  fnSetSkin gdc.MethodBindPtr
  fnGetSkin gdc.MethodBindPtr
  fnRestart gdc.MethodBindPtr
  fnCaptureAabb gdc.MethodBindPtr
  fnSetSubEmitter gdc.MethodBindPtr
  fnGetSubEmitter gdc.MethodBindPtr
  fnEmitParticle gdc.MethodBindPtr
  fnSetTrailEnabled gdc.MethodBindPtr
  fnSetTrailLifetime gdc.MethodBindPtr
  fnIsTrailEnabled gdc.MethodBindPtr
  fnGetTrailLifetime gdc.MethodBindPtr
  fnSetTransformAlign gdc.MethodBindPtr
  fnGetTransformAlign gdc.MethodBindPtr
  fnConvertFromParticles gdc.MethodBindPtr
  fnSetAmountRatio gdc.MethodBindPtr
  fnGetAmountRatio gdc.MethodBindPtr
}

var ptrsForGPUParticles3D ptrsForGPUParticles3DList

func initGPUParticles3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GPUParticles3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_emitting")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_amount")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_lifetime")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_one_shot")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_pre_process_time")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_explosiveness_ratio")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_randomness_ratio")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_visibility_aabb")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetVisibilityAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
  }
  {
    methodName := StringNameFromStr("set_use_local_coordinates")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_fixed_fps")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_fractional_delta")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_interpolate")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_process_material")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetProcessMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
  }
  {
    methodName := StringNameFromStr("set_speed_scale")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_collision_base_size")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetCollisionBaseSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_interp_to_end")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetInterpToEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("is_emitting")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnIsEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_amount")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_lifetime")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_one_shot")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_pre_process_time")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_explosiveness_ratio")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_randomness_ratio")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_visibility_aabb")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetVisibilityAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
  }
  {
    methodName := StringNameFromStr("get_use_local_coordinates")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_fixed_fps")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_fractional_delta")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_interpolate")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_process_material")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetProcessMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
  }
  {
    methodName := StringNameFromStr("get_speed_scale")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_collision_base_size")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetCollisionBaseSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_interp_to_end")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetInterpToEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_draw_order")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1208074815))
  }
  {
    methodName := StringNameFromStr("get_draw_order")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3770381780))
  }
  {
    methodName := StringNameFromStr("set_draw_passes")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetDrawPasses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_draw_pass_mesh")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetDrawPassMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969122797))
  }
  {
    methodName := StringNameFromStr("get_draw_passes")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetDrawPasses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_draw_pass_mesh")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetDrawPassMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1576363275))
  }
  {
    methodName := StringNameFromStr("set_skin")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3971435618))
  }
  {
    methodName := StringNameFromStr("get_skin")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2074563878))
  }
  {
    methodName := StringNameFromStr("restart")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnRestart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("capture_aabb")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnCaptureAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
  }
  {
    methodName := StringNameFromStr("set_sub_emitter")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetSubEmitter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_sub_emitter")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetSubEmitter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("emit_particle")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnEmitParticle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 992173727))
  }
  {
    methodName := StringNameFromStr("set_trail_enabled")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetTrailEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_trail_lifetime")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetTrailLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("is_trail_enabled")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnIsTrailEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_trail_lifetime")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetTrailLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_transform_align")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetTransformAlign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3892425954))
  }
  {
    methodName := StringNameFromStr("get_transform_align")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetTransformAlign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2100992166))
  }
  {
    methodName := StringNameFromStr("convert_from_particles")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnConvertFromParticles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
  {
    methodName := StringNameFromStr("set_amount_ratio")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnSetAmountRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_amount_ratio")
    defer methodName.Destroy()
    ptrsForGPUParticles3D.fnGetAmountRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type GPUParticles3D struct {
  GeometryInstance3D
}

func (me *GPUParticles3D) BaseClass() string {
  return "GPUParticles3D"
}

func NewGPUParticles3D() *GPUParticles3D {
  str := StringNameFromStr("GPUParticles3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GPUParticles3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetEmitting), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetAmount(amount int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetLifetime(secs float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetLifetime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetOneShot(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetOneShot), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetPreProcessTime(secs float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetPreProcessTime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetExplosivenessRatio(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetRandomnessRatio(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetRandomnessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetVisibilityAabb(aabb AABB, )  {
  cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetVisibilityAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetUseLocalCoordinates(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetFixedFps(fps int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetFixedFps), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetFractionalDelta(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetFractionalDelta), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetInterpolate(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetInterpolate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetProcessMaterial(material Material, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetProcessMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetSpeedScale(scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetCollisionBaseSize(size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetCollisionBaseSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetInterpToEnd(interp float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interp) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetInterpToEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) IsEmitting() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnIsEmitting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetAmount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetLifetime() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetLifetime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetOneShot() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetOneShot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetPreProcessTime() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetPreProcessTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetExplosivenessRatio() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetRandomnessRatio() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetRandomnessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetVisibilityAabb() AABB {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetVisibilityAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticles3D) GetUseLocalCoordinates() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetFixedFps() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetFixedFps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetFractionalDelta() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetFractionalDelta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetInterpolate() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetProcessMaterial() Material {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetProcessMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticles3D) GetSpeedScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetCollisionBaseSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetCollisionBaseSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetInterpToEnd() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetInterpToEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) SetDrawOrder(order GPUParticles3DDrawOrder, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetDrawOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) GetDrawOrder() GPUParticles3DDrawOrder {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GPUParticles3DDrawOrder

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetDrawOrder), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GPUParticles3D) SetDrawPasses(passes int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&passes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetDrawPasses), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetDrawPassMesh(pass int64, mesh Mesh, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pass) , mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetDrawPassMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) GetDrawPasses() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetDrawPasses), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetDrawPassMesh(pass int64, ) Mesh {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pass) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()
  pinner.Pin(&pass)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetDrawPassMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticles3D) SetSkin(skin Skin, )  {
  cargs := []gdc.ConstTypePtr{skin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetSkin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) GetSkin() Skin {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetSkin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticles3D) Restart()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnRestart), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) CaptureAabb() AABB {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnCaptureAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticles3D) SetSubEmitter(path NodePath, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetSubEmitter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) GetSubEmitter() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetSubEmitter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticles3D) EmitParticle(xform Transform3D, velocity Vector3, color Color, custom Color, flags int64, )  {
  cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), velocity.AsCTypePtr(), color.AsCTypePtr(), custom.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnEmitParticle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetTrailEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetTrailEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetTrailLifetime(secs float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetTrailLifetime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) IsTrailEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnIsTrailEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) GetTrailLifetime() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetTrailLifetime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticles3D) SetTransformAlign(align GPUParticles3DTransformAlign, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&align) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetTransformAlign), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) GetTransformAlign() GPUParticles3DTransformAlign {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GPUParticles3DTransformAlign

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetTransformAlign), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GPUParticles3D) ConvertFromParticles(particles Node, )  {
  cargs := []gdc.ConstTypePtr{particles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnConvertFromParticles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) SetAmountRatio(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnSetAmountRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticles3D) GetAmountRatio() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles3D.fnGetAmountRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GPUParticles3DFinishedSignalFn func()

func (me *GPUParticles3D) ConnectFinished(subs SignalSubscribers, fn GPUParticles3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GPUParticles3D) DisconnectFinished(subs SignalSubscribers, fn GPUParticles3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
