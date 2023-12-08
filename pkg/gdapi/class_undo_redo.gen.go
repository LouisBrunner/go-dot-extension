// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UndoRedo struct {
  obj gdc.ObjectPtr
}

func (me *UndoRedo) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *UndoRedo) BaseClass() string {
  return "UndoRedo"
}

type UndoRedoMergeMode int
const (
  UndoRedoMergeDisable UndoRedoMergeMode = 0
  UndoRedoMergeEnds UndoRedoMergeMode = 1
  UndoRedoMergeAll UndoRedoMergeMode = 2
)
