// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  ResourceSaverSaverFlagsFlagNone ResourceSaverSaverFlags = 0
  ResourceSaverSaverFlagsFlagRelativePaths ResourceSaverSaverFlags = 1
  ResourceSaverSaverFlagsFlagBundleResources ResourceSaverSaverFlags = 2
  ResourceSaverSaverFlagsFlagChangePath ResourceSaverSaverFlags = 4
  ResourceSaverSaverFlagsFlagOmitEditorProperties ResourceSaverSaverFlags = 8
  ResourceSaverSaverFlagsFlagSaveBigEndian ResourceSaverSaverFlags = 16
  ResourceSaverSaverFlagsFlagCompress ResourceSaverSaverFlags = 32
  ResourceSaverSaverFlagsFlagReplaceSubresourcePaths ResourceSaverSaverFlags = 64
)

func  (me *ResourceSaver) Save(resource Resource, path String, flags ResourceSaverSaverFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceSaver) GetRecognizedExtensions(type_ Resource, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceSaver) AddResourceFormatSaver(format_saver ResourceFormatSaver, at_front bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceSaver) RemoveResourceFormatSaver(format_saver ResourceFormatSaver, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
