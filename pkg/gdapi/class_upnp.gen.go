// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type UPNP struct {
  RefCounted
}

func (me *UPNP) BaseClass() string {
  return "UPNP"
}



// Enums

type UPNPUPNPResult int
const (
  UPNPUPNPResultUpnpResultSuccess UPNPUPNPResult = 0
  UPNPUPNPResultUpnpResultNotAuthorized UPNPUPNPResult = 1
  UPNPUPNPResultUpnpResultPortMappingNotFound UPNPUPNPResult = 2
  UPNPUPNPResultUpnpResultInconsistentParameters UPNPUPNPResult = 3
  UPNPUPNPResultUpnpResultNoSuchEntryInArray UPNPUPNPResult = 4
  UPNPUPNPResultUpnpResultActionFailed UPNPUPNPResult = 5
  UPNPUPNPResultUpnpResultSrcIpWildcardNotPermitted UPNPUPNPResult = 6
  UPNPUPNPResultUpnpResultExtPortWildcardNotPermitted UPNPUPNPResult = 7
  UPNPUPNPResultUpnpResultIntPortWildcardNotPermitted UPNPUPNPResult = 8
  UPNPUPNPResultUpnpResultRemoteHostMustBeWildcard UPNPUPNPResult = 9
  UPNPUPNPResultUpnpResultExtPortMustBeWildcard UPNPUPNPResult = 10
  UPNPUPNPResultUpnpResultNoPortMapsAvailable UPNPUPNPResult = 11
  UPNPUPNPResultUpnpResultConflictWithOtherMechanism UPNPUPNPResult = 12
  UPNPUPNPResultUpnpResultConflictWithOtherMapping UPNPUPNPResult = 13
  UPNPUPNPResultUpnpResultSamePortValuesRequired UPNPUPNPResult = 14
  UPNPUPNPResultUpnpResultOnlyPermanentLeaseSupported UPNPUPNPResult = 15
  UPNPUPNPResultUpnpResultInvalidGateway UPNPUPNPResult = 16
  UPNPUPNPResultUpnpResultInvalidPort UPNPUPNPResult = 17
  UPNPUPNPResultUpnpResultInvalidProtocol UPNPUPNPResult = 18
  UPNPUPNPResultUpnpResultInvalidDuration UPNPUPNPResult = 19
  UPNPUPNPResultUpnpResultInvalidArgs UPNPUPNPResult = 20
  UPNPUPNPResultUpnpResultInvalidResponse UPNPUPNPResult = 21
  UPNPUPNPResultUpnpResultInvalidParam UPNPUPNPResult = 22
  UPNPUPNPResultUpnpResultHttpError UPNPUPNPResult = 23
  UPNPUPNPResultUpnpResultSocketError UPNPUPNPResult = 24
  UPNPUPNPResultUpnpResultMemAllocError UPNPUPNPResult = 25
  UPNPUPNPResultUpnpResultNoGateway UPNPUPNPResult = 26
  UPNPUPNPResultUpnpResultNoDevices UPNPUPNPResult = 27
  UPNPUPNPResultUpnpResultUnknownError UPNPUPNPResult = 28
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

func  (me *UPNP) GetDeviceCount() int {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_device_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) GetDevice(index int, ) UPNPDevice {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2193290270) // FIXME: should cache?
  var ret UPNPDevice
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) AddDevice(device UPNPDevice, )  {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 986715920) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(device.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *UPNP) SetDevice(index int, device UPNPDevice, )  {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3015133723) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(device.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *UPNP) RemoveDevice(index int, )  {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *UPNP) ClearDevices()  {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_devices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *UPNP) GetGateway() UPNPDevice {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gateway")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2276800779) // FIXME: should cache?
  var ret UPNPDevice
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) Discover(timeout int, ttl int, device_filter String, ) int {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("discover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1575334765) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout), gdc.ConstTypePtr(&ttl), gdc.ConstTypePtr(device_filter.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) QueryExternalAddress() String {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("query_external_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) AddPortMapping(port int, port_internal int, desc String, proto String, duration int, ) int {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_port_mapping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 818314583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&port_internal), gdc.ConstTypePtr(desc.AsCTypePtr()), gdc.ConstTypePtr(proto.AsCTypePtr()), gdc.ConstTypePtr(&duration), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) DeletePortMapping(port int, proto String, ) int {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delete_port_mapping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444187325) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(proto.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) SetDiscoverMulticastIf(m_if String, )  {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_discover_multicast_if")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(m_if.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *UPNP) GetDiscoverMulticastIf() String {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_discover_multicast_if")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) SetDiscoverLocalPort(port int, )  {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_discover_local_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *UPNP) GetDiscoverLocalPort() int {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_discover_local_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *UPNP) SetDiscoverIpv6(ipv6 bool, )  {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_discover_ipv6")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ipv6), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *UPNP) IsDiscoverIpv6() bool {
  classNameV := StringNameFromStr("UPNP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_discover_ipv6")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
