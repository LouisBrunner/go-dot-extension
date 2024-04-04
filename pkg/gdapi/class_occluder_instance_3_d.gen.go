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

type OccluderInstance3D struct {
  Node3D
}

func (me *OccluderInstance3D) BaseClass() string {
  return "OccluderInstance3D"
}

func NewOccluderInstance3D() *OccluderInstance3D {
  str := StringNameFromStr("OccluderInstance3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OccluderInstance3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *OccluderInstance3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OccluderInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OccluderInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OccluderInstance3D) SetBakeMask(mask int64, )  {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OccluderInstance3D) GetBakeMask() int64 {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OccluderInstance3D) SetBakeMaskValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OccluderInstance3D) GetBakeMaskValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OccluderInstance3D) SetBakeSimplificationDistance(simplification_distance float64, )  {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_simplification_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&simplification_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OccluderInstance3D) GetBakeSimplificationDistance() float64 {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_simplification_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OccluderInstance3D) SetOccluder(occluder Occluder3D, )  {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_occluder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1664878165) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OccluderInstance3D) GetOccluder() Occluder3D {
  classNameV := StringNameFromStr("OccluderInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occluder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1696836198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewOccluder3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
