// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Lightmapper struct {
  obj gdc.ObjectPtr
}

func (me *Lightmapper) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Lightmapper) BaseClass() string {
  return "Lightmapper"
}



// Enums

func (me *Lightmapper) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Lightmapper) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
