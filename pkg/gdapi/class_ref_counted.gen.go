// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RefCounted struct {
  obj gdc.ObjectPtr
}

func (me *RefCounted) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RefCounted) BaseClass() string {
  return "RefCounted"
}



// Enums

func (me *RefCounted) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RefCounted) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RefCounted) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RefCounted) InitRef()  {
  panic("TODO: implement")
}

func  (me *RefCounted) Reference()  {
  panic("TODO: implement")
}

func  (me *RefCounted) Unreference()  {
  panic("TODO: implement")
}

func  (me *RefCounted) GetReferenceCount()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
