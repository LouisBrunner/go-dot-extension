// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type Container struct {
  Control
}

func (me *Container) BaseClass() string {
  return "Container"
}

func NewContainer() *Container {
  str := StringNameFromStr("Container") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Container{}
  obj.SetBaseObject(objPtr)
  return obj
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
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Container) FitChildInRect(child Control, rect Rect2, )  {
  classNameV := StringNameFromStr("Container")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fit_child_in_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1993438598) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{child.AsCTypePtr(), rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type ContainerPreSortChildrenSignalFn func()

func (me *Container) ConnectPreSortChildren(subs SignalSubscribers, fn ContainerPreSortChildrenSignalFn) {
  sig := StringNameFromStr("pre_sort_children")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Container) DisconnectPreSortChildren(subs SignalSubscribers, fn ContainerPreSortChildrenSignalFn) {
  sig := StringNameFromStr("pre_sort_children")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ContainerSortChildrenSignalFn func()

func (me *Container) ConnectSortChildren(subs SignalSubscribers, fn ContainerSortChildrenSignalFn) {
  sig := StringNameFromStr("sort_children")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Container) DisconnectSortChildren(subs SignalSubscribers, fn ContainerSortChildrenSignalFn) {
  sig := StringNameFromStr("sort_children")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
