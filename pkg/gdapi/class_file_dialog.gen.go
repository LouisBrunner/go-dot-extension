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



// Enums

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

func (me *FileDialog) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FileDialog) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *FileDialog) ClearFilters()  {
  panic("TODO: implement")
}

func  (me *FileDialog) AddFilter(filter String, description String, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetFilters(filters PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetFilters()  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetCurrentDir()  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetCurrentFile()  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetCurrentPath()  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetCurrentDir(dir String, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetCurrentFile(file String, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetCurrentPath(path String, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetModeOverridesTitle(override bool, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) IsModeOverridingTitle()  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetFileMode(mode FileDialogFileMode, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetFileMode()  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetVbox()  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetLineEdit()  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetAccess(access FileDialogAccess, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetAccess()  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetRootSubfolder(dir String, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) GetRootSubfolder()  {
  panic("TODO: implement")
}

func  (me *FileDialog) SetShowHiddenFiles(show bool, )  {
  panic("TODO: implement")
}

func  (me *FileDialog) IsShowingHiddenFiles()  {
  panic("TODO: implement")
}

func  (me *FileDialog) DeselectAll()  {
  panic("TODO: implement")
}

func  (me *FileDialog) Invalidate()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
