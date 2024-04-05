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

type ptrsForMainLoopList struct {
  fnXInitialize gdc.MethodBindPtr
  fnXPhysicsProcess gdc.MethodBindPtr
  fnXProcess gdc.MethodBindPtr
  fnXFinalize gdc.MethodBindPtr
}

var ptrsForMainLoop ptrsForMainLoopList

func initMainLoopPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MainLoop")
  defer className.Destroy()
}

type MainLoop struct {
  Object
}

func (me *MainLoop) BaseClass() string {
  return "MainLoop"
}

func NewMainLoop() *MainLoop {
  str := StringNameFromStr("MainLoop") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MainLoop{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  MainLoopNotificationOsMemoryWarning = 2009
  MainLoopNotificationTranslationChanged = 2010
  MainLoopNotificationWmAbout = 2011
  MainLoopNotificationCrash = 2012
  MainLoopNotificationOsImeUpdate = 2013
  MainLoopNotificationApplicationResumed = 2014
  MainLoopNotificationApplicationPaused = 2015
  MainLoopNotificationApplicationFocusIn = 2016
  MainLoopNotificationApplicationFocusOut = 2017
  MainLoopNotificationTextServerChanged = 2018
)

// Enums

func (me *MainLoop) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MainLoop) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MainLoop) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

type MainLoopOnRequestPermissionsResultSignalFn func(permission String, granted bool, )

func (me *MainLoop) ConnectOnRequestPermissionsResult(subs SignalSubscribers, fn MainLoopOnRequestPermissionsResultSignalFn) {
  sig := StringNameFromStr("on_request_permissions_result")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *MainLoop) DisconnectOnRequestPermissionsResult(subs SignalSubscribers, fn MainLoopOnRequestPermissionsResultSignalFn) {
  sig := StringNameFromStr("on_request_permissions_result")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
