// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderTexture3D) BaseClass() string {
  return "PlaceholderTexture3D"
}



// Enums

func (me *PlaceholderTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PlaceholderTexture3D) SetSize(size Vector3i, )  {
  panic("TODO: implement")
}

func  (me *PlaceholderTexture3D) GetSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
