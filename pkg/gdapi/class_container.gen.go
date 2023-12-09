// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Container struct {
  obj gdc.ObjectPtr
}

func (me *Container) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Container) BaseClass() string {
  return "Container"
}



// Constants

var (
  ContainerNotificationPreSortChildren = "50" // TODO: construct correctly
  ContainerNotificationSortChildren = "51" // TODO: construct correctly
)

// Enums

func (me *Container) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Container) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Container) XGetAllowedSizeFlagsHorizontal()  {
  panic("TODO: implement")
}

func  (me *Container) XGetAllowedSizeFlagsVertical()  {
  panic("TODO: implement")
}

func  (me *Container) QueueSort()  {
  panic("TODO: implement")
}

func  (me *Container) FitChildInRect(child Control, rect Rect2, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
