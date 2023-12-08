// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDExtensionManager struct {
  obj gdc.ObjectPtr
}

func (me *GDExtensionManager) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GDExtensionManager) BaseClass() string {
  return "GDExtensionManager"
}

type GDExtensionManagerLoadStatus int
const (
  GDExtensionManagerLoadStatusLoadStatusOk GDExtensionManagerLoadStatus = 0
  GDExtensionManagerLoadStatusLoadStatusFailed GDExtensionManagerLoadStatus = 1
  GDExtensionManagerLoadStatusLoadStatusAlreadyLoaded GDExtensionManagerLoadStatus = 2
  GDExtensionManagerLoadStatusLoadStatusNotLoaded GDExtensionManagerLoadStatus = 3
  GDExtensionManagerLoadStatusLoadStatusNeedsRestart GDExtensionManagerLoadStatus = 4
)

func  (me *GDExtensionManager) LoadExtension(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GDExtensionManager) ReloadExtension(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GDExtensionManager) UnloadExtension(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GDExtensionManager) IsExtensionLoaded(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GDExtensionManager) GetLoadedExtensions() { // TODO: return value
  // TODO: implement
}

func  (me *GDExtensionManager) GetExtension(path String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
