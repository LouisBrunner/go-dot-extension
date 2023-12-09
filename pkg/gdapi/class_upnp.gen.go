// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UPNP struct {
  obj gdc.ObjectPtr
}

func (me *UPNP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func  (me *UPNP) GetDeviceCount()  {
  panic("TODO: implement")
}

func  (me *UPNP) GetDevice(index int, )  {
  panic("TODO: implement")
}

func  (me *UPNP) AddDevice(device UPNPDevice, )  {
  panic("TODO: implement")
}

func  (me *UPNP) SetDevice(index int, device UPNPDevice, )  {
  panic("TODO: implement")
}

func  (me *UPNP) RemoveDevice(index int, )  {
  panic("TODO: implement")
}

func  (me *UPNP) ClearDevices()  {
  panic("TODO: implement")
}

func  (me *UPNP) GetGateway()  {
  panic("TODO: implement")
}

func  (me *UPNP) Discover(timeout int, ttl int, device_filter String, )  {
  panic("TODO: implement")
}

func  (me *UPNP) QueryExternalAddress()  {
  panic("TODO: implement")
}

func  (me *UPNP) AddPortMapping(port int, port_internal int, desc String, proto String, duration int, )  {
  panic("TODO: implement")
}

func  (me *UPNP) DeletePortMapping(port int, proto String, )  {
  panic("TODO: implement")
}

func  (me *UPNP) SetDiscoverMulticastIf(m_if String, )  {
  panic("TODO: implement")
}

func  (me *UPNP) GetDiscoverMulticastIf()  {
  panic("TODO: implement")
}

func  (me *UPNP) SetDiscoverLocalPort(port int, )  {
  panic("TODO: implement")
}

func  (me *UPNP) GetDiscoverLocalPort()  {
  panic("TODO: implement")
}

func  (me *UPNP) SetDiscoverIpv6(ipv6 bool, )  {
  panic("TODO: implement")
}

func  (me *UPNP) IsDiscoverIpv6()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
