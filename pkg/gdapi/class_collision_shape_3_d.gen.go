// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionShape3D) BaseClass() string {
  return "CollisionShape3D"
}



// Enums

func (me *CollisionShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CollisionShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CollisionShape3D) ResourceChanged(resource Resource, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape3D) SetShape(shape Shape3D, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape3D) GetShape()  {
  panic("TODO: implement")
}

func  (me *CollisionShape3D) SetDisabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionShape3D) IsDisabled()  {
  panic("TODO: implement")
}

func  (me *CollisionShape3D) MakeConvexFromSiblings()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
