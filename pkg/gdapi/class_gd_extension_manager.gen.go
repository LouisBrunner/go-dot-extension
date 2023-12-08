// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  GDExtensionManagerLoadStatusOk GDExtensionManagerLoadStatus = 0
  GDExtensionManagerLoadStatusFailed GDExtensionManagerLoadStatus = 1
  GDExtensionManagerLoadStatusAlreadyLoaded GDExtensionManagerLoadStatus = 2
  GDExtensionManagerLoadStatusNotLoaded GDExtensionManagerLoadStatus = 3
  GDExtensionManagerLoadStatusNeedsRestart GDExtensionManagerLoadStatus = 4
)
