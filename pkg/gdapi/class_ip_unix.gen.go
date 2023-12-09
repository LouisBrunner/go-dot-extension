// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type IPUnix struct {
  obj gdc.ObjectPtr
}

func (me *IPUnix) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *IPUnix) BaseClass() string {
  return "IPUnix"
}



// Enums

func (me *IPUnix) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *IPUnix) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
