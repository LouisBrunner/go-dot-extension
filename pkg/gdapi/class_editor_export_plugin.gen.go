// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForEditorExportPluginList struct {
  fnXExportFile gdc.MethodBindPtr
  fnXExportBegin gdc.MethodBindPtr
  fnXExportEnd gdc.MethodBindPtr
  fnXBeginCustomizeResources gdc.MethodBindPtr
  fnXCustomizeResource gdc.MethodBindPtr
  fnXBeginCustomizeScenes gdc.MethodBindPtr
  fnXCustomizeScene gdc.MethodBindPtr
  fnXGetCustomizationConfigurationHash gdc.MethodBindPtr
  fnXEndCustomizeScenes gdc.MethodBindPtr
  fnXEndCustomizeResources gdc.MethodBindPtr
  fnXGetExportOptions gdc.MethodBindPtr
  fnXShouldUpdateExportOptions gdc.MethodBindPtr
  fnXGetExportOptionWarning gdc.MethodBindPtr
  fnXGetExportFeatures gdc.MethodBindPtr
  fnXGetName gdc.MethodBindPtr
  fnXSupportsPlatform gdc.MethodBindPtr
  fnXGetAndroidDependencies gdc.MethodBindPtr
  fnXGetAndroidDependenciesMavenRepos gdc.MethodBindPtr
  fnXGetAndroidLibraries gdc.MethodBindPtr
  fnXGetAndroidManifestActivityElementContents gdc.MethodBindPtr
  fnXGetAndroidManifestApplicationElementContents gdc.MethodBindPtr
  fnXGetAndroidManifestElementContents gdc.MethodBindPtr
  fnAddSharedObject gdc.MethodBindPtr
  fnAddIosProjectStaticLib gdc.MethodBindPtr
  fnAddFile gdc.MethodBindPtr
  fnAddIosFramework gdc.MethodBindPtr
  fnAddIosEmbeddedFramework gdc.MethodBindPtr
  fnAddIosPlistContent gdc.MethodBindPtr
  fnAddIosLinkerFlags gdc.MethodBindPtr
  fnAddIosBundleFile gdc.MethodBindPtr
  fnAddIosCppCode gdc.MethodBindPtr
  fnAddMacosPluginFile gdc.MethodBindPtr
  fnSkip gdc.MethodBindPtr
  fnGetOption gdc.MethodBindPtr
}

var ptrsForEditorExportPlugin ptrsForEditorExportPluginList

func initEditorExportPluginPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorExportPlugin")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_shared_object")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddSharedObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3098291045))
  }
  {
    methodName := StringNameFromStr("add_ios_project_static_lib")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddIosProjectStaticLib = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("add_file")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 527928637))
  }
  {
    methodName := StringNameFromStr("add_ios_framework")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddIosFramework = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("add_ios_embedded_framework")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddIosEmbeddedFramework = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("add_ios_plist_content")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddIosPlistContent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("add_ios_linker_flags")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddIosLinkerFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("add_ios_bundle_file")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddIosBundleFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("add_ios_cpp_code")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddIosCppCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("add_macos_plugin_file")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnAddMacosPluginFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("skip")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnSkip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_option")
    defer methodName.Destroy()
    ptrsForEditorExportPlugin.fnGetOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
  }
}

type EditorExportPlugin struct {
  RefCounted
}

func (me *EditorExportPlugin) BaseClass() string {
  return "EditorExportPlugin"
}

func NewEditorExportPlugin() *EditorExportPlugin {
  str := StringNameFromStr("EditorExportPlugin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlugin{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), tags.AsCTypePtr(), target.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddSharedObject), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddIosProjectStaticLib(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddIosProjectStaticLib), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddFile(path String, file PackedByteArray, remap bool, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), file.AsCTypePtr(), gdc.ConstTypePtr(&remap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddFile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddIosFramework(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddIosFramework), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddIosEmbeddedFramework(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddIosEmbeddedFramework), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddIosPlistContent(plist_content String, )  {
  cargs := []gdc.ConstTypePtr{plist_content.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddIosPlistContent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddIosLinkerFlags(flags String, )  {
  cargs := []gdc.ConstTypePtr{flags.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddIosLinkerFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddIosBundleFile(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddIosBundleFile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddIosCppCode(code String, )  {
  cargs := []gdc.ConstTypePtr{code.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddIosCppCode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) AddMacosPluginFile(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnAddMacosPluginFile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) Skip()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnSkip), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorExportPlugin) GetOption(name StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorExportPlugin.fnGetOption), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
