// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  EditorFileDialogFileModeOpenFile EditorFileDialogFileMode = 0
  EditorFileDialogFileModeOpenFiles EditorFileDialogFileMode = 1
  EditorFileDialogFileModeOpenDir EditorFileDialogFileMode = 2
  EditorFileDialogFileModeOpenAny EditorFileDialogFileMode = 3
  EditorFileDialogFileModeSaveFile EditorFileDialogFileMode = 4
)

type EditorFileDialogAccess int
const (
  EditorFileDialogAccessResources EditorFileDialogAccess = 0
  EditorFileDialogAccessUserdata EditorFileDialogAccess = 1
  EditorFileDialogAccessFilesystem EditorFileDialogAccess = 2
)

type EditorFileDialogDisplayMode int
const (
  EditorFileDialogDisplayThumbnails EditorFileDialogDisplayMode = 0
  EditorFileDialogDisplayList EditorFileDialogDisplayMode = 1
)
