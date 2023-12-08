// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UPNPDevice struct {
  obj gdc.ObjectPtr
}

func (me *UPNPDevice) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *UPNPDevice) BaseClass() string {
  return "UPNPDevice"
}

type UPNPDeviceIGDStatus int
const (
  UPNPDeviceIGDStatusIgdStatusOk UPNPDeviceIGDStatus = 0
  UPNPDeviceIGDStatusIgdStatusHttpError UPNPDeviceIGDStatus = 1
  UPNPDeviceIGDStatusIgdStatusHttpEmpty UPNPDeviceIGDStatus = 2
  UPNPDeviceIGDStatusIgdStatusNoUrls UPNPDeviceIGDStatus = 3
  UPNPDeviceIGDStatusIgdStatusNoIgd UPNPDeviceIGDStatus = 4
  UPNPDeviceIGDStatusIgdStatusDisconnected UPNPDeviceIGDStatus = 5
  UPNPDeviceIGDStatusIgdStatusUnknownDevice UPNPDeviceIGDStatus = 6
  UPNPDeviceIGDStatusIgdStatusInvalidControl UPNPDeviceIGDStatus = 7
  UPNPDeviceIGDStatusIgdStatusMallocError UPNPDeviceIGDStatus = 8
  UPNPDeviceIGDStatusIgdStatusUnknownError UPNPDeviceIGDStatus = 9
)

func  (me *UPNPDevice) IsValidGateway() { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) QueryExternalAddress() { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) AddPortMapping(port int, port_internal int, desc String, proto String, duration int, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) DeletePortMapping(port int, proto String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) SetDescriptionUrl(url String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) GetDescriptionUrl() { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) SetServiceType(type_ String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) GetServiceType() { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) SetIgdControlUrl(url String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) GetIgdControlUrl() { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) SetIgdServiceType(type_ String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) GetIgdServiceType() { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) SetIgdOurAddr(addr String, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) GetIgdOurAddr() { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) SetIgdStatus(status UPNPDeviceIGDStatus, ) { // TODO: return value
  // TODO: implement
}

func  (me *UPNPDevice) GetIgdStatus() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
