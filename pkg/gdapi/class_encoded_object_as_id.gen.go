// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EncodedObjectAsID struct {
  obj gdc.ObjectPtr
}

func (me *EncodedObjectAsID) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EncodedObjectAsID) BaseClass() string {
  return "EncodedObjectAsID"
}



// Enums

func (me *EncodedObjectAsID) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EncodedObjectAsID) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EncodedObjectAsID) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EncodedObjectAsID) SetObjectId(id int, )  {
  panic("TODO: implement")
}

func  (me *EncodedObjectAsID) GetObjectId()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
