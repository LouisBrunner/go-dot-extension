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

type ptrsForCamera3DList struct {
	fnProjectRayNormal      gdc.MethodBindPtr
	fnProjectLocalRayNormal gdc.MethodBindPtr
	fnProjectRayOrigin      gdc.MethodBindPtr
	fnUnprojectPosition     gdc.MethodBindPtr
	fnIsPositionBehind      gdc.MethodBindPtr
	fnProjectPosition       gdc.MethodBindPtr
	fnSetPerspective        gdc.MethodBindPtr
	fnSetOrthogonal         gdc.MethodBindPtr
	fnSetFrustum            gdc.MethodBindPtr
	fnMakeCurrent           gdc.MethodBindPtr
	fnClearCurrent          gdc.MethodBindPtr
	fnSetCurrent            gdc.MethodBindPtr
	fnIsCurrent             gdc.MethodBindPtr
	fnGetCameraTransform    gdc.MethodBindPtr
	fnGetCameraProjection   gdc.MethodBindPtr
	fnGetFov                gdc.MethodBindPtr
	fnGetFrustumOffset      gdc.MethodBindPtr
	fnGetSize               gdc.MethodBindPtr
	fnGetFar                gdc.MethodBindPtr
	fnGetNear               gdc.MethodBindPtr
	fnSetFov                gdc.MethodBindPtr
	fnSetFrustumOffset      gdc.MethodBindPtr
	fnSetSize               gdc.MethodBindPtr
	fnSetFar                gdc.MethodBindPtr
	fnSetNear               gdc.MethodBindPtr
	fnGetProjection         gdc.MethodBindPtr
	fnSetProjection         gdc.MethodBindPtr
	fnSetHOffset            gdc.MethodBindPtr
	fnGetHOffset            gdc.MethodBindPtr
	fnSetVOffset            gdc.MethodBindPtr
	fnGetVOffset            gdc.MethodBindPtr
	fnSetCullMask           gdc.MethodBindPtr
	fnGetCullMask           gdc.MethodBindPtr
	fnSetEnvironment        gdc.MethodBindPtr
	fnGetEnvironment        gdc.MethodBindPtr
	fnSetAttributes         gdc.MethodBindPtr
	fnGetAttributes         gdc.MethodBindPtr
	fnSetKeepAspectMode     gdc.MethodBindPtr
	fnGetKeepAspectMode     gdc.MethodBindPtr
	fnSetDopplerTracking    gdc.MethodBindPtr
	fnGetDopplerTracking    gdc.MethodBindPtr
	fnGetFrustum            gdc.MethodBindPtr
	fnIsPositionInFrustum   gdc.MethodBindPtr
	fnGetCameraRid          gdc.MethodBindPtr
	fnGetPyramidShapeRid    gdc.MethodBindPtr
	fnSetCullMaskValue      gdc.MethodBindPtr
	fnGetCullMaskValue      gdc.MethodBindPtr
}

var ptrsForCamera3D ptrsForCamera3DList

func initCamera3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Camera3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("project_ray_normal")
		defer methodName.Destroy()
		ptrsForCamera3D.fnProjectRayNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1718073306))
	}
	{
		methodName := StringNameFromStr("project_local_ray_normal")
		defer methodName.Destroy()
		ptrsForCamera3D.fnProjectLocalRayNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1718073306))
	}
	{
		methodName := StringNameFromStr("project_ray_origin")
		defer methodName.Destroy()
		ptrsForCamera3D.fnProjectRayOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1718073306))
	}
	{
		methodName := StringNameFromStr("unproject_position")
		defer methodName.Destroy()
		ptrsForCamera3D.fnUnprojectPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3758901831))
	}
	{
		methodName := StringNameFromStr("is_position_behind")
		defer methodName.Destroy()
		ptrsForCamera3D.fnIsPositionBehind = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3108956480))
	}
	{
		methodName := StringNameFromStr("project_position")
		defer methodName.Destroy()
		ptrsForCamera3D.fnProjectPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2171975744))
	}
	{
		methodName := StringNameFromStr("set_perspective")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetPerspective = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2385087082))
	}
	{
		methodName := StringNameFromStr("set_orthogonal")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetOrthogonal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2385087082))
	}
	{
		methodName := StringNameFromStr("set_frustum")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetFrustum = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 354890663))
	}
	{
		methodName := StringNameFromStr("make_current")
		defer methodName.Destroy()
		ptrsForCamera3D.fnMakeCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_current")
		defer methodName.Destroy()
		ptrsForCamera3D.fnClearCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3216645846))
	}
	{
		methodName := StringNameFromStr("set_current")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_current")
		defer methodName.Destroy()
		ptrsForCamera3D.fnIsCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_camera_transform")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetCameraTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("get_camera_projection")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetCameraProjection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2910717950))
	}
	{
		methodName := StringNameFromStr("get_fov")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_frustum_offset")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetFrustumOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_far")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetFar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_near")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetNear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fov")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_frustum_offset")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetFrustumOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_far")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetFar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_near")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetNear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_projection")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetProjection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2624185235))
	}
	{
		methodName := StringNameFromStr("set_projection")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetProjection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4218540108))
	}
	{
		methodName := StringNameFromStr("set_h_offset")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetHOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_h_offset")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetHOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_v_offset")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetVOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_v_offset")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetVOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_cull_mask")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_cull_mask")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_environment")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4143518816))
	}
	{
		methodName := StringNameFromStr("get_environment")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3082064660))
	}
	{
		methodName := StringNameFromStr("set_attributes")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2817810567))
	}
	{
		methodName := StringNameFromStr("get_attributes")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3921283215))
	}
	{
		methodName := StringNameFromStr("set_keep_aspect_mode")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetKeepAspectMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740651252))
	}
	{
		methodName := StringNameFromStr("get_keep_aspect_mode")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetKeepAspectMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2790278316))
	}
	{
		methodName := StringNameFromStr("set_doppler_tracking")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetDopplerTracking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3109431270))
	}
	{
		methodName := StringNameFromStr("get_doppler_tracking")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetDopplerTracking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1584483649))
	}
	{
		methodName := StringNameFromStr("get_frustum")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetFrustum = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("is_position_in_frustum")
		defer methodName.Destroy()
		ptrsForCamera3D.fnIsPositionInFrustum = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3108956480))
	}
	{
		methodName := StringNameFromStr("get_camera_rid")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetCameraRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_pyramid_shape_rid")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetPyramidShapeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("set_cull_mask_value")
		defer methodName.Destroy()
		ptrsForCamera3D.fnSetCullMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_cull_mask_value")
		defer methodName.Destroy()
		ptrsForCamera3D.fnGetCullMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
}

