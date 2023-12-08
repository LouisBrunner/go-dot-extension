// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInterface struct {
  obj gdc.ObjectPtr
}

func (me *EditorInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorInterface) BaseClass() string {
  return "EditorInterface"
}

func  (me *EditorInterface) RestartEditor(save bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetCommandPalette() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetResourceFilesystem() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetEditorPaths() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetResourcePreviewer() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetSelection() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetEditorSettings() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) MakeMeshPreviews(meshes Mesh, preview_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) SetPluginEnabled(plugin String, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) IsPluginEnabled(plugin String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetBaseControl() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetEditorMainScreen() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetScriptEditor() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) SetMainScreenEditor(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) SetDistractionFreeMode(enter bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) IsDistractionFreeModeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetEditorScale() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) PopupDialog(dialog Window, rect Rect2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) PopupDialogCentered(dialog Window, minsize Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) PopupDialogCenteredRatio(dialog Window, ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) PopupDialogCenteredClamped(dialog Window, minsize Vector2i, fallback_ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetFileSystemDock() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) SelectFile(file String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetSelectedPaths() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetCurrentPath() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetCurrentDirectory() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetInspector() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) InspectObject(object Object, for_property String, inspector_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) EditResource(resource Resource, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) EditNode(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) EditScript(script Script, line int, column int, grab_focus bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) OpenSceneFromPath(scene_filepath String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) ReloadSceneFromPath(scene_filepath String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetOpenScenes() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetEditedSceneRoot() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) SaveScene() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) SaveSceneAs(path String, with_preview bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) MarkSceneAsUnsaved() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) PlayMainScene() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) PlayCurrentScene() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) PlayCustomScene(scene_filepath String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) StopPlayingScene() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) IsPlayingScene() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) GetPlayingScene() { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) SetMovieMakerEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorInterface) IsMovieMakerEnabled() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
