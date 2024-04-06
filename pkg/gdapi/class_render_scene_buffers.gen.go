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

type ptrsForRenderSceneBuffersList struct {
	fnConfigure gdc.MethodBindPtr
}

var ptrsForRenderSceneBuffers ptrsForRenderSceneBuffersList

func initRenderSceneBuffersPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderSceneBuffers")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("configure")
		defer methodName.Destroy()
		ptrsForRenderSceneBuffers.fnConfigure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3072623270))
	}
}

type RenderSceneBuffers struct {
	RefCounted
}

func (me *RenderSceneBuffers) BaseClass() string {
	return "RenderSceneBuffers"
}

func NewRenderSceneBuffers() *RenderSceneBuffers {
	str := StringNameFromStr("RenderSceneBuffers") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderSceneBuffers{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderSceneBuffers) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderSceneBuffers) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderSceneBuffers) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RenderSceneBuffers) Configure(config RenderSceneBuffersConfiguration) {
	cargs := []gdc.ConstTypePtr{config.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneBuffers.fnConfigure), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
