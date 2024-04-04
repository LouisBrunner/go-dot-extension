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

type GLTFDocument struct {
  Resource
}

func (me *GLTFDocument) BaseClass() string {
  return "GLTFDocument"
}

func NewGLTFDocument() *GLTFDocument {
  str := StringNameFromStr("GLTFDocument") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFDocument{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type GLTFDocumentRootNodeMode int
const (
  GLTFDocumentRootNodeModeRootNodeModeSingleRoot GLTFDocumentRootNodeMode = 0
  GLTFDocumentRootNodeModeRootNodeModeKeepRoot GLTFDocumentRootNodeMode = 1
  GLTFDocumentRootNodeModeRootNodeModeMultiRoot GLTFDocumentRootNodeMode = 2
)

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

func  (me *GLTFDocument) AppendFromFile(path String, state GLTFState, flags int64, base_path String, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_from_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866380864) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), state.AsCTypePtr(), gdc.ConstTypePtr(&flags) , base_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&flags)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GLTFDocument) AppendFromBuffer(bytes PackedByteArray, base_path String, state GLTFState, flags int64, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1616081266) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{bytes.AsCTypePtr(), base_path.AsCTypePtr(), state.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&flags)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GLTFDocument) AppendFromScene(node Node, state GLTFState, flags int64, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_from_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1622574258) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), state.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&flags)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GLTFDocument) GenerateScene(state GLTFState, bake_fps float64, trimming bool, remove_immutable_tracks bool, ) Node {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 596118388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{state.AsCTypePtr(), gdc.ConstTypePtr(&bake_fps) , gdc.ConstTypePtr(&trimming) , gdc.ConstTypePtr(&remove_immutable_tracks) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()
  pinner.Pin(&bake_fps)
  pinner.Pin(&trimming)
  pinner.Pin(&remove_immutable_tracks)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFDocument) GenerateBuffer(state GLTFState, ) PackedByteArray {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 741783455) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{state.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFDocument) WriteToFilesystem(state GLTFState, path String, ) Error {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("write_to_filesystem")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1784551478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{state.AsCTypePtr(), path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GLTFDocument) SetImageFormat(image_format String, )  {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_image_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{image_format.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFDocument) GetImageFormat() String {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_image_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFDocument) SetLossyQuality(lossy_quality float64, )  {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lossy_quality")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lossy_quality) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFDocument) GetLossyQuality() float64 {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lossy_quality")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFDocument) SetRootNodeMode(root_node_mode GLTFDocumentRootNodeMode, )  {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_node_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 463633402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&root_node_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFDocument) GetRootNodeMode() GLTFDocumentRootNodeMode {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_node_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 948057992) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GLTFDocumentRootNodeMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  GLTFDocumentRegisterGltfDocumentExtension(extension GLTFDocumentExtension, first_priority bool, )  {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_gltf_document_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3752678331) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{extension.AsCTypePtr(), gdc.ConstTypePtr(&first_priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)

}

func  GLTFDocumentUnregisterGltfDocumentExtension(extension GLTFDocumentExtension, )  {
  classNameV := StringNameFromStr("GLTFDocument")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unregister_gltf_document_extension")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684415758) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{extension.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
