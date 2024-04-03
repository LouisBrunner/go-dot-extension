// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGPrimitive3D struct {
  CSGShape3D
}

func (me *CSGPrimitive3D) BaseClass() string {
  return "CSGPrimitive3D"
}

func NewCSGPrimitive3D() *CSGPrimitive3D {
  str := StringNameFromStr("CSGPrimitive3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CSGPrimitive3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
