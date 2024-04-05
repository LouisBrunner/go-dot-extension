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

type ptrsForVoxelGIList struct {
  fnSetProbeData gdc.MethodBindPtr
  fnGetProbeData gdc.MethodBindPtr
  fnSetSubdiv gdc.MethodBindPtr
  fnGetSubdiv gdc.MethodBindPtr
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetCameraAttributes gdc.MethodBindPtr
  fnGetCameraAttributes gdc.MethodBindPtr
  fnBake gdc.MethodBindPtr
  fnDebugBake gdc.MethodBindPtr
}

var ptrsForVoxelGI ptrsForVoxelGIList

func initVoxelGIPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VoxelGI")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_probe_data")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnSetProbeData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1637849675))
  }
  {
    methodName := StringNameFromStr("get_probe_data")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnGetProbeData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1730645405))
  }
  {
    methodName := StringNameFromStr("set_subdiv")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnSetSubdiv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240898472))
  }
  {
    methodName := StringNameFromStr("get_subdiv")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnGetSubdiv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4261647950))
  }
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_camera_attributes")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnSetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2817810567))
  }
  {
    methodName := StringNameFromStr("get_camera_attributes")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnGetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3921283215))
  }
  {
    methodName := StringNameFromStr("bake")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnBake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2781551026))
  }
  {
    methodName := StringNameFromStr("debug_bake")
    defer methodName.Destroy()
    ptrsForVoxelGI.fnDebugBake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

type VoxelGI struct {
  VisualInstance3D
}

func (me *VoxelGI) BaseClass() string {
  return "VoxelGI"
}

func NewVoxelGI() *VoxelGI {
  str := StringNameFromStr("VoxelGI") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VoxelGI{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VoxelGISubdiv int
const (
  VoxelGISubdivSubdiv64 VoxelGISubdiv = 0
  VoxelGISubdivSubdiv128 VoxelGISubdiv = 1
  VoxelGISubdivSubdiv256 VoxelGISubdiv = 2
  VoxelGISubdivSubdiv512 VoxelGISubdiv = 3
  VoxelGISubdivSubdivMax VoxelGISubdiv = 4
)

func (me *VoxelGI) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VoxelGI) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VoxelGI) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VoxelGI) SetProbeData(data VoxelGIData, )  {
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnSetProbeData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGI) GetProbeData() VoxelGIData {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVoxelGIData()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnGetProbeData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGI) SetSubdiv(subdiv VoxelGISubdiv, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdiv) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnSetSubdiv), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGI) GetSubdiv() VoxelGISubdiv {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VoxelGISubdiv

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnGetSubdiv), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VoxelGI) SetSize(size Vector3, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGI) GetSize() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGI) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  cargs := []gdc.ConstTypePtr{camera_attributes.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnSetCameraAttributes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGI) GetCameraAttributes() CameraAttributes {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCameraAttributes()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnGetCameraAttributes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGI) Bake(from_node Node, create_visual_debug bool, )  {
  cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&create_visual_debug) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnBake), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGI) DebugBake()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGI.fnDebugBake), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
