// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRAnchor3D struct {
  obj gdc.ObjectPtr
}

func (me *XRAnchor3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRAnchor3D) BaseClass() string {
  return "XRAnchor3D"
}



// Enums

func (me *XRAnchor3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRAnchor3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRAnchor3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XRAnchor3D) GetSize()  {
  panic("TODO: implement")
}

func  (me *XRAnchor3D) GetPlane()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
