// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MainLoop struct {
  obj gdc.ObjectPtr
}

func (me *MainLoop) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MainLoop) BaseClass() string {
  return "MainLoop"
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
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *MainLoop) DisconnectOnRequestPermissionsResult(subs SignalSubscribers, fn MainLoopOnRequestPermissionsResultSignalFn) {
  sig := StringNameFromStr("on_request_permissions_result")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
