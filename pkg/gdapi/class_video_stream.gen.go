// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VideoStream struct {
  obj gdc.ObjectPtr
}

func (me *VideoStream) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VideoStream) BaseClass() string {
  return "VideoStream"
}



// Enums

func (me *VideoStream) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VideoStream) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VideoStream) XInstantiatePlayback()  {
  panic("TODO: implement")
}

func  (me *VideoStream) SetFile(file String, )  {
  panic("TODO: implement")
}

func  (me *VideoStream) GetFile()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
