// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *MainLoop) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MainLoop) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MainLoop) XInitialize()  {
  panic("TODO: implement")
}

func  (me *MainLoop) XPhysicsProcess(delta float32, )  {
  panic("TODO: implement")
}

func  (me *MainLoop) XProcess(delta float32, )  {
  panic("TODO: implement")
}

func  (me *MainLoop) XFinalize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
