// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VoxelGI struct {
  obj gdc.ObjectPtr
}

func (me *VoxelGI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VoxelGI) BaseClass() string {
  return "VoxelGI"
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
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_probe_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1637849675) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGI) GetProbeData() VoxelGIData {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_probe_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1730645405) // FIXME: should cache?
  var ret VoxelGIData
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGI) SetSubdiv(subdiv VoxelGISubdiv, )  {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subdiv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240898472) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdiv), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGI) GetSubdiv() VoxelGISubdiv {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subdiv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4261647950) // FIXME: should cache?
  var ret VoxelGISubdiv
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGI) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGI) GetSize() Vector3 {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGI) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2817810567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(camera_attributes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGI) GetCameraAttributes() CameraAttributes {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3921283215) // FIXME: should cache?
  var ret CameraAttributes
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGI) Bake(from_node Node, create_visual_debug bool, )  {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2781551026) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(&create_visual_debug), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGI) DebugBake()  {
  classNameV := StringNameFromStr("VoxelGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("debug_bake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
