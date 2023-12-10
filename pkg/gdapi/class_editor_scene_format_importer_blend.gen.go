// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSceneFormatImporterBlend struct {
  obj gdc.ObjectPtr
}

func (me *EditorSceneFormatImporterBlend) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorSceneFormatImporterBlend) BaseClass() string {
  return "EditorSceneFormatImporterBlend"
}



// Enums

func (me *EditorSceneFormatImporterBlend) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterBlend) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterBlend) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties