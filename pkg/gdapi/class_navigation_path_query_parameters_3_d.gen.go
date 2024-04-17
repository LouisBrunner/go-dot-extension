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

type ptrsForNavigationPathQueryParameters3DList struct {
	fnSetPathfindingAlgorithm gdc.MethodBindPtr
	fnGetPathfindingAlgorithm gdc.MethodBindPtr
	fnSetPathPostprocessing   gdc.MethodBindPtr
	fnGetPathPostprocessing   gdc.MethodBindPtr
	fnSetMap                  gdc.MethodBindPtr
	fnGetMap                  gdc.MethodBindPtr
	fnSetStartPosition        gdc.MethodBindPtr
	fnGetStartPosition        gdc.MethodBindPtr
	fnSetTargetPosition       gdc.MethodBindPtr
	fnGetTargetPosition       gdc.MethodBindPtr
	fnSetNavigationLayers     gdc.MethodBindPtr
	fnGetNavigationLayers     gdc.MethodBindPtr
	fnSetMetadataFlags        gdc.MethodBindPtr
	fnGetMetadataFlags        gdc.MethodBindPtr
	fnSetSimplifyPath         gdc.MethodBindPtr
	fnGetSimplifyPath         gdc.MethodBindPtr
	fnSetSimplifyEpsilon      gdc.MethodBindPtr
	fnGetSimplifyEpsilon      gdc.MethodBindPtr
}

var ptrsForNavigationPathQueryParameters3D ptrsForNavigationPathQueryParameters3DList

func initNavigationPathQueryParameters3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationPathQueryParameters3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pathfinding_algorithm")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 394560454))
	}
	{
		methodName := StringNameFromStr("get_pathfinding_algorithm")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3398491350))
	}
	{
		methodName := StringNameFromStr("set_path_postprocessing")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2267362344))
	}
	{
		methodName := StringNameFromStr("get_path_postprocessing")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3883858360))
	}
	{
		methodName := StringNameFromStr("set_map")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_map")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_start_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_start_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_target_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_target_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_metadata_flags")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2713846708))
	}
	{
		methodName := StringNameFromStr("get_metadata_flags")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1582332802))
	}
	{
		methodName := StringNameFromStr("set_simplify_path")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetSimplifyPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_simplify_path")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetSimplifyPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_simplify_epsilon")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnSetSimplifyEpsilon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_simplify_epsilon")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters3D.fnGetSimplifyEpsilon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type NavigationPathQueryParameters3D struct {
	RefCounted
}

func (me *NavigationPathQueryParameters3D) BaseClass() string {
	return "NavigationPathQueryParameters3D"
}

func NewNavigationPathQueryParameters3D() *NavigationPathQueryParameters3D {
	str := StringNameFromStr("NavigationPathQueryParameters3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationPathQueryParameters3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type NavigationPathQueryParameters3DPathfindingAlgorithm int

const (
	NavigationPathQueryParameters3DPathfindingAlgorithmPathfindingAlgorithmAstar NavigationPathQueryParameters3DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters3DPathPostProcessing int

const (
	NavigationPathQueryParameters3DPathPostProcessingPathPostprocessingCorridorfunnel NavigationPathQueryParameters3DPathPostProcessing = 0
	NavigationPathQueryParameters3DPathPostProcessingPathPostprocessingEdgecentered   NavigationPathQueryParameters3DPathPostProcessing = 1
)

type NavigationPathQueryParameters3DPathMetadataFlags int

const (
	NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeNone   NavigationPathQueryParameters3DPathMetadataFlags = 0
	NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeTypes  NavigationPathQueryParameters3DPathMetadataFlags = 1
	NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeRids   NavigationPathQueryParameters3DPathMetadataFlags = 2
	NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeOwners NavigationPathQueryParameters3DPathMetadataFlags = 4
	NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeAll    NavigationPathQueryParameters3DPathMetadataFlags = 7
)

func (me *NavigationPathQueryParameters3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationPathQueryParameters3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationPathQueryParameters3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetPathfindingAlgorithm() NavigationPathQueryParameters3DPathfindingAlgorithm {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters3DPathfindingAlgorithm

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationPathQueryParameters3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetPathPostprocessing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetPathPostprocessing() NavigationPathQueryParameters3DPathPostProcessing {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters3DPathPostProcessing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetPathPostprocessing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationPathQueryParameters3D) SetMap(map_ RID) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryParameters3D) SetStartPosition(start_position Vector3) {
	cargs := []gdc.ConstTypePtr{start_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetStartPosition() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryParameters3D) SetTargetPosition(target_position Vector3) {
	cargs := []gdc.ConstTypePtr{target_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetTargetPosition() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryParameters3D) SetNavigationLayers(navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetNavigationLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationPathQueryParameters3D) SetMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetMetadataFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetMetadataFlags() NavigationPathQueryParameters3DPathMetadataFlags {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters3DPathMetadataFlags

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetMetadataFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationPathQueryParameters3D) SetSimplifyPath(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetSimplifyPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetSimplifyPath() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetSimplifyPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationPathQueryParameters3D) SetSimplifyEpsilon(epsilon float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&epsilon)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnSetSimplifyEpsilon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters3D) GetSimplifyEpsilon() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters3D.fnGetSimplifyEpsilon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
