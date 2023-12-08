// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  UPNPDeviceIgdStatusOk UPNPDeviceIGDStatus = 0
  UPNPDeviceIgdStatusHttpError UPNPDeviceIGDStatus = 1
  UPNPDeviceIgdStatusHttpEmpty UPNPDeviceIGDStatus = 2
  UPNPDeviceIgdStatusNoUrls UPNPDeviceIGDStatus = 3
  UPNPDeviceIgdStatusNoIgd UPNPDeviceIGDStatus = 4
  UPNPDeviceIgdStatusDisconnected UPNPDeviceIGDStatus = 5
  UPNPDeviceIgdStatusUnknownDevice UPNPDeviceIGDStatus = 6
  UPNPDeviceIgdStatusInvalidControl UPNPDeviceIGDStatus = 7
  UPNPDeviceIgdStatusMallocError UPNPDeviceIGDStatus = 8
  UPNPDeviceIgdStatusUnknownError UPNPDeviceIGDStatus = 9
)
