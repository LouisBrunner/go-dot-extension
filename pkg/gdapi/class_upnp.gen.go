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

type ptrsForUPNPList struct {
	fnGetDeviceCount         gdc.MethodBindPtr
	fnGetDevice              gdc.MethodBindPtr
	fnAddDevice              gdc.MethodBindPtr
	fnSetDevice              gdc.MethodBindPtr
	fnRemoveDevice           gdc.MethodBindPtr
	fnClearDevices           gdc.MethodBindPtr
	fnGetGateway             gdc.MethodBindPtr
	fnDiscover               gdc.MethodBindPtr
	fnQueryExternalAddress   gdc.MethodBindPtr
	fnAddPortMapping         gdc.MethodBindPtr
	fnDeletePortMapping      gdc.MethodBindPtr
	fnSetDiscoverMulticastIf gdc.MethodBindPtr
	fnGetDiscoverMulticastIf gdc.MethodBindPtr
	fnSetDiscoverLocalPort   gdc.MethodBindPtr
	fnGetDiscoverLocalPort   gdc.MethodBindPtr
	fnSetDiscoverIpv6        gdc.MethodBindPtr
	fnIsDiscoverIpv6         gdc.MethodBindPtr
}

var ptrsForUPNP ptrsForUPNPList

func initUPNPPtrs(iface gdc.Interface) {

	className := StringNameFromStr("UPNP")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_device_count")
		defer methodName.Destroy()
		ptrsForUPNP.fnGetDeviceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_device")
		defer methodName.Destroy()
		ptrsForUPNP.fnGetDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2193290270))
	}
	{
		methodName := StringNameFromStr("add_device")
		defer methodName.Destroy()
		ptrsForUPNP.fnAddDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 986715920))
	}
	{
		methodName := StringNameFromStr("set_device")
		defer methodName.Destroy()
		ptrsForUPNP.fnSetDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3015133723))
	}
	{
		methodName := StringNameFromStr("remove_device")
		defer methodName.Destroy()
		ptrsForUPNP.fnRemoveDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("clear_devices")
		defer methodName.Destroy()
		ptrsForUPNP.fnClearDevices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_gateway")
		defer methodName.Destroy()
		ptrsForUPNP.fnGetGateway = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2276800779))
	}
	{
		methodName := StringNameFromStr("discover")
		defer methodName.Destroy()
		ptrsForUPNP.fnDiscover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1575334765))
	}
	{
		methodName := StringNameFromStr("query_external_address")
		defer methodName.Destroy()
		ptrsForUPNP.fnQueryExternalAddress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("add_port_mapping")
		defer methodName.Destroy()
		ptrsForUPNP.fnAddPortMapping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 818314583))
	}
	{
		methodName := StringNameFromStr("delete_port_mapping")
		defer methodName.Destroy()
		ptrsForUPNP.fnDeletePortMapping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444187325))
	}
	{
		methodName := StringNameFromStr("set_discover_multicast_if")
		defer methodName.Destroy()
		ptrsForUPNP.fnSetDiscoverMulticastIf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_discover_multicast_if")
		defer methodName.Destroy()
		ptrsForUPNP.fnGetDiscoverMulticastIf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_discover_local_port")
		defer methodName.Destroy()
		ptrsForUPNP.fnSetDiscoverLocalPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_discover_local_port")
		defer methodName.Destroy()
		ptrsForUPNP.fnGetDiscoverLocalPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_discover_ipv6")
		defer methodName.Destroy()
		ptrsForUPNP.fnSetDiscoverIpv6 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_discover_ipv6")
		defer methodName.Destroy()
		ptrsForUPNP.fnIsDiscoverIpv6 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type UPNP struct {
	RefCounted
}

func (me *UPNP) BaseClass() string {
	return "UPNP"
}

func NewUPNP() *UPNP {
	str := StringNameFromStr("UPNP") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &UPNP{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type UPNPUPNPResult int

const (
	UPNPUPNPResultUpnpResultSuccess                     UPNPUPNPResult = 0
	UPNPUPNPResultUpnpResultNotAuthorized               UPNPUPNPResult = 1
	UPNPUPNPResultUpnpResultPortMappingNotFound         UPNPUPNPResult = 2
	UPNPUPNPResultUpnpResultInconsistentParameters      UPNPUPNPResult = 3
	UPNPUPNPResultUpnpResultNoSuchEntryInArray          UPNPUPNPResult = 4
	UPNPUPNPResultUpnpResultActionFailed                UPNPUPNPResult = 5
	UPNPUPNPResultUpnpResultSrcIpWildcardNotPermitted   UPNPUPNPResult = 6
	UPNPUPNPResultUpnpResultExtPortWildcardNotPermitted UPNPUPNPResult = 7
	UPNPUPNPResultUpnpResultIntPortWildcardNotPermitted UPNPUPNPResult = 8
	UPNPUPNPResultUpnpResultRemoteHostMustBeWildcard    UPNPUPNPResult = 9
	UPNPUPNPResultUpnpResultExtPortMustBeWildcard       UPNPUPNPResult = 10
	UPNPUPNPResultUpnpResultNoPortMapsAvailable         UPNPUPNPResult = 11
	UPNPUPNPResultUpnpResultConflictWithOtherMechanism  UPNPUPNPResult = 12
	UPNPUPNPResultUpnpResultConflictWithOtherMapping    UPNPUPNPResult = 13
	UPNPUPNPResultUpnpResultSamePortValuesRequired      UPNPUPNPResult = 14
	UPNPUPNPResultUpnpResultOnlyPermanentLeaseSupported UPNPUPNPResult = 15
	UPNPUPNPResultUpnpResultInvalidGateway              UPNPUPNPResult = 16
	UPNPUPNPResultUpnpResultInvalidPort                 UPNPUPNPResult = 17
	UPNPUPNPResultUpnpResultInvalidProtocol             UPNPUPNPResult = 18
	UPNPUPNPResultUpnpResultInvalidDuration             UPNPUPNPResult = 19
	UPNPUPNPResultUpnpResultInvalidArgs                 UPNPUPNPResult = 20
	UPNPUPNPResultUpnpResultInvalidResponse             UPNPUPNPResult = 21
	UPNPUPNPResultUpnpResultInvalidParam                UPNPUPNPResult = 22
	UPNPUPNPResultUpnpResultHttpError                   UPNPUPNPResult = 23
	UPNPUPNPResultUpnpResultSocketError                 UPNPUPNPResult = 24
	UPNPUPNPResultUpnpResultMemAllocError               UPNPUPNPResult = 25
	UPNPUPNPResultUpnpResultNoGateway                   UPNPUPNPResult = 26
	UPNPUPNPResultUpnpResultNoDevices                   UPNPUPNPResult = 27
	UPNPUPNPResultUpnpResultUnknownError                UPNPUPNPResult = 28
)

func (me *UPNP) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *UPNP) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *UPNP) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *UPNP) GetDeviceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnGetDeviceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNP) GetDevice(index int64) UPNPDevice {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewUPNPDevice()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnGetDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNP) AddDevice(device UPNPDevice) {
	cargs := []gdc.ConstTypePtr{device.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnAddDevice), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNP) SetDevice(index int64, device UPNPDevice) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), device.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnSetDevice), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNP) RemoveDevice(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnRemoveDevice), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNP) ClearDevices() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnClearDevices), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNP) GetGateway() UPNPDevice {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewUPNPDevice()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnGetGateway), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNP) Discover(timeout int64, ttl int64, device_filter String) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout), gdc.ConstTypePtr(&ttl), device_filter.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&timeout)
	pinner.Pin(&ttl)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnDiscover), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNP) QueryExternalAddress() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnQueryExternalAddress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNP) AddPortMapping(port int64, port_internal int64, desc String, proto String, duration int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&port_internal), desc.AsCTypePtr(), proto.AsCTypePtr(), gdc.ConstTypePtr(&duration)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port)
	pinner.Pin(&port_internal)
	pinner.Pin(&duration)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnAddPortMapping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNP) DeletePortMapping(port int64, proto String) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), proto.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnDeletePortMapping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNP) SetDiscoverMulticastIf(m_if String) {
	cargs := []gdc.ConstTypePtr{m_if.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnSetDiscoverMulticastIf), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNP) GetDiscoverMulticastIf() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnGetDiscoverMulticastIf), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNP) SetDiscoverLocalPort(port int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnSetDiscoverLocalPort), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNP) GetDiscoverLocalPort() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnGetDiscoverLocalPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNP) SetDiscoverIpv6(ipv6 bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ipv6)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnSetDiscoverIpv6), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNP) IsDiscoverIpv6() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNP.fnIsDiscoverIpv6), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
