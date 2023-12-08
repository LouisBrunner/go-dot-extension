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

func  (me *UPNP) GetDeviceCount() { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) GetDevice(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) AddDevice(device UPNPDevice, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) SetDevice(index int, device UPNPDevice, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) RemoveDevice(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) ClearDevices() { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) GetGateway() { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) Discover(timeout int, ttl int, device_filter String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) QueryExternalAddress() { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) AddPortMapping(port int, port_internal int, desc String, proto String, duration int, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) DeletePortMapping(port int, proto String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) SetDiscoverMulticastIf(m_if String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) GetDiscoverMulticastIf() { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) SetDiscoverLocalPort(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) GetDiscoverLocalPort() { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) SetDiscoverIpv6(ipv6 bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNP) IsDiscoverIpv6() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
