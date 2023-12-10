// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *EditorExportPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorExportPlugin) AddSharedObject(path String, tags PackedStringArray, target String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_shared_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3098291045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(tags.AsCTypePtr()), gdc.ConstTypePtr(target.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddIosProjectStaticLib(path String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ios_project_static_lib")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddFile(path String, file PackedByteArray, remap bool, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 527928637) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(file.AsCTypePtr()), gdc.ConstTypePtr(&remap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddIosFramework(path String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ios_framework")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddIosEmbeddedFramework(path String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ios_embedded_framework")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddIosPlistContent(plist_content String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ios_plist_content")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plist_content.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddIosLinkerFlags(flags String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ios_linker_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(flags.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddIosBundleFile(path String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ios_bundle_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddIosCppCode(code String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_ios_cpp_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(code.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) AddMacosPluginFile(path String, )  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_macos_plugin_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) Skip()  {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("skip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorExportPlugin) GetOption(name StringName, ) Variant {
  classNameV := StringNameFromStr("EditorExportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_option")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties