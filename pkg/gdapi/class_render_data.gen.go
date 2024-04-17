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

type ptrsForRenderDataList struct {
	fnGetRenderSceneBuffers gdc.MethodBindPtr
	fnGetRenderSceneData    gdc.MethodBindPtr
	fnGetEnvironment        gdc.MethodBindPtr
	fnGetCameraAttributes   gdc.MethodBindPtr
}

var ptrsForRenderData ptrsForRenderDataList

func initRenderDataPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderData")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_render_scene_buffers")
		defer methodName.Destroy()
		ptrsForRenderData.fnGetRenderSceneBuffers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2793216201))
	}
	{
		methodName := StringNameFromStr("get_render_scene_data")
		defer methodName.Destroy()
		ptrsForRenderData.fnGetRenderSceneData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1288715698))
	}
	{
		methodName := StringNameFromStr("get_environment")
		defer methodName.Destroy()
		ptrsForRenderData.fnGetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_camera_attributes")
		defer methodName.Destroy()
		ptrsForRenderData.fnGetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}

}

type RenderData struct {
	Object
}

func (me *RenderData) BaseClass() string {
	return "RenderData"
}

func NewRenderData() *RenderData {
	str := StringNameFromStr("RenderData") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderData{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderData) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderData) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderData) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RenderData) GetRenderSceneBuffers() RenderSceneBuffers {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRenderSceneBuffers()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderData.fnGetRenderSceneBuffers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderData) GetRenderSceneData() RenderSceneData {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRenderSceneData()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderData.fnGetRenderSceneData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderData) GetEnvironment() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderData.fnGetEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderData) GetCameraAttributes() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderData.fnGetCameraAttributes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
