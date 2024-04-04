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

type HeightMapShape3D struct {
  Shape3D
}

func (me *HeightMapShape3D) BaseClass() string {
  return "HeightMapShape3D"
}

func NewHeightMapShape3D() *HeightMapShape3D {
  str := StringNameFromStr("HeightMapShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HeightMapShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *HeightMapShape3D) SetMapWidth(width int64, )  {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HeightMapShape3D) GetMapWidth() int64 {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HeightMapShape3D) SetMapDepth(height int64, )  {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HeightMapShape3D) GetMapDepth() int64 {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HeightMapShape3D) SetMapData(data PackedFloat32Array, )  {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HeightMapShape3D) GetMapData() PackedFloat32Array {
  classNameV := StringNameFromStr("HeightMapShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
