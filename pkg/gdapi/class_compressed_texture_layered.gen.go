// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedTextureLayered struct {
  obj gdc.ObjectPtr
}

func (me *CompressedTextureLayered) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedTextureLayered) BaseClass() string {
  return "CompressedTextureLayered"
}



// Enums

func (me *CompressedTextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedTextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CompressedTextureLayered) Load(path String, )  {
  panic("TODO: implement")
}

func  (me *CompressedTextureLayered) GetLoadPath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
