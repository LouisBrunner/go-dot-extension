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

type ptrsForXRServerList struct {
	fnGetWorldScale       gdc.MethodBindPtr
	fnSetWorldScale       gdc.MethodBindPtr
	fnGetWorldOrigin      gdc.MethodBindPtr
	fnSetWorldOrigin      gdc.MethodBindPtr
	fnGetReferenceFrame   gdc.MethodBindPtr
	fnCenterOnHmd         gdc.MethodBindPtr
	fnGetHmdTransform     gdc.MethodBindPtr
	fnAddInterface        gdc.MethodBindPtr
	fnGetInterfaceCount   gdc.MethodBindPtr
	fnRemoveInterface     gdc.MethodBindPtr
	fnGetInterface        gdc.MethodBindPtr
	fnGetInterfaces       gdc.MethodBindPtr
	fnFindInterface       gdc.MethodBindPtr
	fnAddTracker          gdc.MethodBindPtr
	fnRemoveTracker       gdc.MethodBindPtr
	fnGetTrackers         gdc.MethodBindPtr
	fnGetTracker          gdc.MethodBindPtr
	fnGetPrimaryInterface gdc.MethodBindPtr
	fnSetPrimaryInterface gdc.MethodBindPtr
}

var ptrsForXRServer ptrsForXRServerList

func initXRServerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRServer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_world_scale")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetWorldScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_world_scale")
		defer methodName.Destroy()
		ptrsForXRServer.fnSetWorldScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_world_origin")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetWorldOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("set_world_origin")
		defer methodName.Destroy()
		ptrsForXRServer.fnSetWorldOrigin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_reference_frame")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetReferenceFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("center_on_hmd")
		defer methodName.Destroy()
		ptrsForXRServer.fnCenterOnHmd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1450904707))
	}
	{
		methodName := StringNameFromStr("get_hmd_transform")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetHmdTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4183770049))
	}
	{
		methodName := StringNameFromStr("add_interface")
		defer methodName.Destroy()
		ptrsForXRServer.fnAddInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1898711491))
	}
	{
		methodName := StringNameFromStr("get_interface_count")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetInterfaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("remove_interface")
		defer methodName.Destroy()
		ptrsForXRServer.fnRemoveInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1898711491))
	}
	{
		methodName := StringNameFromStr("get_interface")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4237347919))
	}
	{
		methodName := StringNameFromStr("get_interfaces")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetInterfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("find_interface")
		defer methodName.Destroy()
		ptrsForXRServer.fnFindInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1395192955))
	}
	{
		methodName := StringNameFromStr("add_tracker")
		defer methodName.Destroy()
		ptrsForXRServer.fnAddTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2692800323))
	}
	{
		methodName := StringNameFromStr("remove_tracker")
		defer methodName.Destroy()
		ptrsForXRServer.fnRemoveTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2692800323))
	}
	{
		methodName := StringNameFromStr("get_trackers")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetTrackers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3554694381))
	}
	{
		methodName := StringNameFromStr("get_tracker")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2742084544))
	}
	{
		methodName := StringNameFromStr("get_primary_interface")
		defer methodName.Destroy()
		ptrsForXRServer.fnGetPrimaryInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2143545064))
	}
	{
		methodName := StringNameFromStr("set_primary_interface")
		defer methodName.Destroy()
		ptrsForXRServer.fnSetPrimaryInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1898711491))
	}
}

type XRServer struct {
	Object
}

func (me *XRServer) BaseClass() string {
	return "XRServer"
}

