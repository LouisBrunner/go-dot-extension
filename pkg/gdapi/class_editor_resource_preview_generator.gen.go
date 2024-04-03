// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorResourcePreviewGenerator struct {
  RefCounted
}

func (me *EditorResourcePreviewGenerator) BaseClass() string {
  return "EditorResourcePreviewGenerator"
}

func NewEditorResourcePreviewGenerator() *EditorResourcePreviewGenerator {
  str := StringNameFromStr("EditorResourcePreviewGenerator") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorResourcePreviewGenerator{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorResourcePreviewGenerator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorResourcePreviewGenerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorResourcePreviewGenerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
