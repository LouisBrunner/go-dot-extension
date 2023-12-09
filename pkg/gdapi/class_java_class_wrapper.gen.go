// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JavaClassWrapper struct {
  obj gdc.ObjectPtr
}

func (me *JavaClassWrapper) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JavaClassWrapper) BaseClass() string {
  return "JavaClassWrapper"
}



// Enums

func (me *JavaClassWrapper) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaClassWrapper) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaClassWrapper) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *JavaClassWrapper) Wrap(name String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
