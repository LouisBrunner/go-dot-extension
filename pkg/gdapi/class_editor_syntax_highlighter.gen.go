// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSyntaxHighlighter struct {
  SyntaxHighlighter
}

func (me *EditorSyntaxHighlighter) BaseClass() string {
  return "EditorSyntaxHighlighter"
}

func NewEditorSyntaxHighlighter() *EditorSyntaxHighlighter {
  str := StringNameFromStr("EditorSyntaxHighlighter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorSyntaxHighlighter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorSyntaxHighlighter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSyntaxHighlighter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSyntaxHighlighter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
