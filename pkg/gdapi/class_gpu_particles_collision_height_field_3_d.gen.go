// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type GPUParticlesCollisionHeightField3D struct {
  GPUParticlesCollision3D
}

func (me *GPUParticlesCollisionHeightField3D) BaseClass() string {
  return "GPUParticlesCollisionHeightField3D"
}

func NewGPUParticlesCollisionHeightField3D() *GPUParticlesCollisionHeightField3D {
  str := StringNameFromStr("GPUParticlesCollisionHeightField3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GPUParticlesCollisionHeightField3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionHeightField3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticlesCollisionHeightField3D) SetResolution(resolution GPUParticlesCollisionHeightField3DResolution, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1009996517) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resolution) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionHeightField3D) GetResolution() GPUParticlesCollisionHeightField3DResolution {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1156065644) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GPUParticlesCollisionHeightField3DResolution

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GPUParticlesCollisionHeightField3D) SetUpdateMode(update_mode GPUParticlesCollisionHeightField3DUpdateMode, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 673680859) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionHeightField3D) GetUpdateMode() GPUParticlesCollisionHeightField3DUpdateMode {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1998141380) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GPUParticlesCollisionHeightField3DUpdateMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GPUParticlesCollisionHeightField3D) SetFollowCameraEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_follow_camera_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionHeightField3D) IsFollowCameraEnabled() bool {
  classNameV := StringNameFromStr("GPUParticlesCollisionHeightField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_follow_camera_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
