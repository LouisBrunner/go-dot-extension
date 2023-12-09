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



// Enums

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

func (me *EditorFileDialog) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorFileDialog) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorFileDialog) ClearFilters()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) AddFilter(filter String, description String, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetFilters(filters PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetFilters()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetCurrentDir()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetCurrentFile()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetCurrentPath()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetCurrentDir(dir String, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetCurrentFile(file String, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetCurrentPath(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetFileMode(mode EditorFileDialogFileMode, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetFileMode()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetVbox()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetLineEdit()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetAccess(access EditorFileDialogAccess, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetAccess()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetShowHiddenFiles(show bool, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) IsShowingHiddenFiles()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetDisplayMode(mode EditorFileDialogDisplayMode, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) GetDisplayMode()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) SetDisableOverwriteWarning(disable bool, )  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) IsOverwriteWarningDisabled()  {
  panic("TODO: implement")
}

func  (me *EditorFileDialog) Invalidate()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
