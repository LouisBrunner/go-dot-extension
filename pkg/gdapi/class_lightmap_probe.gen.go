// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightmapProbe struct {
  obj gdc.ObjectPtr
}

func (me *LightmapProbe) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightmapProbe) BaseClass() string {
  return "LightmapProbe"
}



// Enums

func (me *LightmapProbe) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightmapProbe) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
