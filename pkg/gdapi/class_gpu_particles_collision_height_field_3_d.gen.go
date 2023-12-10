// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GPUParticlesCollisionHeightField3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollisionHeightField3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollisionHeightField3D) BaseClass() string {
  return "GPUParticlesCollisionHeightField3D"
}



// Enums

type GPUParticlesCollisionHeightField3DResolution int
const (
  GPUParticlesCollisionHeightField3DResolutionResolution256 GPUParticlesCollisionHeightField3DResolution = 0
  GPUParticlesCollisionHeightField3DResolutionResolution512 GPUParticlesCollisionHeightField3DResolution = 1
  GPUParticlesCollisionHeightField3DResolutionResolution1024 GPUParticlesCollisionHeightField3DResolution = 2
  GPUParticlesCollisionHeightField3DResolutionResolution2048 GPUParticlesCollisionHeightField3DResolution = 3
  GPUParticlesCollisionHeightField3DResolutionResolution4096 GPUParticlesCollisionHeightField3DResolution = 4
  GPUParticlesCollisionHeightField3DResolutionResolution8192 GPUParticlesCollisionHeightField3DResolution = 5
  GPUParticlesCollisionHeightField3DResolutionResolutionMax GPUParticlesCollisionHeightField3DResolution = 6
)

type GPUParticlesCollisionHeightField3DUpdateMode int
const (
  GPUParticlesCollisionHeightField3DUpdateModeUpdateModeWhenMoved GPUParticlesCollisionHeightField3DUpdateMode = 0
  GPUParticlesCollisionHeightField3DUpdateModeUpdateModeAlways GPUParticlesCollisionHeightField3DUpdateMode = 1
)

func (me *GPUParticlesCollisionHeightField3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesCollisionHeightField3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionHeightField3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GPUParticlesCollisionHeightField3D) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionHeightField3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionHeightField3D) SetResolution(resolution GPUParticlesCollisionHeightField3DResolution, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1009996517) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resolution), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionHeightField3D) GetResolution() GPUParticlesCollisionHeightField3DResolution {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1156065644) // FIXME: should cache?
  var ret GPUParticlesCollisionHeightField3DResolution
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionHeightField3D) SetUpdateMode(update_mode GPUParticlesCollisionHeightField3DUpdateMode, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 673680859) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionHeightField3D) GetUpdateMode() GPUParticlesCollisionHeightField3DUpdateMode {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1998141380) // FIXME: should cache?
  var ret GPUParticlesCollisionHeightField3DUpdateMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionHeightField3D) SetFollowCameraEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_follow_camera_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionHeightField3D) IsFollowCameraEnabled() bool {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_follow_camera_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *GPUParticlesCollisionHeightField3D) GetPropSize() Vector3 {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionHeightField3D) SetPropSize(value Vector3) {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionHeightField3D) GetPropResolution() int {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionHeightField3D) SetPropResolution(value int) {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionHeightField3D) GetPropUpdateMode() int {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionHeightField3D) SetPropUpdateMode(value int) {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionHeightField3D) GetPropFollowCameraEnabled() bool {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionHeightField3D) SetPropFollowCameraEnabled(value bool) {
  panic("TODO: implement")
}