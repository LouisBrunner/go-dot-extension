// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorFileSystemDirectory struct {
  obj gdc.ObjectPtr
}

func createEditorFileSystemDirectory(obj gdc.ObjectPtr) *EditorFileSystemDirectory {
  return &EditorFileSystemDirectory{
    obj: obj,
  }
}

func (me *EditorFileSystemDirectory) BaseClass() string {
  return "EditorFileSystemDirectory"
}
