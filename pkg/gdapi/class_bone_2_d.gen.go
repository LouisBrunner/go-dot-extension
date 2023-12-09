// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Bone2D struct {
  obj gdc.ObjectPtr
}

func (me *Bone2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Bone2D) BaseClass() string {
  return "Bone2D"
}



// Enums

func (me *Bone2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Bone2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Bone2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Bone2D) SetRest(rest Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Bone2D) GetRest()  {
  panic("TODO: implement")
}

func  (me *Bone2D) ApplyRest()  {
  panic("TODO: implement")
}

func  (me *Bone2D) GetSkeletonRest()  {
  panic("TODO: implement")
}

func  (me *Bone2D) GetIndexInSkeleton()  {
  panic("TODO: implement")
}

func  (me *Bone2D) SetAutocalculateLengthAndAngle(auto_calculate bool, )  {
  panic("TODO: implement")
}

func  (me *Bone2D) GetAutocalculateLengthAndAngle()  {
  panic("TODO: implement")
}

func  (me *Bone2D) SetLength(length float32, )  {
  panic("TODO: implement")
}

func  (me *Bone2D) GetLength()  {
  panic("TODO: implement")
}

func  (me *Bone2D) SetBoneAngle(angle float32, )  {
  panic("TODO: implement")
}

func  (me *Bone2D) GetBoneAngle()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
