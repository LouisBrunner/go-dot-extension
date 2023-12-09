// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorExportPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorExportPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorExportPlugin) BaseClass() string {
  return "EditorExportPlugin"
}



// Enums

func (me *EditorExportPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorExportPlugin) XExportFile(path String, type_ String, features PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XExportBegin(features PackedStringArray, is_debug bool, path String, flags int, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XExportEnd()  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XBeginCustomizeResources(platform EditorExportPlatform, features PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XCustomizeResource(resource Resource, path String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XBeginCustomizeScenes(platform EditorExportPlatform, features PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XCustomizeScene(scene Node, path String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XGetCustomizationConfigurationHash()  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XEndCustomizeScenes()  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XEndCustomizeResources()  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XGetExportOptions(platform EditorExportPlatform, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XShouldUpdateExportOptions(platform EditorExportPlatform, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XGetExportFeatures(platform EditorExportPlatform, debug bool, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) XGetName()  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddSharedObject(path String, tags PackedStringArray, target String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddIosProjectStaticLib(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddFile(path String, file PackedByteArray, remap bool, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddIosFramework(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddIosEmbeddedFramework(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddIosPlistContent(plist_content String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddIosLinkerFlags(flags String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddIosBundleFile(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddIosCppCode(code String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) AddMacosPluginFile(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) Skip()  {
  panic("TODO: implement")
}

func  (me *EditorExportPlugin) GetOption(name StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
