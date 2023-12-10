// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSceneFormatImporterGLTF struct {
  obj gdc.ObjectPtr
}

func (me *EditorSceneFormatImporterGLTF) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorSceneFormatImporterGLTF) BaseClass() string {
  return "EditorSceneFormatImporterGLTF"
}



// Enums

func (me *EditorSceneFormatImporterGLTF) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterGLTF) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterGLTF) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties