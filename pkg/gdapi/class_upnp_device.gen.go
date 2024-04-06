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

type ptrsForUPNPDeviceList struct {
	fnIsValidGateway       gdc.MethodBindPtr
	fnQueryExternalAddress gdc.MethodBindPtr
	fnAddPortMapping       gdc.MethodBindPtr
	fnDeletePortMapping    gdc.MethodBindPtr
	fnSetDescriptionUrl    gdc.MethodBindPtr
	fnGetDescriptionUrl    gdc.MethodBindPtr
	fnSetServiceType       gdc.MethodBindPtr
	fnGetServiceType       gdc.MethodBindPtr
	fnSetIgdControlUrl     gdc.MethodBindPtr
	fnGetIgdControlUrl     gdc.MethodBindPtr
	fnSetIgdServiceType    gdc.MethodBindPtr
	fnGetIgdServiceType    gdc.MethodBindPtr
	fnSetIgdOurAddr        gdc.MethodBindPtr
	fnGetIgdOurAddr        gdc.MethodBindPtr
	fnSetIgdStatus         gdc.MethodBindPtr
	fnGetIgdStatus         gdc.MethodBindPtr
}

var ptrsForUPNPDevice ptrsForUPNPDeviceList

func initUPNPDevicePtrs(iface gdc.Interface) {

	className := StringNameFromStr("UPNPDevice")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_valid_gateway")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnIsValidGateway = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("query_external_address")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnQueryExternalAddress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("add_port_mapping")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnAddPortMapping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 818314583))
	}
	{
		methodName := StringNameFromStr("delete_port_mapping")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnDeletePortMapping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444187325))
	}
	{
		methodName := StringNameFromStr("set_description_url")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnSetDescriptionUrl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_description_url")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnGetDescriptionUrl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_service_type")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnSetServiceType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_service_type")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnGetServiceType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_igd_control_url")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnSetIgdControlUrl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_igd_control_url")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnGetIgdControlUrl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_igd_service_type")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnSetIgdServiceType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_igd_service_type")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnGetIgdServiceType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_igd_our_addr")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnSetIgdOurAddr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_igd_our_addr")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnGetIgdOurAddr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_igd_status")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnSetIgdStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 519504122))
	}
	{
		methodName := StringNameFromStr("get_igd_status")
		defer methodName.Destroy()
		ptrsForUPNPDevice.fnGetIgdStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 180887011))
	}

}

type UPNPDevice struct {
	RefCounted
}

func (me *UPNPDevice) BaseClass() string {
	return "UPNPDevice"
}

func NewUPNPDevice() *UPNPDevice {
	str := StringNameFromStr("UPNPDevice") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &UPNPDevice{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type UPNPDeviceIGDStatus int

const (
	UPNPDeviceIGDStatusIgdStatusOk             UPNPDeviceIGDStatus = 0
	UPNPDeviceIGDStatusIgdStatusHttpError      UPNPDeviceIGDStatus = 1
	UPNPDeviceIGDStatusIgdStatusHttpEmpty      UPNPDeviceIGDStatus = 2
	UPNPDeviceIGDStatusIgdStatusNoUrls         UPNPDeviceIGDStatus = 3
	UPNPDeviceIGDStatusIgdStatusNoIgd          UPNPDeviceIGDStatus = 4
	UPNPDeviceIGDStatusIgdStatusDisconnected   UPNPDeviceIGDStatus = 5
	UPNPDeviceIGDStatusIgdStatusUnknownDevice  UPNPDeviceIGDStatus = 6
	UPNPDeviceIGDStatusIgdStatusInvalidControl UPNPDeviceIGDStatus = 7
	UPNPDeviceIGDStatusIgdStatusMallocError    UPNPDeviceIGDStatus = 8
	UPNPDeviceIGDStatusIgdStatusUnknownError   UPNPDeviceIGDStatus = 9
)

func (me *UPNPDevice) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *UPNPDevice) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *UPNPDevice) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *UPNPDevice) IsValidGateway() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnIsValidGateway), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNPDevice) QueryExternalAddress() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnQueryExternalAddress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNPDevice) AddPortMapping(port int64, port_internal int64, desc String, proto String, duration int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&port_internal), desc.AsCTypePtr(), proto.AsCTypePtr(), gdc.ConstTypePtr(&duration)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port)
	pinner.Pin(&port_internal)
	pinner.Pin(&duration)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnAddPortMapping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNPDevice) DeletePortMapping(port int64, proto String) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), proto.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnDeletePortMapping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UPNPDevice) SetDescriptionUrl(url String) {
	cargs := []gdc.ConstTypePtr{url.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnSetDescriptionUrl), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNPDevice) GetDescriptionUrl() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnGetDescriptionUrl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNPDevice) SetServiceType(type_ String) {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnSetServiceType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNPDevice) GetServiceType() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnGetServiceType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNPDevice) SetIgdControlUrl(url String) {
	cargs := []gdc.ConstTypePtr{url.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnSetIgdControlUrl), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNPDevice) GetIgdControlUrl() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnGetIgdControlUrl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNPDevice) SetIgdServiceType(type_ String) {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnSetIgdServiceType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNPDevice) GetIgdServiceType() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnGetIgdServiceType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNPDevice) SetIgdOurAddr(addr String) {
	cargs := []gdc.ConstTypePtr{addr.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnSetIgdOurAddr), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNPDevice) GetIgdOurAddr() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnGetIgdOurAddr), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UPNPDevice) SetIgdStatus(status UPNPDeviceIGDStatus) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&status)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnSetIgdStatus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UPNPDevice) GetIgdStatus() UPNPDeviceIGDStatus {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret UPNPDeviceIGDStatus

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUPNPDevice.fnGetIgdStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
