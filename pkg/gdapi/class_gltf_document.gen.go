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

func  (me *GLTFDocument) AppendFromFile(path String, state GLTFState, flags int, base_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocument) AppendFromBuffer(bytes PackedByteArray, base_path String, state GLTFState, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocument) AppendFromScene(node Node, state GLTFState, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocument) GenerateScene(state GLTFState, bake_fps float32, trimming bool, remove_immutable_tracks bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocument) GenerateBuffer(state GLTFState, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocument) WriteToFilesystem(state GLTFState, path String, ) { // TODO: return value
  // TODO: implement
}

func  GLTFDocumentRegisterGltfDocumentExtension(extension GLTFDocumentExtension, first_priority bool, ) { // TODO: return value
  // TODO: implement
}

func  GLTFDocumentUnregisterGltfDocumentExtension(extension GLTFDocumentExtension, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
