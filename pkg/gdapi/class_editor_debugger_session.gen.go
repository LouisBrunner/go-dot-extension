// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorDebuggerSession struct {
  obj gdc.ObjectPtr
}

func createEditorDebuggerSession(obj gdc.ObjectPtr) *EditorDebuggerSession {
  return &EditorDebuggerSession{
    obj: obj,
  }
}

func (me *EditorDebuggerSession) BaseClass() string {
  return "EditorDebuggerSession"
}
