// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JavaClass struct {
  obj gdc.ObjectPtr
}

func (me *JavaClass) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JavaClass) BaseClass() string {
  return "JavaClass"
}



// Enums

func (me *JavaClass) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaClass) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaClass) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
