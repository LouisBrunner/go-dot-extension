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

type ptrsForEditorResourcePreviewList struct {
	fnQueueResourcePreview       gdc.MethodBindPtr
	fnQueueEditedResourcePreview gdc.MethodBindPtr
	fnAddPreviewGenerator        gdc.MethodBindPtr
	fnRemovePreviewGenerator     gdc.MethodBindPtr
	fnCheckForInvalidation       gdc.MethodBindPtr
}

var ptrsForEditorResourcePreview ptrsForEditorResourcePreviewList

func initEditorResourcePreviewPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorResourcePreview")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("queue_resource_preview")
		defer methodName.Destroy()
		ptrsForEditorResourcePreview.fnQueueResourcePreview = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 233177534))
	}
	{
		methodName := StringNameFromStr("queue_edited_resource_preview")
		defer methodName.Destroy()
		ptrsForEditorResourcePreview.fnQueueEditedResourcePreview = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1608376650))
	}
	{
		methodName := StringNameFromStr("add_preview_generator")
		defer methodName.Destroy()
		ptrsForEditorResourcePreview.fnAddPreviewGenerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 332288124))
	}
	{
		methodName := StringNameFromStr("remove_preview_generator")
		defer methodName.Destroy()
		ptrsForEditorResourcePreview.fnRemovePreviewGenerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 332288124))
	}
	{
		methodName := StringNameFromStr("check_for_invalidation")
		defer methodName.Destroy()
		ptrsForEditorResourcePreview.fnCheckForInvalidation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
}

type EditorResourcePreview struct {
	Node
}

func (me *EditorResourcePreview) BaseClass() string {
	return "EditorResourcePreview"
}

func NewEditorResourcePreview() *EditorResourcePreview {
	str := StringNameFromStr("EditorResourcePreview") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorResourcePreview{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorResourcePreview) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorResourcePreview) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorResourcePreview) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorResourcePreview) QueueResourcePreview(path String, receiver Object, receiver_func StringName, userdata Variant) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), receiver.AsCTypePtr(), receiver_func.AsCTypePtr(), userdata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePreview.fnQueueResourcePreview), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePreview) QueueEditedResourcePreview(resource Resource, receiver Object, receiver_func StringName, userdata Variant) {
	cargs := []gdc.ConstTypePtr{resource.AsCTypePtr(), receiver.AsCTypePtr(), receiver_func.AsCTypePtr(), userdata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePreview.fnQueueEditedResourcePreview), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePreview) AddPreviewGenerator(generator EditorResourcePreviewGenerator) {
	cargs := []gdc.ConstTypePtr{generator.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePreview.fnAddPreviewGenerator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePreview) RemovePreviewGenerator(generator EditorResourcePreviewGenerator) {
	cargs := []gdc.ConstTypePtr{generator.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePreview.fnRemovePreviewGenerator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorResourcePreview) CheckForInvalidation(path String) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourcePreview.fnCheckForInvalidation), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type EditorResourcePreviewPreviewInvalidatedSignalFn func(path String)

func (me *EditorResourcePreview) ConnectPreviewInvalidated(subs SignalSubscribers, fn EditorResourcePreviewPreviewInvalidatedSignalFn) {
	sig := StringNameFromStr("preview_invalidated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorResourcePreview) DisconnectPreviewInvalidated(subs SignalSubscribers, fn EditorResourcePreviewPreviewInvalidatedSignalFn) {
	sig := StringNameFromStr("preview_invalidated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
