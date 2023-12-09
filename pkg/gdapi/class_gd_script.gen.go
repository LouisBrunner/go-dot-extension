// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDScript struct {
  obj gdc.ObjectPtr
}

func (me *GDScript) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GDScript) BaseClass() string {
  return "GDScript"
}



// Enums

func (me *GDScript) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GDScript) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDScript) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GDScript) New()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
