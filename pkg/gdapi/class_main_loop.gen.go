// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  MainLoopNotificationOsMemoryWarning = "2009" // TODO: construct correctly
  MainLoopNotificationTranslationChanged = "2010" // TODO: construct correctly
  MainLoopNotificationWmAbout = "2011" // TODO: construct correctly
  MainLoopNotificationCrash = "2012" // TODO: construct correctly
  MainLoopNotificationOsImeUpdate = "2013" // TODO: construct correctly
  MainLoopNotificationApplicationResumed = "2014" // TODO: construct correctly
  MainLoopNotificationApplicationPaused = "2015" // TODO: construct correctly
  MainLoopNotificationApplicationFocusIn = "2016" // TODO: construct correctly
  MainLoopNotificationApplicationFocusOut = "2017" // TODO: construct correctly
  MainLoopNotificationTextServerChanged = "2018" // TODO: construct correctly
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