func NewXRServer() *XRServer {
	str := StringNameFromStr("XRServer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRServer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XRServerTrackerType int

const (
	XRServerTrackerTypeTrackerHead        XRServerTrackerType = 1
	XRServerTrackerTypeTrackerController  XRServerTrackerType = 2
	XRServerTrackerTypeTrackerBasestation XRServerTrackerType = 4
	XRServerTrackerTypeTrackerAnchor      XRServerTrackerType = 8
	XRServerTrackerTypeTrackerAnyKnown    XRServerTrackerType = 127
	XRServerTrackerTypeTrackerUnknown     XRServerTrackerType = 128
	XRServerTrackerTypeTrackerAny         XRServerTrackerType = 255
)

type XRServerRotationMode int

const (
	XRServerRotationModeResetFullRotation XRServerRotationMode = 0
	XRServerRotationModeResetButKeepTilt  XRServerRotationMode = 1
	XRServerRotationModeDontResetRotation XRServerRotationMode = 2
)

func (me *XRServer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRServer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRServer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRServer) GetWorldScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetWorldScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRServer) SetWorldScale(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnSetWorldScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRServer) GetWorldOrigin() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetWorldOrigin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) SetWorldOrigin(world_origin Transform3D) {
	cargs := []gdc.ConstTypePtr{world_origin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnSetWorldOrigin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRServer) GetReferenceFrame() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetReferenceFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) CenterOnHmd(rotation_mode XRServerRotationMode, keep_height bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation_mode), gdc.ConstTypePtr(&keep_height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnCenterOnHmd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRServer) GetHmdTransform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetHmdTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) AddInterface(interface_ XRInterface) {
	cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnAddInterface), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRServer) GetInterfaceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetInterfaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRServer) RemoveInterface(interface_ XRInterface) {
	cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnRemoveInterface), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRServer) GetInterface(idx int64) XRInterface {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewXRInterface()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) GetInterfaces() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetInterfaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *XRServer) FindInterface(name String) XRInterface {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewXRInterface()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnFindInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) AddTracker(tracker XRPositionalTracker) {
	cargs := []gdc.ConstTypePtr{tracker.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnAddTracker), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRServer) RemoveTracker(tracker XRPositionalTracker) {
	cargs := []gdc.ConstTypePtr{tracker.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnRemoveTracker), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRServer) GetTrackers(tracker_types int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tracker_types)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&tracker_types)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetTrackers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) GetTracker(tracker_name StringName) XRPositionalTracker {
	cargs := []gdc.ConstTypePtr{tracker_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewXRPositionalTracker()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetTracker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) GetPrimaryInterface() XRInterface {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewXRInterface()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnGetPrimaryInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRServer) SetPrimaryInterface(interface_ XRInterface) {
	cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRServer.fnSetPrimaryInterface), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type XRServerInterfaceAddedSignalFn func(interface_name StringName)

func (me *XRServer) ConnectInterfaceAdded(subs SignalSubscribers, fn XRServerInterfaceAddedSignalFn) {
	sig := StringNameFromStr("interface_added")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectInterfaceAdded(subs SignalSubscribers, fn XRServerInterfaceAddedSignalFn) {
	sig := StringNameFromStr("interface_added")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerInterfaceRemovedSignalFn func(interface_name StringName)

func (me *XRServer) ConnectInterfaceRemoved(subs SignalSubscribers, fn XRServerInterfaceRemovedSignalFn) {
	sig := StringNameFromStr("interface_removed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectInterfaceRemoved(subs SignalSubscribers, fn XRServerInterfaceRemovedSignalFn) {
	sig := StringNameFromStr("interface_removed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerTrackerAddedSignalFn func(tracker_name StringName, type_ int)

func (me *XRServer) ConnectTrackerAdded(subs SignalSubscribers, fn XRServerTrackerAddedSignalFn) {
	sig := StringNameFromStr("tracker_added")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectTrackerAdded(subs SignalSubscribers, fn XRServerTrackerAddedSignalFn) {
	sig := StringNameFromStr("tracker_added")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerTrackerUpdatedSignalFn func(tracker_name StringName, type_ int)

func (me *XRServer) ConnectTrackerUpdated(subs SignalSubscribers, fn XRServerTrackerUpdatedSignalFn) {
	sig := StringNameFromStr("tracker_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectTrackerUpdated(subs SignalSubscribers, fn XRServerTrackerUpdatedSignalFn) {
	sig := StringNameFromStr("tracker_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type XRServerTrackerRemovedSignalFn func(tracker_name StringName, type_ int)

func (me *XRServer) ConnectTrackerRemoved(subs SignalSubscribers, fn XRServerTrackerRemovedSignalFn) {
	sig := StringNameFromStr("tracker_removed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *XRServer) DisconnectTrackerRemoved(subs SignalSubscribers, fn XRServerTrackerRemovedSignalFn) {
	sig := StringNameFromStr("tracker_removed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
