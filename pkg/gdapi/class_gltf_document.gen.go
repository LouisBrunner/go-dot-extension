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

type ptrsForGLTFDocumentList struct {
	fnAppendFromFile                  gdc.MethodBindPtr
	fnAppendFromBuffer                gdc.MethodBindPtr
	fnAppendFromScene                 gdc.MethodBindPtr
	fnGenerateScene                   gdc.MethodBindPtr
	fnGenerateBuffer                  gdc.MethodBindPtr
	fnWriteToFilesystem               gdc.MethodBindPtr
	fnSetImageFormat                  gdc.MethodBindPtr
	fnGetImageFormat                  gdc.MethodBindPtr
	fnSetLossyQuality                 gdc.MethodBindPtr
	fnGetLossyQuality                 gdc.MethodBindPtr
	fnSetRootNodeMode                 gdc.MethodBindPtr
	fnGetRootNodeMode                 gdc.MethodBindPtr
	fnRegisterGltfDocumentExtension   gdc.MethodBindPtr
	fnUnregisterGltfDocumentExtension gdc.MethodBindPtr
}

var ptrsForGLTFDocument ptrsForGLTFDocumentList

func initGLTFDocumentPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFDocument")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("append_from_file")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnAppendFromFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866380864))
	}
	{
		methodName := StringNameFromStr("append_from_buffer")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnAppendFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1616081266))
	}
	{
		methodName := StringNameFromStr("append_from_scene")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnAppendFromScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1622574258))
	}
	{
		methodName := StringNameFromStr("generate_scene")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnGenerateScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 596118388))
	}
	{
		methodName := StringNameFromStr("generate_buffer")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnGenerateBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 741783455))
	}
	{
		methodName := StringNameFromStr("write_to_filesystem")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnWriteToFilesystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1784551478))
	}
	{
		methodName := StringNameFromStr("set_image_format")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnSetImageFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_image_format")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnGetImageFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_lossy_quality")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnSetLossyQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_lossy_quality")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnGetLossyQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_root_node_mode")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnSetRootNodeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 463633402))
	}
	{
		methodName := StringNameFromStr("get_root_node_mode")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnGetRootNodeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 948057992))
	}
	{
		methodName := StringNameFromStr("register_gltf_document_extension")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnRegisterGltfDocumentExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3752678331))
	}
	{
		methodName := StringNameFromStr("unregister_gltf_document_extension")
		defer methodName.Destroy()
		ptrsForGLTFDocument.fnUnregisterGltfDocumentExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684415758))
	}
}

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
	GLTFDocumentRootNodeModeRootNodeModeKeepRoot   GLTFDocumentRootNodeMode = 1
	GLTFDocumentRootNodeModeRootNodeModeMultiRoot  GLTFDocumentRootNodeMode = 2
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

func (me *GLTFDocument) AppendFromFile(path String, state GLTFState, flags int64, base_path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), state.AsCTypePtr(), gdc.ConstTypePtr(&flags), base_path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnAppendFromFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GLTFDocument) AppendFromBuffer(bytes PackedByteArray, base_path String, state GLTFState, flags int64) Error {
	cargs := []gdc.ConstTypePtr{bytes.AsCTypePtr(), base_path.AsCTypePtr(), state.AsCTypePtr(), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnAppendFromBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GLTFDocument) AppendFromScene(node Node, state GLTFState, flags int64) Error {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), state.AsCTypePtr(), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnAppendFromScene), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GLTFDocument) GenerateScene(state GLTFState, bake_fps float64, trimming bool, remove_immutable_tracks bool) Node {
	cargs := []gdc.ConstTypePtr{state.AsCTypePtr(), gdc.ConstTypePtr(&bake_fps), gdc.ConstTypePtr(&trimming), gdc.ConstTypePtr(&remove_immutable_tracks)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()
	pinner.Pin(&bake_fps)
	pinner.Pin(&trimming)
	pinner.Pin(&remove_immutable_tracks)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnGenerateScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFDocument) GenerateBuffer(state GLTFState) PackedByteArray {
	cargs := []gdc.ConstTypePtr{state.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnGenerateBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFDocument) WriteToFilesystem(state GLTFState, path String) Error {
	cargs := []gdc.ConstTypePtr{state.AsCTypePtr(), path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnWriteToFilesystem), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GLTFDocument) SetImageFormat(image_format String) {
	cargs := []gdc.ConstTypePtr{image_format.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnSetImageFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFDocument) GetImageFormat() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnGetImageFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFDocument) SetLossyQuality(lossy_quality float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lossy_quality)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnSetLossyQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFDocument) GetLossyQuality() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnGetLossyQuality), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFDocument) SetRootNodeMode(root_node_mode GLTFDocumentRootNodeMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&root_node_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnSetRootNodeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFDocument) GetRootNodeMode() GLTFDocumentRootNodeMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GLTFDocumentRootNodeMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnGetRootNodeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func GLTFDocumentRegisterGltfDocumentExtension(extension GLTFDocumentExtension, first_priority bool) {
	cargs := []gdc.ConstTypePtr{extension.AsCTypePtr(), gdc.ConstTypePtr(&first_priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnRegisterGltfDocumentExtension), nil, unsafe.SliceData(cargs), nil)

}

func GLTFDocumentUnregisterGltfDocumentExtension(extension GLTFDocumentExtension) {
	cargs := []gdc.ConstTypePtr{extension.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFDocument.fnUnregisterGltfDocumentExtension), nil, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
