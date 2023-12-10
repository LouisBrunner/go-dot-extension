// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGPrimitive3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGPrimitive3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGPrimitive3D) BaseClass() string {
  return "CSGPrimitive3D"
}



// Enums

func (me *CSGPrimitive3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGPrimitive3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGPrimitive3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGPrimitive3D) SetFlipFaces(flip_faces bool, )  {
  classNameV := StringNameFromStr("CSGPrimitive3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_faces), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGPrimitive3D) GetFlipFaces() bool {
  classNameV := StringNameFromStr("CSGPrimitive3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flip_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CSGPrimitive3D) GetPropFlipFaces() bool {
  panic("TODO: implement")
}

func (me *CSGPrimitive3D) SetPropFlipFaces(value bool) {
  panic("TODO: implement")
}