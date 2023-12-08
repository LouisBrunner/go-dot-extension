// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorFileDialog struct {
  obj gdc.ObjectPtr
}

func (me *EditorFileDialog) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorFileDialog) BaseClass() string {
  return "EditorFileDialog"
}

type EditorFileDialogFileMode int
const (
  EditorFileDialogFileModeFileModeOpenFile EditorFileDialogFileMode = 0
  EditorFileDialogFileModeFileModeOpenFiles EditorFileDialogFileMode = 1
  EditorFileDialogFileModeFileModeOpenDir EditorFileDialogFileMode = 2
  EditorFileDialogFileModeFileModeOpenAny EditorFileDialogFileMode = 3
  EditorFileDialogFileModeFileModeSaveFile EditorFileDialogFileMode = 4
)

type EditorFileDialogAccess int
const (
  EditorFileDialogAccessAccessResources EditorFileDialogAccess = 0
  EditorFileDialogAccessAccessUserdata EditorFileDialogAccess = 1
  EditorFileDialogAccessAccessFilesystem EditorFileDialogAccess = 2
)

type EditorFileDialogDisplayMode int
const (
  EditorFileDialogDisplayModeDisplayThumbnails EditorFileDialogDisplayMode = 0
  EditorFileDialogDisplayModeDisplayList EditorFileDialogDisplayMode = 1
)

func  (me *EditorFileDialog) ClearFilters() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) AddFilter(filter String, description String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetFilters(filters PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetFilters() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetCurrentDir() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetCurrentFile() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetCurrentPath() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetCurrentDir(dir String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetCurrentFile(file String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetCurrentPath(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetFileMode(mode EditorFileDialogFileMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetFileMode() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetVbox() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetLineEdit() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetAccess(access EditorFileDialogAccess, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetAccess() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetShowHiddenFiles(show bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) IsShowingHiddenFiles() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetDisplayMode(mode EditorFileDialogDisplayMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) GetDisplayMode() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) SetDisableOverwriteWarning(disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) IsOverwriteWarningDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *EditorFileDialog) Invalidate() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
