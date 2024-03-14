// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HeightMapShape3D struct {
  Shape3D
}

func (me *HeightMapShape3D) BaseClass() string {
  return "HeightMapShape3D"
}



// Enums

func (me *HeightMapShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HeightMapShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HeightMapShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *HeightMapShape3D) SetMapWidth(width int, )  {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *HeightMapShape3D) GetMapWidth() int {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *HeightMapShape3D) SetMapDepth(height int, )  {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *HeightMapShape3D) GetMapDepth() int {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *HeightMapShape3D) SetMapData(data PackedFloat32Array, )  {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *HeightMapShape3D) GetMapData() PackedFloat32Array {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
