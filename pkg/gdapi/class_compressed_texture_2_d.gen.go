// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedTexture2D struct {
  obj gdc.ObjectPtr
}

func (me *CompressedTexture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedTexture2D) BaseClass() string {
  return "CompressedTexture2D"
}



// Enums

func (me *CompressedTexture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CompressedTexture2D) Load(path String, )  {
  panic("TODO: implement")
}

func  (me *CompressedTexture2D) GetLoadPath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
