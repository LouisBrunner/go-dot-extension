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

// TODO: needed?
// const (
// // )

var (
  ContainerNotificationPreSortChildren = "50" // TODO: construct correctly
  ContainerNotificationSortChildren = "51" // TODO: construct correctly
)

func  (me *Container) XGetAllowedSizeFlagsHorizontal() { // TODO: return value
  // TODO: implement
}

func  (me *Container) XGetAllowedSizeFlagsVertical() { // TODO: return value
  // TODO: implement
}

func  (me *Container) QueueSort() { // TODO: return value
  // TODO: implement
}

func  (me *Container) FitChildInRect(child Control, rect Rect2, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
