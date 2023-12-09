// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDExtension struct {
  obj gdc.ObjectPtr
}

func (me *GDExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GDExtension) BaseClass() string {
  return "GDExtension"
}



// Enums

type GDExtensionInitializationLevel int
const (
  GDExtensionInitializationLevelInitializationLevelCore GDExtensionInitializationLevel = 0
  GDExtensionInitializationLevelInitializationLevelServers GDExtensionInitializationLevel = 1
  GDExtensionInitializationLevelInitializationLevelScene GDExtensionInitializationLevel = 2
  GDExtensionInitializationLevelInitializationLevelEditor GDExtensionInitializationLevel = 3
)

func (me *GDExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GDExtension) OpenLibrary(path String, entry_symbol String, )  {
  panic("TODO: implement")
}

func  (me *GDExtension) CloseLibrary()  {
  panic("TODO: implement")
}

func  (me *GDExtension) IsLibraryOpen()  {
  panic("TODO: implement")
}

func  (me *GDExtension) GetMinimumLibraryInitializationLevel()  {
  panic("TODO: implement")
}

func  (me *GDExtension) InitializeLibrary(level GDExtensionInitializationLevel, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
