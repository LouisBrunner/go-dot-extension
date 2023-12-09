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



// Enums

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

func (me *ResourceSaver) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceSaver) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceSaver) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ResourceSaver) Save(resource Resource, path String, flags ResourceSaverSaverFlags, )  {
  panic("TODO: implement")
}

func  (me *ResourceSaver) GetRecognizedExtensions(type_ Resource, )  {
  panic("TODO: implement")
}

func  (me *ResourceSaver) AddResourceFormatSaver(format_saver ResourceFormatSaver, at_front bool, )  {
  panic("TODO: implement")
}

func  (me *ResourceSaver) RemoveResourceFormatSaver(format_saver ResourceFormatSaver, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
