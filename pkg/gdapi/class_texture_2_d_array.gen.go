// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture2DArray struct {
  obj gdc.ObjectPtr
}

func (me *Texture2DArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture2DArray) BaseClass() string {
  return "Texture2DArray"
}



// Enums

func (me *Texture2DArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture2DArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Texture2DArray) CreatePlaceholder()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
