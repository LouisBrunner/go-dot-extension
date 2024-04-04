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

type EditorInterface struct {
  Object
}

func (me *EditorInterface) BaseClass() string {
  return "EditorInterface"
}

func NewEditorInterface() *EditorInterface {
  str := StringNameFromStr("EditorInterface") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorInterface{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorInterface) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorInterface) RestartEditor(save bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("restart_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&save) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetCommandPalette() EditorCommandPalette {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_command_palette")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2471163807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorCommandPalette()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetResourceFilesystem() EditorFileSystem {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource_filesystem")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 780151678) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorFileSystem()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorPaths() EditorPaths {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1595760068) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorPaths()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetResourcePreviewer() EditorResourcePreview {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource_previewer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 943486957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorResourcePreview()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetSelection() EditorSelection {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2690272531) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorSelection()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorSettings() EditorSettings {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_settings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4086932459) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorSettings()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) MakeMeshPreviews(meshes []Mesh, preview_size int64, ) []Texture2D {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_mesh_previews")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 878078554) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&meshes) , gdc.ConstTypePtr(&preview_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&meshes)
  pinner.Pin(&preview_size)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Texture2D](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *EditorInterface) SetPluginEnabled(plugin String, enabled bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_plugin_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsPluginEnabled(plugin String, ) bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_plugin_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) GetEditorTheme() Theme {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3846893731) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTheme()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetBaseControl() Control {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewControl()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorMainScreen() VBoxContainer {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_main_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706218421) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVBoxContainer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetScriptEditor() ScriptEditor {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 90868003) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewScriptEditor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorViewport2D() SubViewport {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_viewport_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3750751911) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSubViewport()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorViewport3D(idx int64, ) SubViewport {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_viewport_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1970834490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSubViewport()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SetMainScreenEditor(name String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_main_screen_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) SetDistractionFreeMode(enter bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distraction_free_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsDistractionFreeModeEnabled() bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_distraction_free_mode_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) GetEditorScale() float64 {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) PopupDialog(dialog Window, rect Rect2i, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2015770942) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PopupDialogCentered(dialog Window, minsize Vector2i, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 346557367) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), minsize.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PopupDialogCenteredRatio(dialog Window, ratio float64, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog_centered_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2093669136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PopupDialogCenteredClamped(dialog Window, minsize Vector2i, fallback_ratio float64, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog_centered_clamped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3763385571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), minsize.AsCTypePtr(), gdc.ConstTypePtr(&fallback_ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetCurrentFeatureProfile() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_feature_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SetCurrentFeatureProfile(profile_name String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_feature_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{profile_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetFileSystemDock() FileSystemDock {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_system_dock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3751012327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFileSystemDock()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SelectFile(file String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetSelectedPaths() PackedStringArray {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetCurrentPath() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetCurrentDirectory() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_directory")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetInspector() EditorInspector {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inspector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3517113938) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorInspector()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) InspectObject(object Object, for_property String, inspector_only bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("inspect_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 127962172) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), for_property.AsCTypePtr(), gdc.ConstTypePtr(&inspector_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) EditResource(resource Resource, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("edit_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 968641751) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{resource.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) EditNode(node Node, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("edit_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) EditScript(script Script, line int64, column int64, grab_focus bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("edit_script")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 219829402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&grab_focus) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) OpenSceneFromPath(scene_filepath String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_scene_from_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_filepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) ReloadSceneFromPath(scene_filepath String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reload_scene_from_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_filepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetOpenScenes() PackedStringArray {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_scenes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditedSceneRoot() Node {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_scene_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SaveScene() Error {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *EditorInterface) SaveSceneAs(path String, with_preview bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_scene_as")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3647332257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&with_preview) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) SaveAllScenes()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_all_scenes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) MarkSceneAsUnsaved()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mark_scene_as_unsaved")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PlayMainScene()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_main_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PlayCurrentScene()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PlayCustomScene(scene_filepath String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_custom_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_filepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) StopPlayingScene()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop_playing_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsPlayingScene() bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) GetPlayingScene() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playing_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SetMovieMakerEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_movie_maker_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsMovieMakerEnabled() bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_movie_maker_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
