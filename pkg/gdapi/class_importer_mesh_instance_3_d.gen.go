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

type ptrsForImporterMeshInstance3DList struct {
	fnSetMesh                       gdc.MethodBindPtr
	fnGetMesh                       gdc.MethodBindPtr
	fnSetSkin                       gdc.MethodBindPtr
	fnGetSkin                       gdc.MethodBindPtr
	fnSetSkeletonPath               gdc.MethodBindPtr
	fnGetSkeletonPath               gdc.MethodBindPtr
	fnSetLayerMask                  gdc.MethodBindPtr
	fnGetLayerMask                  gdc.MethodBindPtr
	fnSetCastShadowsSetting         gdc.MethodBindPtr
	fnGetCastShadowsSetting         gdc.MethodBindPtr
	fnSetVisibilityRangeEndMargin   gdc.MethodBindPtr
	fnGetVisibilityRangeEndMargin   gdc.MethodBindPtr
	fnSetVisibilityRangeEnd         gdc.MethodBindPtr
	fnGetVisibilityRangeEnd         gdc.MethodBindPtr
	fnSetVisibilityRangeBeginMargin gdc.MethodBindPtr
	fnGetVisibilityRangeBeginMargin gdc.MethodBindPtr
	fnSetVisibilityRangeBegin       gdc.MethodBindPtr
	fnGetVisibilityRangeBegin       gdc.MethodBindPtr
	fnSetVisibilityRangeFadeMode    gdc.MethodBindPtr
	fnGetVisibilityRangeFadeMode    gdc.MethodBindPtr
}

var ptrsForImporterMeshInstance3D ptrsForImporterMeshInstance3DList

func initImporterMeshInstance3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ImporterMeshInstance3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mesh")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2255166972))
	}
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3161779525))
	}
	{
		methodName := StringNameFromStr("set_skin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3971435618))
	}
	{
		methodName := StringNameFromStr("get_skin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2074563878))
	}
	{
		methodName := StringNameFromStr("set_skeleton_path")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetSkeletonPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_skeleton_path")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetSkeletonPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_layer_mask")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetLayerMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_layer_mask")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetLayerMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_cast_shadows_setting")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetCastShadowsSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 856677339))
	}
	{
		methodName := StringNameFromStr("get_cast_shadows_setting")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetCastShadowsSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3383019359))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_end_margin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetVisibilityRangeEndMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_end_margin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetVisibilityRangeEndMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_end")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetVisibilityRangeEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_end")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetVisibilityRangeEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_begin_margin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetVisibilityRangeBeginMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_begin_margin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetVisibilityRangeBeginMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_begin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetVisibilityRangeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_begin")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetVisibilityRangeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_fade_mode")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnSetVisibilityRangeFadeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1440117808))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_fade_mode")
		defer methodName.Destroy()
		ptrsForImporterMeshInstance3D.fnGetVisibilityRangeFadeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2067221882))
	}

}

type ImporterMeshInstance3D struct {
	Node3D
}

func (me *ImporterMeshInstance3D) BaseClass() string {
	return "ImporterMeshInstance3D"
}

func NewImporterMeshInstance3D() *ImporterMeshInstance3D {
	str := StringNameFromStr("ImporterMeshInstance3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ImporterMeshInstance3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ImporterMeshInstance3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ImporterMeshInstance3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ImporterMeshInstance3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ImporterMeshInstance3D) SetMesh(mesh ImporterMesh) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetMesh() ImporterMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImporterMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMeshInstance3D) SetSkin(skin Skin) {
	cargs := []gdc.ConstTypePtr{skin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetSkin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetSkin() Skin {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetSkin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMeshInstance3D) SetSkeletonPath(skeleton_path NodePath) {
	cargs := []gdc.ConstTypePtr{skeleton_path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetSkeletonPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetSkeletonPath() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetSkeletonPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMeshInstance3D) SetLayerMask(layer_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetLayerMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetLayerMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetLayerMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMeshInstance3D) SetCastShadowsSetting(shadow_casting_setting GeometryInstance3DShadowCastingSetting) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shadow_casting_setting)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetCastShadowsSetting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetCastShadowsSetting() GeometryInstance3DShadowCastingSetting {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GeometryInstance3DShadowCastingSetting

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetCastShadowsSetting), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ImporterMeshInstance3D) SetVisibilityRangeEndMargin(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetVisibilityRangeEndMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetVisibilityRangeEndMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetVisibilityRangeEndMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMeshInstance3D) SetVisibilityRangeEnd(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetVisibilityRangeEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetVisibilityRangeEnd() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetVisibilityRangeEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMeshInstance3D) SetVisibilityRangeBeginMargin(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetVisibilityRangeBeginMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetVisibilityRangeBeginMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetVisibilityRangeBeginMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMeshInstance3D) SetVisibilityRangeBegin(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetVisibilityRangeBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetVisibilityRangeBegin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetVisibilityRangeBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMeshInstance3D) SetVisibilityRangeFadeMode(mode GeometryInstance3DVisibilityRangeFadeMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnSetVisibilityRangeFadeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMeshInstance3D) GetVisibilityRangeFadeMode() GeometryInstance3DVisibilityRangeFadeMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GeometryInstance3DVisibilityRangeFadeMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMeshInstance3D.fnGetVisibilityRangeFadeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
