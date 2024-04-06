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

type ptrsForCSGPolygon3DList struct {
	fnSetPolygon           gdc.MethodBindPtr
	fnGetPolygon           gdc.MethodBindPtr
	fnSetMode              gdc.MethodBindPtr
	fnGetMode              gdc.MethodBindPtr
	fnSetDepth             gdc.MethodBindPtr
	fnGetDepth             gdc.MethodBindPtr
	fnSetSpinDegrees       gdc.MethodBindPtr
	fnGetSpinDegrees       gdc.MethodBindPtr
	fnSetSpinSides         gdc.MethodBindPtr
	fnGetSpinSides         gdc.MethodBindPtr
	fnSetPathNode          gdc.MethodBindPtr
	fnGetPathNode          gdc.MethodBindPtr
	fnSetPathIntervalType  gdc.MethodBindPtr
	fnGetPathIntervalType  gdc.MethodBindPtr
	fnSetPathInterval      gdc.MethodBindPtr
	fnGetPathInterval      gdc.MethodBindPtr
	fnSetPathSimplifyAngle gdc.MethodBindPtr
	fnGetPathSimplifyAngle gdc.MethodBindPtr
	fnSetPathRotation      gdc.MethodBindPtr
	fnGetPathRotation      gdc.MethodBindPtr
	fnSetPathLocal         gdc.MethodBindPtr
	fnIsPathLocal          gdc.MethodBindPtr
	fnSetPathContinuousU   gdc.MethodBindPtr
	fnIsPathContinuousU    gdc.MethodBindPtr
	fnSetPathUDistance     gdc.MethodBindPtr
	fnGetPathUDistance     gdc.MethodBindPtr
	fnSetPathJoined        gdc.MethodBindPtr
	fnIsPathJoined         gdc.MethodBindPtr
	fnSetMaterial          gdc.MethodBindPtr
	fnGetMaterial          gdc.MethodBindPtr
	fnSetSmoothFaces       gdc.MethodBindPtr
	fnGetSmoothFaces       gdc.MethodBindPtr
}

var ptrsForCSGPolygon3D ptrsForCSGPolygon3DList

func initCSGPolygon3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CSGPolygon3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_polygon")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_polygon")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_mode")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3158377035))
	}
	{
		methodName := StringNameFromStr("get_mode")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1201612222))
	}
	{
		methodName := StringNameFromStr("set_depth")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_spin_degrees")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetSpinDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_spin_degrees")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetSpinDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_spin_sides")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetSpinSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_spin_sides")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetSpinSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_path_node")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_path_node")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetPathNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_path_interval_type")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathIntervalType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744240707))
	}
	{
		methodName := StringNameFromStr("get_path_interval_type")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetPathIntervalType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3434618397))
	}
	{
		methodName := StringNameFromStr("set_path_interval")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_path_interval")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetPathInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_path_simplify_angle")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathSimplifyAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_path_simplify_angle")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetPathSimplifyAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_path_rotation")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1412947288))
	}
	{
		methodName := StringNameFromStr("get_path_rotation")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetPathRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 647219346))
	}
	{
		methodName := StringNameFromStr("set_path_local")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_path_local")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnIsPathLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_path_continuous_u")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathContinuousU = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_path_continuous_u")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnIsPathContinuousU = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_path_u_distance")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathUDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_path_u_distance")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetPathUDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_path_joined")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetPathJoined = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_path_joined")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnIsPathJoined = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("set_smooth_faces")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnSetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_smooth_faces")
		defer methodName.Destroy()
		ptrsForCSGPolygon3D.fnGetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type CSGPolygon3D struct {
	CSGPrimitive3D
}

func (me *CSGPolygon3D) BaseClass() string {
	return "CSGPolygon3D"
}

func NewCSGPolygon3D() *CSGPolygon3D {
	str := StringNameFromStr("CSGPolygon3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CSGPolygon3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CSGPolygon3DMode int

const (
	CSGPolygon3DModeModeDepth CSGPolygon3DMode = 0
	CSGPolygon3DModeModeSpin  CSGPolygon3DMode = 1
	CSGPolygon3DModeModePath  CSGPolygon3DMode = 2
)

type CSGPolygon3DPathRotation int

const (
	CSGPolygon3DPathRotationPathRotationPolygon    CSGPolygon3DPathRotation = 0
	CSGPolygon3DPathRotationPathRotationPath       CSGPolygon3DPathRotation = 1
	CSGPolygon3DPathRotationPathRotationPathFollow CSGPolygon3DPathRotation = 2
)

type CSGPolygon3DPathIntervalType int

const (
	CSGPolygon3DPathIntervalTypePathIntervalDistance  CSGPolygon3DPathIntervalType = 0
	CSGPolygon3DPathIntervalTypePathIntervalSubdivide CSGPolygon3DPathIntervalType = 1
)

func (me *CSGPolygon3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CSGPolygon3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CSGPolygon3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CSGPolygon3D) SetPolygon(polygon PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetPolygon() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CSGPolygon3D) SetMode(mode CSGPolygon3DMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetMode() CSGPolygon3DMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CSGPolygon3DMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CSGPolygon3D) SetDepth(depth float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetSpinDegrees(degrees float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetSpinDegrees), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetSpinDegrees() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetSpinDegrees), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetSpinSides(spin_sides int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spin_sides)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetSpinSides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetSpinSides() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetSpinSides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetPathNode(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetPathNode() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetPathNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CSGPolygon3D) SetPathIntervalType(interval_type CSGPolygon3DPathIntervalType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathIntervalType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetPathIntervalType() CSGPolygon3DPathIntervalType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CSGPolygon3DPathIntervalType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetPathIntervalType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CSGPolygon3D) SetPathInterval(interval float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetPathInterval() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetPathInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetPathSimplifyAngle(degrees float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathSimplifyAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetPathSimplifyAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetPathSimplifyAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetPathRotation(path_rotation CSGPolygon3DPathRotation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_rotation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetPathRotation() CSGPolygon3DPathRotation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CSGPolygon3DPathRotation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetPathRotation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CSGPolygon3D) SetPathLocal(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathLocal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) IsPathLocal() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnIsPathLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetPathContinuousU(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathContinuousU), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) IsPathContinuousU() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnIsPathContinuousU), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetPathUDistance(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathUDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetPathUDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetPathUDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetPathJoined(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetPathJoined), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) IsPathJoined() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnIsPathJoined), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGPolygon3D) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CSGPolygon3D) SetSmoothFaces(smooth_faces bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth_faces)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnSetSmoothFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGPolygon3D) GetSmoothFaces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGPolygon3D.fnGetSmoothFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
