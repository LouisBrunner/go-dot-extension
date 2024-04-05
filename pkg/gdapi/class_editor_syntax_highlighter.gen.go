// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForEditorSyntaxHighlighterList struct {
  fnXGetName gdc.MethodBindPtr
  fnXGetSupportedLanguages gdc.MethodBindPtr
}

var ptrsForEditorSyntaxHighlighter ptrsForEditorSyntaxHighlighterList

func initEditorSyntaxHighlighterPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorSyntaxHighlighter")
  defer className.Destroy()
}

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
