// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *GLTFDocument) AppendFromFile(path String, state GLTFState, flags int, base_path String, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_from_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1862991421) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(state.AsCTypePtr()), gdc.ConstTypePtr(&flags), gdc.ConstTypePtr(base_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFDocument) AppendFromBuffer(bytes PackedByteArray, base_path String, state GLTFState, flags int, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2818062664) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bytes.AsCTypePtr()), gdc.ConstTypePtr(base_path.AsCTypePtr()), gdc.ConstTypePtr(state.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFDocument) AppendFromScene(node Node, state GLTFState, flags int, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_from_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 374125375) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(state.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFDocument) GenerateScene(state GLTFState, bake_fps float32, trimming bool, remove_immutable_tracks bool, ) Node {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2770277081) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(state.AsCTypePtr()), gdc.ConstTypePtr(&bake_fps), gdc.ConstTypePtr(&trimming), gdc.ConstTypePtr(&remove_immutable_tracks), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFDocument) GenerateBuffer(state GLTFState, ) PackedByteArray {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 741783455) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(state.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFDocument) WriteToFilesystem(state GLTFState, path String, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("write_to_filesystem")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1784551478) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(state.AsCTypePtr()), gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  GLTFDocumentRegisterGltfDocumentExtension(extension GLTFDocumentExtension, first_priority bool, )  {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_gltf_document_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3752678331) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extension.AsCTypePtr()), gdc.ConstTypePtr(&first_priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)
}

func  GLTFDocumentUnregisterGltfDocumentExtension(extension GLTFDocumentExtension, )  {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unregister_gltf_document_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684415758) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extension.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)
}

// Signals
