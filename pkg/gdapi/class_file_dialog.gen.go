// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileDialog struct {
  obj gdc.ObjectPtr
}

func (me *FileDialog) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FileDialog) BaseClass() string {
  return "FileDialog"
}

type FileDialogFileMode int
const (
  FileDialogFileModeFileModeOpenFile FileDialogFileMode = 0
  FileDialogFileModeFileModeOpenFiles FileDialogFileMode = 1
  FileDialogFileModeFileModeOpenDir FileDialogFileMode = 2
  FileDialogFileModeFileModeOpenAny FileDialogFileMode = 3
  FileDialogFileModeFileModeSaveFile FileDialogFileMode = 4
)

type FileDialogAccess int
const (
  FileDialogAccessAccessResources FileDialogAccess = 0
  FileDialogAccessAccessUserdata FileDialogAccess = 1
  FileDialogAccessAccessFilesystem FileDialogAccess = 2
)

func  (me *FileDialog) ClearFilters() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) AddFilter(filter String, description String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetFilters(filters PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetFilters() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetCurrentDir() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetCurrentFile() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetCurrentPath() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetCurrentDir(dir String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetCurrentFile(file String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetCurrentPath(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetModeOverridesTitle(override bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) IsModeOverridingTitle() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetFileMode(mode FileDialogFileMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetFileMode() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetVbox() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetLineEdit() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetAccess(access FileDialogAccess, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetAccess() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetRootSubfolder(dir String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) GetRootSubfolder() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) SetShowHiddenFiles(show bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) IsShowingHiddenFiles() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) DeselectAll() { // TODO: return value
  // TODO: implement
}

func  (me *FileDialog) Invalidate() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
