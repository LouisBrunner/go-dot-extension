// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  FileDialogFileModeOpenFile FileDialogFileMode = 0
  FileDialogFileModeOpenFiles FileDialogFileMode = 1
  FileDialogFileModeOpenDir FileDialogFileMode = 2
  FileDialogFileModeOpenAny FileDialogFileMode = 3
  FileDialogFileModeSaveFile FileDialogFileMode = 4
)

type FileDialogAccess int
const (
  FileDialogAccessResources FileDialogAccess = 0
  FileDialogAccessUserdata FileDialogAccess = 1
  FileDialogAccessFilesystem FileDialogAccess = 2
)
