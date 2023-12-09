// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkinReference struct {
  obj gdc.ObjectPtr
}

func (me *SkinReference) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkinReference) BaseClass() string {
  return "SkinReference"
}



// Enums

func (me *SkinReference) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkinReference) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SkinReference) GetSkeleton()  {
  panic("TODO: implement")
}

func  (me *SkinReference) GetSkin()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
