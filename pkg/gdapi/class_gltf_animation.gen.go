// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFAnimation struct {
  obj gdc.ObjectPtr
}

func (me *GLTFAnimation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFAnimation) BaseClass() string {
  return "GLTFAnimation"
}



// Enums

func (me *GLTFAnimation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFAnimation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GLTFAnimation) GetLoop()  {
  panic("TODO: implement")
}

func  (me *GLTFAnimation) SetLoop(loop bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
