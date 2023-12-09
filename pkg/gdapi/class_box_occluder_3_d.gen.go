// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *BoxOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoxOccluder3D) BaseClass() string {
  return "BoxOccluder3D"
}



// Enums

func (me *BoxOccluder3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoxOccluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoxOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BoxOccluder3D) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *BoxOccluder3D) GetSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
