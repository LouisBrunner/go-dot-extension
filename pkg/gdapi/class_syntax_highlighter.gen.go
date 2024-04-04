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

type SyntaxHighlighter struct {
  Resource
}

func (me *SyntaxHighlighter) BaseClass() string {
  return "SyntaxHighlighter"
}

func NewSyntaxHighlighter() *SyntaxHighlighter {
  str := StringNameFromStr("SyntaxHighlighter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SyntaxHighlighter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SyntaxHighlighter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SyntaxHighlighter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SyntaxHighlighter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SyntaxHighlighter) GetLineSyntaxHighlighting(line int64, ) Dictionary {
  classNameV := StringNameFromStr("SyntaxHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_syntax_highlighting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3554694381) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SyntaxHighlighter) UpdateCache()  {
  classNameV := StringNameFromStr("SyntaxHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SyntaxHighlighter) ClearHighlightingCache()  {
  classNameV := StringNameFromStr("SyntaxHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_highlighting_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SyntaxHighlighter) GetTextEdit() TextEdit {
  classNameV := StringNameFromStr("SyntaxHighlighter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_edit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1893027089) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTextEdit()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
