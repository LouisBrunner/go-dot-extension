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



// Enums

func (me *EditorResourcePreview) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorResourcePreview) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorResourcePreview) QueueResourcePreview(path String, receiver Object, receiver_func StringName, userdata Variant, )  {
  panic("TODO: implement")
}

func  (me *EditorResourcePreview) QueueEditedResourcePreview(resource Resource, receiver Object, receiver_func StringName, userdata Variant, )  {
  panic("TODO: implement")
}

func  (me *EditorResourcePreview) AddPreviewGenerator(generator EditorResourcePreviewGenerator, )  {
  panic("TODO: implement")
}

func  (me *EditorResourcePreview) RemovePreviewGenerator(generator EditorResourcePreviewGenerator, )  {
  panic("TODO: implement")
}

func  (me *EditorResourcePreview) CheckForInvalidation(path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
