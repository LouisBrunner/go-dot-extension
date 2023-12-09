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



// Enums

type GDExtensionManagerLoadStatus int
const (
  GDExtensionManagerLoadStatusLoadStatusOk GDExtensionManagerLoadStatus = 0
  GDExtensionManagerLoadStatusLoadStatusFailed GDExtensionManagerLoadStatus = 1
  GDExtensionManagerLoadStatusLoadStatusAlreadyLoaded GDExtensionManagerLoadStatus = 2
  GDExtensionManagerLoadStatusLoadStatusNotLoaded GDExtensionManagerLoadStatus = 3
  GDExtensionManagerLoadStatusLoadStatusNeedsRestart GDExtensionManagerLoadStatus = 4
)

func (me *GDExtensionManager) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GDExtensionManager) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDExtensionManager) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GDExtensionManager) LoadExtension(path String, )  {
  panic("TODO: implement")
}

func  (me *GDExtensionManager) ReloadExtension(path String, )  {
  panic("TODO: implement")
}

func  (me *GDExtensionManager) UnloadExtension(path String, )  {
  panic("TODO: implement")
}

func  (me *GDExtensionManager) IsExtensionLoaded(path String, )  {
  panic("TODO: implement")
}

func  (me *GDExtensionManager) GetLoadedExtensions()  {
  panic("TODO: implement")
}

func  (me *GDExtensionManager) GetExtension(path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
