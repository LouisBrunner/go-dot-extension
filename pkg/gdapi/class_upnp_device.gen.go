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



// Enums

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

func (me *UPNPDevice) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *UPNPDevice) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *UPNPDevice) IsValidGateway()  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) QueryExternalAddress()  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) AddPortMapping(port int, port_internal int, desc String, proto String, duration int, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) DeletePortMapping(port int, proto String, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) SetDescriptionUrl(url String, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) GetDescriptionUrl()  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) SetServiceType(type_ String, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) GetServiceType()  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) SetIgdControlUrl(url String, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) GetIgdControlUrl()  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) SetIgdServiceType(type_ String, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) GetIgdServiceType()  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) SetIgdOurAddr(addr String, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) GetIgdOurAddr()  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) SetIgdStatus(status UPNPDeviceIGDStatus, )  {
  panic("TODO: implement")
}

func  (me *UPNPDevice) GetIgdStatus()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
