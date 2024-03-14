// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorInterface struct {
  Object
}

func (me *EditorInterface) BaseClass() string {
  return "EditorInterface"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&save), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) GetCommandPalette() EditorCommandPalette {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_command_palette")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2471163807) // FIXME: should cache?
  var ret EditorCommandPalette
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetResourceFilesystem() EditorFileSystem {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource_filesystem")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 780151678) // FIXME: should cache?
  var ret EditorFileSystem
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditorPaths() EditorPaths {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1595760068) // FIXME: should cache?
  var ret EditorPaths
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetResourcePreviewer() EditorResourcePreview {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource_previewer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 943486957) // FIXME: should cache?
  var ret EditorResourcePreview
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetSelection() EditorSelection {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2690272531) // FIXME: should cache?
  var ret EditorSelection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditorSettings() EditorSettings {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_settings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4086932459) // FIXME: should cache?
  var ret EditorSettings
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) MakeMeshPreviews(meshes Mesh, preview_size int, ) Texture2D {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_mesh_previews")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 878078554) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(meshes.AsCTypePtr()), gdc.ConstTypePtr(&preview_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) SetPluginEnabled(plugin String, enabled bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_plugin_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plugin.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) IsPluginEnabled(plugin String, ) bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_plugin_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plugin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditorTheme() Theme {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3846893731) // FIXME: should cache?
  var ret Theme
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetBaseControl() Control {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditorMainScreen() VBoxContainer {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_main_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706218421) // FIXME: should cache?
  var ret VBoxContainer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetScriptEditor() ScriptEditor {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 90868003) // FIXME: should cache?
  var ret ScriptEditor
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditorViewport2D() SubViewport {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_viewport_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3750751911) // FIXME: should cache?
  var ret SubViewport
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditorViewport3D(idx int, ) SubViewport {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_viewport_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1970834490) // FIXME: should cache?
  var ret SubViewport
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) SetMainScreenEditor(name String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_main_screen_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) SetDistractionFreeMode(enter bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distraction_free_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) IsDistractionFreeModeEnabled() bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_distraction_free_mode_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditorScale() float32 {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) PopupDialog(dialog Window, rect Rect2i, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2015770942) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dialog.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) PopupDialogCentered(dialog Window, minsize Vector2i, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 346557367) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dialog.AsCTypePtr()), gdc.ConstTypePtr(minsize.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) PopupDialogCenteredRatio(dialog Window, ratio float32, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog_centered_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2093669136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dialog.AsCTypePtr()), gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) PopupDialogCenteredClamped(dialog Window, minsize Vector2i, fallback_ratio float32, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("popup_dialog_centered_clamped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3763385571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dialog.AsCTypePtr()), gdc.ConstTypePtr(minsize.AsCTypePtr()), gdc.ConstTypePtr(&fallback_ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) GetCurrentFeatureProfile() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_feature_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) SetCurrentFeatureProfile(profile_name String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_feature_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(profile_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) GetFileSystemDock() FileSystemDock {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_system_dock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3751012327) // FIXME: should cache?
  var ret FileSystemDock
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) SelectFile(file String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) GetSelectedPaths() PackedStringArray {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetCurrentPath() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetCurrentDirectory() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_directory")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetInspector() EditorInspector {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inspector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3517113938) // FIXME: should cache?
  var ret EditorInspector
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) InspectObject(object Object, for_property String, inspector_only bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("inspect_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 127962172) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(for_property.AsCTypePtr()), gdc.ConstTypePtr(&inspector_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) EditResource(resource Resource, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("edit_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 968641751) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(resource.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) EditNode(node Node, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("edit_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) EditScript(script Script, line int, column int, grab_focus bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("edit_script")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 219829402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(script.AsCTypePtr()), gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&grab_focus), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) OpenSceneFromPath(scene_filepath String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_scene_from_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scene_filepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) ReloadSceneFromPath(scene_filepath String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reload_scene_from_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scene_filepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) GetOpenScenes() PackedStringArray {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_scenes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetEditedSceneRoot() Node {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_scene_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) SaveScene() Error {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) SaveSceneAs(path String, with_preview bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_scene_as")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3647332257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&with_preview), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) SaveAllScenes()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_all_scenes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) MarkSceneAsUnsaved()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mark_scene_as_unsaved")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) PlayMainScene()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_main_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) PlayCurrentScene()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_current_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) PlayCustomScene(scene_filepath String, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_custom_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scene_filepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) StopPlayingScene()  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop_playing_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) IsPlayingScene() bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) GetPlayingScene() String {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playing_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorInterface) SetMovieMakerEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_movie_maker_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorInterface) IsMovieMakerEnabled() bool {
  classNameV := StringNameFromStr("EditorInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_movie_maker_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
