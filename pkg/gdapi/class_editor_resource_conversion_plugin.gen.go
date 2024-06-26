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

type ptrsForEditorResourceConversionPluginList struct {
	fnXConvertsTo gdc.MethodBindPtr
	fnXHandles    gdc.MethodBindPtr
	fnXConvert    gdc.MethodBindPtr
}

var ptrsForEditorResourceConversionPlugin ptrsForEditorResourceConversionPluginList

func initEditorResourceConversionPluginPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorResourceConversionPlugin")
	defer className.Destroy()

}

type EditorResourceConversionPlugin struct {
	RefCounted
}

func (me *EditorResourceConversionPlugin) BaseClass() string {
	return "EditorResourceConversionPlugin"
}

func NewEditorResourceConversionPlugin() *EditorResourceConversionPlugin {
	str := StringNameFromStr("EditorResourceConversionPlugin") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorResourceConversionPlugin{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorResourceConversionPlugin) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorResourceConversionPlugin) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorResourceConversionPlugin) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
