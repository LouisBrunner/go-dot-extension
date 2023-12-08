// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorVCSInterface struct {
  obj gdc.ObjectPtr
}

func (me *EditorVCSInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorVCSInterface) BaseClass() string {
  return "EditorVCSInterface"
}

type EditorVCSInterfaceChangeType int
const (
  EditorVCSInterfaceChangeTypeNew EditorVCSInterfaceChangeType = 0
  EditorVCSInterfaceChangeTypeModified EditorVCSInterfaceChangeType = 1
  EditorVCSInterfaceChangeTypeRenamed EditorVCSInterfaceChangeType = 2
  EditorVCSInterfaceChangeTypeDeleted EditorVCSInterfaceChangeType = 3
  EditorVCSInterfaceChangeTypeTypechange EditorVCSInterfaceChangeType = 4
  EditorVCSInterfaceChangeTypeUnmerged EditorVCSInterfaceChangeType = 5
)

type EditorVCSInterfaceTreeArea int
const (
  EditorVCSInterfaceTreeAreaCommit EditorVCSInterfaceTreeArea = 0
  EditorVCSInterfaceTreeAreaStaged EditorVCSInterfaceTreeArea = 1
  EditorVCSInterfaceTreeAreaUnstaged EditorVCSInterfaceTreeArea = 2
)
