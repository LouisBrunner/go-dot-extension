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

type ptrsForRenderSceneDataList struct {
	fnGetCamTransform   gdc.MethodBindPtr
	fnGetCamProjection  gdc.MethodBindPtr
	fnGetViewCount      gdc.MethodBindPtr
	fnGetViewEyeOffset  gdc.MethodBindPtr
	fnGetViewProjection gdc.MethodBindPtr
	fnGetUniformBuffer  gdc.MethodBindPtr
}

var ptrsForRenderSceneData ptrsForRenderSceneDataList

func initRenderSceneDataPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderSceneData")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_cam_transform")
		defer methodName.Destroy()
		ptrsForRenderSceneData.fnGetCamTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("get_cam_projection")
		defer methodName.Destroy()
		ptrsForRenderSceneData.fnGetCamProjection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2910717950))
	}
	{
		methodName := StringNameFromStr("get_view_count")
		defer methodName.Destroy()
		ptrsForRenderSceneData.fnGetViewCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_view_eye_offset")
		defer methodName.Destroy()
		ptrsForRenderSceneData.fnGetViewEyeOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_view_projection")
		defer methodName.Destroy()
		ptrsForRenderSceneData.fnGetViewProjection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3179846605))
	}
	{
		methodName := StringNameFromStr("get_uniform_buffer")
		defer methodName.Destroy()
		ptrsForRenderSceneData.fnGetUniformBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}

}

type RenderSceneData struct {
	Object
}

func (me *RenderSceneData) BaseClass() string {
	return "RenderSceneData"
}

func NewRenderSceneData() *RenderSceneData {
	str := StringNameFromStr("RenderSceneData") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderSceneData{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RenderSceneData) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderSceneData) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderSceneData) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RenderSceneData) GetCamTransform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneData.fnGetCamTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneData) GetCamProjection() Projection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewProjection()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneData.fnGetCamProjection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneData) GetViewCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneData.fnGetViewCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderSceneData) GetViewEyeOffset(view int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&view)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneData.fnGetViewEyeOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneData) GetViewProjection(view int64) Projection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&view)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewProjection()
	pinner.Pin(&view)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneData.fnGetViewProjection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderSceneData) GetUniformBuffer() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderSceneData.fnGetUniformBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
