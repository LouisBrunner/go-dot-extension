// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *Container) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Container) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Container) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Container) QueueSort()  {
  classNameV := StringNameFromStr("Container")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_sort")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Container) FitChildInRect(child Control, rect Rect2, )  {
  classNameV := StringNameFromStr("Container")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fit_child_in_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1993438598) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(child.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API