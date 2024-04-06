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

type ptrsForPerformanceList struct {
	fnGetMonitor                 gdc.MethodBindPtr
	fnAddCustomMonitor           gdc.MethodBindPtr
	fnRemoveCustomMonitor        gdc.MethodBindPtr
	fnHasCustomMonitor           gdc.MethodBindPtr
	fnGetCustomMonitor           gdc.MethodBindPtr
	fnGetMonitorModificationTime gdc.MethodBindPtr
	fnGetCustomMonitorNames      gdc.MethodBindPtr
}

var ptrsForPerformance ptrsForPerformanceList

func initPerformancePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Performance")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_monitor")
		defer methodName.Destroy()
		ptrsForPerformance.fnGetMonitor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1943275655))
	}
	{
		methodName := StringNameFromStr("add_custom_monitor")
		defer methodName.Destroy()
		ptrsForPerformance.fnAddCustomMonitor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4099036814))
	}
	{
		methodName := StringNameFromStr("remove_custom_monitor")
		defer methodName.Destroy()
		ptrsForPerformance.fnRemoveCustomMonitor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("has_custom_monitor")
		defer methodName.Destroy()
		ptrsForPerformance.fnHasCustomMonitor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2041966384))
	}
	{
		methodName := StringNameFromStr("get_custom_monitor")
		defer methodName.Destroy()
		ptrsForPerformance.fnGetCustomMonitor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2138907829))
	}
	{
		methodName := StringNameFromStr("get_monitor_modification_time")
		defer methodName.Destroy()
		ptrsForPerformance.fnGetMonitorModificationTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_custom_monitor_names")
		defer methodName.Destroy()
		ptrsForPerformance.fnGetCustomMonitorNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}

}

type Performance struct {
	Object
}

func (me *Performance) BaseClass() string {
	return "Performance"
}

func NewPerformance() *Performance {
	str := StringNameFromStr("Performance") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Performance{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type PerformanceMonitor int

const (
	PerformanceMonitorTimeFps                       PerformanceMonitor = 0
	PerformanceMonitorTimeProcess                   PerformanceMonitor = 1
	PerformanceMonitorTimePhysicsProcess            PerformanceMonitor = 2
	PerformanceMonitorTimeNavigationProcess         PerformanceMonitor = 3
	PerformanceMonitorMemoryStatic                  PerformanceMonitor = 4
	PerformanceMonitorMemoryStaticMax               PerformanceMonitor = 5
	PerformanceMonitorMemoryMessageBufferMax        PerformanceMonitor = 6
	PerformanceMonitorObjectCount                   PerformanceMonitor = 7
	PerformanceMonitorObjectResourceCount           PerformanceMonitor = 8
	PerformanceMonitorObjectNodeCount               PerformanceMonitor = 9
	PerformanceMonitorObjectOrphanNodeCount         PerformanceMonitor = 10
	PerformanceMonitorRenderTotalObjectsInFrame     PerformanceMonitor = 11
	PerformanceMonitorRenderTotalPrimitivesInFrame  PerformanceMonitor = 12
	PerformanceMonitorRenderTotalDrawCallsInFrame   PerformanceMonitor = 13
	PerformanceMonitorRenderVideoMemUsed            PerformanceMonitor = 14
	PerformanceMonitorRenderTextureMemUsed          PerformanceMonitor = 15
	PerformanceMonitorRenderBufferMemUsed           PerformanceMonitor = 16
	PerformanceMonitorPhysics2DActiveObjects        PerformanceMonitor = 17
	PerformanceMonitorPhysics2DCollisionPairs       PerformanceMonitor = 18
	PerformanceMonitorPhysics2DIslandCount          PerformanceMonitor = 19
	PerformanceMonitorPhysics3DActiveObjects        PerformanceMonitor = 20
	PerformanceMonitorPhysics3DCollisionPairs       PerformanceMonitor = 21
	PerformanceMonitorPhysics3DIslandCount          PerformanceMonitor = 22
	PerformanceMonitorAudioOutputLatency            PerformanceMonitor = 23
	PerformanceMonitorNavigationActiveMaps          PerformanceMonitor = 24
	PerformanceMonitorNavigationRegionCount         PerformanceMonitor = 25
	PerformanceMonitorNavigationAgentCount          PerformanceMonitor = 26
	PerformanceMonitorNavigationLinkCount           PerformanceMonitor = 27
	PerformanceMonitorNavigationPolygonCount        PerformanceMonitor = 28
	PerformanceMonitorNavigationEdgeCount           PerformanceMonitor = 29
	PerformanceMonitorNavigationEdgeMergeCount      PerformanceMonitor = 30
	PerformanceMonitorNavigationEdgeConnectionCount PerformanceMonitor = 31
	PerformanceMonitorNavigationEdgeFreeCount       PerformanceMonitor = 32
	PerformanceMonitorMonitorMax                    PerformanceMonitor = 33
)

func (me *Performance) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Performance) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Performance) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Performance) GetMonitor(monitor PerformanceMonitor) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&monitor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&monitor)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPerformance.fnGetMonitor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Performance) AddCustomMonitor(id StringName, callable Callable, arguments Array) {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr(), callable.AsCTypePtr(), arguments.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPerformance.fnAddCustomMonitor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Performance) RemoveCustomMonitor(id StringName) {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPerformance.fnRemoveCustomMonitor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Performance) HasCustomMonitor(id StringName) bool {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPerformance.fnHasCustomMonitor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Performance) GetCustomMonitor(id StringName) Variant {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPerformance.fnGetCustomMonitor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Performance) GetMonitorModificationTime() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPerformance.fnGetMonitorModificationTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Performance) GetCustomMonitorNames() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPerformance.fnGetCustomMonitorNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Signals
