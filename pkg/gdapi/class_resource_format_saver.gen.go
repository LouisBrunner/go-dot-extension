// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceFormatSaver struct {
  obj gdc.ObjectPtr
}

func (me *ResourceFormatSaver) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceFormatSaver) BaseClass() string {
  return "ResourceFormatSaver"
}



// Enums

func (me *ResourceFormatSaver) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceFormatSaver) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceFormatSaver) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ResourceFormatSaver) XSave(resource Resource, path String, flags int, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatSaver) XSetUid(path String, uid int, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatSaver) XRecognize(resource Resource, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatSaver) XGetRecognizedExtensions(resource Resource, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatSaver) XRecognizePath(resource Resource, path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
