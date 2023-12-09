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



// Enums

type CSGShape3DOperation int
const (
  CSGShape3DOperationOperationUnion CSGShape3DOperation = 0
  CSGShape3DOperationOperationIntersection CSGShape3DOperation = 1
  CSGShape3DOperationOperationSubtraction CSGShape3DOperation = 2
)

func (me *CSGShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGShape3D) IsRootShape()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetOperation(operation CSGShape3DOperation, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetOperation()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetSnap(snap float32, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetSnap()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetUseCollision(operation bool, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) IsUsingCollision()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetCollisionLayer(layer int, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetCollisionLayer()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetCollisionLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetCollisionLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetCollisionPriority(priority float32, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetCollisionPriority()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) SetCalculateTangents(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) IsCalculatingTangents()  {
  panic("TODO: implement")
}

func  (me *CSGShape3D) GetMeshes()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
