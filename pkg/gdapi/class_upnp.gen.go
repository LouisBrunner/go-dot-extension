// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  UPNPUpnpResultSuccess UPNPUPNPResult = 0
  UPNPUpnpResultNotAuthorized UPNPUPNPResult = 1
  UPNPUpnpResultPortMappingNotFound UPNPUPNPResult = 2
  UPNPUpnpResultInconsistentParameters UPNPUPNPResult = 3
  UPNPUpnpResultNoSuchEntryInArray UPNPUPNPResult = 4
  UPNPUpnpResultActionFailed UPNPUPNPResult = 5
  UPNPUpnpResultSrcIpWildcardNotPermitted UPNPUPNPResult = 6
  UPNPUpnpResultExtPortWildcardNotPermitted UPNPUPNPResult = 7
  UPNPUpnpResultIntPortWildcardNotPermitted UPNPUPNPResult = 8
  UPNPUpnpResultRemoteHostMustBeWildcard UPNPUPNPResult = 9
  UPNPUpnpResultExtPortMustBeWildcard UPNPUPNPResult = 10
  UPNPUpnpResultNoPortMapsAvailable UPNPUPNPResult = 11
  UPNPUpnpResultConflictWithOtherMechanism UPNPUPNPResult = 12
  UPNPUpnpResultConflictWithOtherMapping UPNPUPNPResult = 13
  UPNPUpnpResultSamePortValuesRequired UPNPUPNPResult = 14
  UPNPUpnpResultOnlyPermanentLeaseSupported UPNPUPNPResult = 15
  UPNPUpnpResultInvalidGateway UPNPUPNPResult = 16
  UPNPUpnpResultInvalidPort UPNPUPNPResult = 17
  UPNPUpnpResultInvalidProtocol UPNPUPNPResult = 18
  UPNPUpnpResultInvalidDuration UPNPUPNPResult = 19
  UPNPUpnpResultInvalidArgs UPNPUPNPResult = 20
  UPNPUpnpResultInvalidResponse UPNPUPNPResult = 21
  UPNPUpnpResultInvalidParam UPNPUPNPResult = 22
  UPNPUpnpResultHttpError UPNPUPNPResult = 23
  UPNPUpnpResultSocketError UPNPUPNPResult = 24
  UPNPUpnpResultMemAllocError UPNPUPNPResult = 25
  UPNPUpnpResultNoGateway UPNPUPNPResult = 26
  UPNPUpnpResultNoDevices UPNPUPNPResult = 27
  UPNPUpnpResultUnknownError UPNPUPNPResult = 28
)
