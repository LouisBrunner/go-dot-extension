// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *CompressedTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedTexture3D) BaseClass() string {
  return "CompressedTexture3D"
}



// Enums

func (me *CompressedTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CompressedTexture3D) Load(path String, )  {
  panic("TODO: implement")
}

func  (me *CompressedTexture3D) GetLoadPath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
