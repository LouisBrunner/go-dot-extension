// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GridContainer struct {
  obj gdc.ObjectPtr
}

func (me *GridContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GridContainer) BaseClass() string {
  return "GridContainer"
}



// Enums

func (me *GridContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GridContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GridContainer) SetColumns(columns int, )  {
  panic("TODO: implement")
}

func  (me *GridContainer) GetColumns()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
