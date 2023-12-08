// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorResourcePreview struct {
  obj gdc.ObjectPtr
}

func (me *EditorResourcePreview) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorResourcePreview) BaseClass() string {
  return "EditorResourcePreview"
}

func  (me *EditorResourcePreview) QueueResourcePreview(path String, receiver Object, receiver_func StringName, userdata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorResourcePreview) QueueEditedResourcePreview(resource Resource, receiver Object, receiver_func StringName, userdata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorResourcePreview) AddPreviewGenerator(generator EditorResourcePreviewGenerator, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorResourcePreview) RemovePreviewGenerator(generator EditorResourcePreviewGenerator, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorResourcePreview) CheckForInvalidation(path String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
