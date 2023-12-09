// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorResourceConversionPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorResourceConversionPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorResourceConversionPlugin) BaseClass() string {
  return "EditorResourceConversionPlugin"
}



// Enums

func (me *EditorResourceConversionPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorResourceConversionPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorResourceConversionPlugin) XConvertsTo()  {
  panic("TODO: implement")
}

func  (me *EditorResourceConversionPlugin) XHandles(resource Resource, )  {
  panic("TODO: implement")
}

func  (me *EditorResourceConversionPlugin) XConvert(resource Resource, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
