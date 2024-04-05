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

type ptrsForContainerList struct {
  fnXGetAllowedSizeFlagsHorizontal gdc.MethodBindPtr
  fnXGetAllowedSizeFlagsVertical gdc.MethodBindPtr
  fnQueueSort gdc.MethodBindPtr
  fnFitChildInRect gdc.MethodBindPtr
}

var ptrsForContainer ptrsForContainerList

func initContainerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Container")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("queue_sort")
    defer methodName.Destroy()
    ptrsForContainer.fnQueueSort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("fit_child_in_rect")
    defer methodName.Destroy()
    ptrsForContainer.fnFitChildInRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1993438598))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForContainer.fnQueueSort), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Container) FitChildInRect(child Control, rect Rect2, )  {
  cargs := []gdc.ConstTypePtr{child.AsCTypePtr(), rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForContainer.fnFitChildInRect), me.obj, unsafe.SliceData(cargs), nil)

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
