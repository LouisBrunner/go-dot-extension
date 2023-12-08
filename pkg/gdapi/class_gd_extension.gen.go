// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  GDExtensionInitializationLevelCore GDExtensionInitializationLevel = 0
  GDExtensionInitializationLevelServers GDExtensionInitializationLevel = 1
  GDExtensionInitializationLevelScene GDExtensionInitializationLevel = 2
  GDExtensionInitializationLevelEditor GDExtensionInitializationLevel = 3
)
