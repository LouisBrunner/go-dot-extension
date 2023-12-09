// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFDocument struct {
  obj gdc.ObjectPtr
}

func (me *GLTFDocument) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFDocument) BaseClass() string {
  return "GLTFDocument"
}



// Enums

func (me *GLTFDocument) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFDocument) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFDocument) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GLTFDocument) AppendFromFile(path String, state GLTFState, flags int, base_path String, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocument) AppendFromBuffer(bytes PackedByteArray, base_path String, state GLTFState, flags int, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocument) AppendFromScene(node Node, state GLTFState, flags int, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocument) GenerateScene(state GLTFState, bake_fps float32, trimming bool, remove_immutable_tracks bool, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocument) GenerateBuffer(state GLTFState, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocument) WriteToFilesystem(state GLTFState, path String, )  {
  panic("TODO: implement")
}

func  GLTFDocumentRegisterGltfDocumentExtension(extension GLTFDocumentExtension, first_priority bool, )  {
  panic("TODO: implement")
}

func  GLTFDocumentUnregisterGltfDocumentExtension(extension GLTFDocumentExtension, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
