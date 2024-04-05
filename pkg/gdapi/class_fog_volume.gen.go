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

type ptrsForFogVolumeList struct {
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetShape gdc.MethodBindPtr
  fnGetShape gdc.MethodBindPtr
  fnSetMaterial gdc.MethodBindPtr
  fnGetMaterial gdc.MethodBindPtr
}

var ptrsForFogVolume ptrsForFogVolumeList

func initFogVolumePtrs(iface gdc.Interface) {

  className := StringNameFromStr("FogVolume")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForFogVolume.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForFogVolume.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_shape")
    defer methodName.Destroy()
    ptrsForFogVolume.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1416323362))
  }
  {
    methodName := StringNameFromStr("get_shape")
    defer methodName.Destroy()
    ptrsForFogVolume.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3920334604))
  }
  {
    methodName := StringNameFromStr("set_material")
    defer methodName.Destroy()
    ptrsForFogVolume.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
  }
  {
    methodName := StringNameFromStr("get_material")
    defer methodName.Destroy()
    ptrsForFogVolume.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
  }
}

type FogVolume struct {
  VisualInstance3D
}

func (me *FogVolume) BaseClass() string {
  return "FogVolume"
}

func NewFogVolume() *FogVolume {
  str := StringNameFromStr("FogVolume") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &FogVolume{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *FogVolume) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FogVolume) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FogVolume) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *FogVolume) SetSize(size Vector3, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogVolume.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogVolume) GetSize() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogVolume.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FogVolume) SetShape(shape RenderingServerFogVolumeShape, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogVolume.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogVolume) GetShape() RenderingServerFogVolumeShape {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingServerFogVolumeShape

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogVolume.fnGetShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FogVolume) SetMaterial(material Material, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogVolume.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FogVolume) GetMaterial() Material {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFogVolume.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