type Camera3D struct {
	Node3D
}

func (me *Camera3D) BaseClass() string {
	return "Camera3D"
}

func NewCamera3D() *Camera3D {
	str := StringNameFromStr("Camera3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Camera3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type Camera3DProjectionType int

const (
	Camera3DProjectionTypeProjectionPerspective Camera3DProjectionType = 0
	Camera3DProjectionTypeProjectionOrthogonal  Camera3DProjectionType = 1
	Camera3DProjectionTypeProjectionFrustum     Camera3DProjectionType = 2
)

type Camera3DKeepAspect int

const (
	Camera3DKeepAspectKeepWidth  Camera3DKeepAspect = 0
	Camera3DKeepAspectKeepHeight Camera3DKeepAspect = 1
)

type Camera3DDopplerTracking int

const (
	Camera3DDopplerTrackingDopplerTrackingDisabled    Camera3DDopplerTracking = 0
	Camera3DDopplerTrackingDopplerTrackingIdleStep    Camera3DDopplerTracking = 1
	Camera3DDopplerTrackingDopplerTrackingPhysicsStep Camera3DDopplerTracking = 2
)

func (me *Camera3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Camera3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Camera3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Camera3D) ProjectRayNormal(screen_point Vector2) Vector3 {
	cargs := []gdc.ConstTypePtr{screen_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnProjectRayNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) ProjectLocalRayNormal(screen_point Vector2) Vector3 {
	cargs := []gdc.ConstTypePtr{screen_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnProjectLocalRayNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) ProjectRayOrigin(screen_point Vector2) Vector3 {
	cargs := []gdc.ConstTypePtr{screen_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnProjectRayOrigin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) UnprojectPosition(world_point Vector3) Vector2 {
	cargs := []gdc.ConstTypePtr{world_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnUnprojectPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) IsPositionBehind(world_point Vector3) bool {
	cargs := []gdc.ConstTypePtr{world_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnIsPositionBehind), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) ProjectPosition(screen_point Vector2, z_depth float64) Vector3 {
	cargs := []gdc.ConstTypePtr{screen_point.AsCTypePtr(), gdc.ConstTypePtr(&z_depth)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&z_depth)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnProjectPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) SetPerspective(fov float64, z_near float64, z_far float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fov), gdc.ConstTypePtr(&z_near), gdc.ConstTypePtr(&z_far)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetPerspective), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetOrthogonal(size float64, z_near float64, z_far float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&z_near), gdc.ConstTypePtr(&z_far)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetOrthogonal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetFrustum(size float64, offset Vector2, z_near float64, z_far float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), offset.AsCTypePtr(), gdc.ConstTypePtr(&z_near), gdc.ConstTypePtr(&z_far)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetFrustum), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) MakeCurrent() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnMakeCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) ClearCurrent(enable_next bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable_next)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnClearCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetCurrent(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) IsCurrent() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnIsCurrent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) GetCameraTransform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetCameraTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) GetCameraProjection() Projection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewProjection()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetCameraProjection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) GetFov() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetFov), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) GetFrustumOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetFrustumOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) GetSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) GetFar() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetFar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) GetNear() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetNear), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) SetFov(fov float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fov)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetFov), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetFrustumOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetFrustumOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetSize(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetFar(far float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&far)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetFar), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetNear(near float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&near)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetNear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetProjection() Camera3DProjectionType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Camera3DProjectionType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetProjection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Camera3D) SetProjection(mode Camera3DProjectionType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetProjection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) SetHOffset(offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetHOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetHOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetHOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) SetVOffset(offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetVOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetVOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetVOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) SetCullMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetCullMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) SetEnvironment(env Environment) {
	cargs := []gdc.ConstTypePtr{env.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetEnvironment() Environment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewEnvironment()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) SetAttributes(env CameraAttributes) {
	cargs := []gdc.ConstTypePtr{env.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetAttributes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetAttributes() CameraAttributes {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCameraAttributes()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetAttributes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) SetKeepAspectMode(mode Camera3DKeepAspect) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetKeepAspectMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetKeepAspectMode() Camera3DKeepAspect {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Camera3DKeepAspect

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetKeepAspectMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Camera3D) SetDopplerTracking(mode Camera3DDopplerTracking) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetDopplerTracking), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetDopplerTracking() Camera3DDopplerTracking {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Camera3DDopplerTracking

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetDopplerTracking), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Camera3D) GetFrustum() []Plane {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetFrustum), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Plane](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Camera3D) IsPositionInFrustum(world_point Vector3) bool {
	cargs := []gdc.ConstTypePtr{world_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnIsPositionInFrustum), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Camera3D) GetCameraRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetCameraRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) GetPyramidShapeRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetPyramidShapeRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Camera3D) SetCullMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnSetCullMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Camera3D) GetCullMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCamera3D.fnGetCullMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
