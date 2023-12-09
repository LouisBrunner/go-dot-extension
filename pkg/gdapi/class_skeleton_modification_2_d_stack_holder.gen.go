// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModification2DStackHolder struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2DStackHolder) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2DStackHolder) BaseClass() string {
  return "SkeletonModification2DStackHolder"
}



// Enums

func (me *SkeletonModification2DStackHolder) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DStackHolder) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DStackHolder) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SkeletonModification2DStackHolder) SetHeldModificationStack(held_modification_stack SkeletonModificationStack2D, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2DStackHolder) GetHeldModificationStack()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
