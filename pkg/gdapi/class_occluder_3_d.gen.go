// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Occluder3D struct {
  obj gdc.ObjectPtr
}

func (me *Occluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Occluder3D) BaseClass() string {
  return "Occluder3D"
}



// Enums

func (me *Occluder3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Occluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Occluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Occluder3D) GetVertices()  {
  panic("TODO: implement")
}

func  (me *Occluder3D) GetIndices()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
