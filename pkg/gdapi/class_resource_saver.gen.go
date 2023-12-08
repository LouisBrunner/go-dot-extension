// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceSaver struct {
  obj gdc.ObjectPtr
}

func (me *ResourceSaver) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceSaver) BaseClass() string {
  return "ResourceSaver"
}

type ResourceSaverSaverFlags int
const (
  ResourceSaverFlagNone ResourceSaverSaverFlags = 0
  ResourceSaverFlagRelativePaths ResourceSaverSaverFlags = 1
  ResourceSaverFlagBundleResources ResourceSaverSaverFlags = 2
  ResourceSaverFlagChangePath ResourceSaverSaverFlags = 4
  ResourceSaverFlagOmitEditorProperties ResourceSaverSaverFlags = 8
  ResourceSaverFlagSaveBigEndian ResourceSaverSaverFlags = 16
  ResourceSaverFlagCompress ResourceSaverSaverFlags = 32
  ResourceSaverFlagReplaceSubresourcePaths ResourceSaverSaverFlags = 64
)
