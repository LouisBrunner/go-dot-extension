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

type MultiMeshTransformFormat int
const (
  MultiMeshTransformFormatTransform2D MultiMeshTransformFormat = 0
  MultiMeshTransformFormatTransform3D MultiMeshTransformFormat = 1
)

func  (me *MultiMesh) SetMesh(mesh Mesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetMesh() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetUseColors(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) IsUsingColors() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetUseCustomData(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) IsUsingCustomData() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetTransformFormat(format MultiMeshTransformFormat, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetTransformFormat() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetInstanceCount(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetInstanceCount() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetVisibleInstanceCount(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetVisibleInstanceCount() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetInstanceTransform(instance int, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetInstanceTransform2D(instance int, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetInstanceTransform(instance int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetInstanceTransform2D(instance int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetInstanceColor(instance int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetInstanceColor(instance int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetInstanceCustomData(instance int, custom_data Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetInstanceCustomData(instance int, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetAabb() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) GetBuffer() { // TODO: return value
  // TODO: implement
}

func  (me *MultiMesh) SetBuffer(buffer PackedFloat32Array, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
