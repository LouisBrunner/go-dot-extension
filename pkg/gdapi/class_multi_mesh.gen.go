// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMesh struct {
  obj gdc.ObjectPtr
}

func (me *MultiMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMesh) BaseClass() string {
  return "MultiMesh"
}



// Enums

type MultiMeshTransformFormat int
const (
  MultiMeshTransformFormatTransform2D MultiMeshTransformFormat = 0
  MultiMeshTransformFormatTransform3D MultiMeshTransformFormat = 1
)

func (me *MultiMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiMesh) SetMesh(mesh Mesh, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetMesh()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetUseColors(enable bool, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) IsUsingColors()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetUseCustomData(enable bool, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) IsUsingCustomData()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetTransformFormat(format MultiMeshTransformFormat, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetTransformFormat()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetInstanceCount(count int, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetInstanceCount()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetVisibleInstanceCount(count int, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetVisibleInstanceCount()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetInstanceTransform(instance int, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetInstanceTransform2D(instance int, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetInstanceTransform(instance int, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetInstanceTransform2D(instance int, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetInstanceColor(instance int, color Color, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetInstanceColor(instance int, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetInstanceCustomData(instance int, custom_data Color, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetInstanceCustomData(instance int, )  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetAabb()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) GetBuffer()  {
  panic("TODO: implement")
}

func  (me *MultiMesh) SetBuffer(buffer PackedFloat32Array, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
