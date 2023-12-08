// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  MainLoopNOTIFICATION_OS_MEMORY_WARNING = 2009
  MainLoopNOTIFICATION_TRANSLATION_CHANGED = 2010
  MainLoopNOTIFICATION_WM_ABOUT = 2011
  MainLoopNOTIFICATION_CRASH = 2012
  MainLoopNOTIFICATION_OS_IME_UPDATE = 2013
  MainLoopNOTIFICATION_APPLICATION_RESUMED = 2014
  MainLoopNOTIFICATION_APPLICATION_PAUSED = 2015
  MainLoopNOTIFICATION_APPLICATION_FOCUS_IN = 2016
  MainLoopNOTIFICATION_APPLICATION_FOCUS_OUT = 2017
  MainLoopNOTIFICATION_TEXT_SERVER_CHANGED = 2018
)
