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

type GDExtensionInitializationLevel int
const (
  GDExtensionInitializationLevelInitializationLevelCore GDExtensionInitializationLevel = 0
  GDExtensionInitializationLevelInitializationLevelServers GDExtensionInitializationLevel = 1
  GDExtensionInitializationLevelInitializationLevelScene GDExtensionInitializationLevel = 2
  GDExtensionInitializationLevelInitializationLevelEditor GDExtensionInitializationLevel = 3
)

func  (me *GDExtension) OpenLibrary(path String, entry_symbol String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GDExtension) CloseLibrary() { // TODO: return value
  // TODO: implement
}

func  (me *GDExtension) IsLibraryOpen() { // TODO: return value
  // TODO: implement
}

func  (me *GDExtension) GetMinimumLibraryInitializationLevel() { // TODO: return value
  // TODO: implement
}

func  (me *GDExtension) InitializeLibrary(level GDExtensionInitializationLevel, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
