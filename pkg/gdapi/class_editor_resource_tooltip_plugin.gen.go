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

type ptrsForEditorResourceTooltipPluginList struct {
	fnXHandles            gdc.MethodBindPtr
	fnXMakeTooltipForPath gdc.MethodBindPtr
	fnRequestThumbnail    gdc.MethodBindPtr
}

var ptrsForEditorResourceTooltipPlugin ptrsForEditorResourceTooltipPluginList

func initEditorResourceTooltipPluginPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorResourceTooltipPlugin")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("request_thumbnail")
		defer methodName.Destroy()
		ptrsForEditorResourceTooltipPlugin.fnRequestThumbnail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3245519720))
	}
}

type EditorResourceTooltipPlugin struct {
	RefCounted
}

func (me *EditorResourceTooltipPlugin) BaseClass() string {
	return "EditorResourceTooltipPlugin"
}

func NewEditorResourceTooltipPlugin() *EditorResourceTooltipPlugin {
	str := StringNameFromStr("EditorResourceTooltipPlugin") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorResourceTooltipPlugin{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorResourceTooltipPlugin) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorResourceTooltipPlugin) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorResourceTooltipPlugin) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorResourceTooltipPlugin) RequestThumbnail(path String, control TextureRect) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorResourceTooltipPlugin.fnRequestThumbnail), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
