// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGShape3D) BaseClass() string {
  return "CSGShape3D"
}

type CSGShape3DOperation int
const (
  CSGShape3DOperationOperationUnion CSGShape3DOperation = 0
  CSGShape3DOperationOperationIntersection CSGShape3DOperation = 1
  CSGShape3DOperationOperationSubtraction CSGShape3DOperation = 2
)

func  (me *CSGShape3D) IsRootShape() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetOperation(operation CSGShape3DOperation, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetOperation() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetSnap(snap float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetSnap() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetUseCollision(operation bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) IsUsingCollision() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetCollisionLayer(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetCollisionLayer() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetCollisionMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetCollisionMask() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetCollisionMaskValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetCollisionMaskValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetCollisionLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetCollisionLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetCollisionPriority(priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetCollisionPriority() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) SetCalculateTangents(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) IsCalculatingTangents() { // TODO: return value
  // TODO: implement
}

func  (me *CSGShape3D) GetMeshes() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
