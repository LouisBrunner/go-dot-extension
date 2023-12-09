// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInspector struct {
  obj gdc.ObjectPtr
}

func (me *EditorInspector) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorInspector) BaseClass() string {
  return "EditorInspector"
}



// Enums

func (me *EditorInspector) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorInspector) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorInspector) GetSelectedPath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
