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

type ptrsForNavigationPathQueryParameters2DList struct {
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

var ptrsForNavigationPathQueryParameters2D ptrsForNavigationPathQueryParameters2DList

func initNavigationPathQueryParameters2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationPathQueryParameters2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pathfinding_algorithm")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783519915))
	}
	{
		methodName := StringNameFromStr("get_pathfinding_algorithm")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3000421146))
	}
	{
		methodName := StringNameFromStr("set_path_postprocessing")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2864409082))
	}
	{
		methodName := StringNameFromStr("get_path_postprocessing")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3798118993))
	}
	{
		methodName := StringNameFromStr("set_map")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_map")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_start_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_start_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_target_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_target_position")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_metadata_flags")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 24274129))
	}
	{
		methodName := StringNameFromStr("get_metadata_flags")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 488152976))
	}
	{
		methodName := StringNameFromStr("set_simplify_path")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetSimplifyPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_simplify_path")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetSimplifyPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_simplify_epsilon")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnSetSimplifyEpsilon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_simplify_epsilon")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryParameters2D.fnGetSimplifyEpsilon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type NavigationPathQueryParameters2D struct {
	RefCounted
}

func (me *NavigationPathQueryParameters2D) BaseClass() string {
	return "NavigationPathQueryParameters2D"
}

func NewNavigationPathQueryParameters2D() *NavigationPathQueryParameters2D {
	str := StringNameFromStr("NavigationPathQueryParameters2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationPathQueryParameters2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type NavigationPathQueryParameters2DPathfindingAlgorithm int

const (
	NavigationPathQueryParameters2DPathfindingAlgorithmPathfindingAlgorithmAstar NavigationPathQueryParameters2DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters2DPathPostProcessing int

const (
	NavigationPathQueryParameters2DPathPostProcessingPathPostprocessingCorridorfunnel NavigationPathQueryParameters2DPathPostProcessing = 0
	NavigationPathQueryParameters2DPathPostProcessingPathPostprocessingEdgecentered   NavigationPathQueryParameters2DPathPostProcessing = 1
)

type NavigationPathQueryParameters2DPathMetadataFlags int

const (
	NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeNone   NavigationPathQueryParameters2DPathMetadataFlags = 0
	NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeTypes  NavigationPathQueryParameters2DPathMetadataFlags = 1
	NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeRids   NavigationPathQueryParameters2DPathMetadataFlags = 2
	NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeOwners NavigationPathQueryParameters2DPathMetadataFlags = 4
	NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeAll    NavigationPathQueryParameters2DPathMetadataFlags = 7
)

func (me *NavigationPathQueryParameters2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationPathQueryParameters2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationPathQueryParameters2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetPathfindingAlgorithm() NavigationPathQueryParameters2DPathfindingAlgorithm {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters2DPathfindingAlgorithm

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationPathQueryParameters2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetPathPostprocessing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetPathPostprocessing() NavigationPathQueryParameters2DPathPostProcessing {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters2DPathPostProcessing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetPathPostprocessing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationPathQueryParameters2D) SetMap(map_ RID) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryParameters2D) SetStartPosition(start_position Vector2) {
	cargs := []gdc.ConstTypePtr{start_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetStartPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryParameters2D) SetTargetPosition(target_position Vector2) {
	cargs := []gdc.ConstTypePtr{target_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetTargetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryParameters2D) SetNavigationLayers(navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetNavigationLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationPathQueryParameters2D) SetMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetMetadataFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetMetadataFlags() NavigationPathQueryParameters2DPathMetadataFlags {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters2DPathMetadataFlags

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetMetadataFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationPathQueryParameters2D) SetSimplifyPath(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetSimplifyPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetSimplifyPath() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetSimplifyPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationPathQueryParameters2D) SetSimplifyEpsilon(epsilon float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&epsilon)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnSetSimplifyEpsilon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryParameters2D) GetSimplifyEpsilon() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryParameters2D.fnGetSimplifyEpsilon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
