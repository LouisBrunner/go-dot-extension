// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorResourcePreviewGenerator struct {
  obj gdc.ObjectPtr
}

func (me *EditorResourcePreviewGenerator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorResourcePreviewGenerator) BaseClass() string {
  return "EditorResourcePreviewGenerator"
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

// Properties