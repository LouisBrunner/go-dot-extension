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

type ptrsForTimeList struct {
	fnGetDatetimeDictFromUnixTime       gdc.MethodBindPtr
	fnGetDateDictFromUnixTime           gdc.MethodBindPtr
	fnGetTimeDictFromUnixTime           gdc.MethodBindPtr
	fnGetDatetimeStringFromUnixTime     gdc.MethodBindPtr
	fnGetDateStringFromUnixTime         gdc.MethodBindPtr
	fnGetTimeStringFromUnixTime         gdc.MethodBindPtr
	fnGetDatetimeDictFromDatetimeString gdc.MethodBindPtr
	fnGetDatetimeStringFromDatetimeDict gdc.MethodBindPtr
	fnGetUnixTimeFromDatetimeDict       gdc.MethodBindPtr
	fnGetUnixTimeFromDatetimeString     gdc.MethodBindPtr
	fnGetOffsetStringFromOffsetMinutes  gdc.MethodBindPtr
	fnGetDatetimeDictFromSystem         gdc.MethodBindPtr
	fnGetDateDictFromSystem             gdc.MethodBindPtr
	fnGetTimeDictFromSystem             gdc.MethodBindPtr
	fnGetDatetimeStringFromSystem       gdc.MethodBindPtr
	fnGetDateStringFromSystem           gdc.MethodBindPtr
	fnGetTimeStringFromSystem           gdc.MethodBindPtr
	fnGetTimeZoneFromSystem             gdc.MethodBindPtr
	fnGetUnixTimeFromSystem             gdc.MethodBindPtr
	fnGetTicksMsec                      gdc.MethodBindPtr
	fnGetTicksUsec                      gdc.MethodBindPtr
}

var ptrsForTime ptrsForTimeList

func initTimePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Time")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_datetime_dict_from_unix_time")
		defer methodName.Destroy()
		ptrsForTime.fnGetDatetimeDictFromUnixTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3485342025))
	}
	{
		methodName := StringNameFromStr("get_date_dict_from_unix_time")
		defer methodName.Destroy()
		ptrsForTime.fnGetDateDictFromUnixTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3485342025))
	}
	{
		methodName := StringNameFromStr("get_time_dict_from_unix_time")
		defer methodName.Destroy()
		ptrsForTime.fnGetTimeDictFromUnixTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3485342025))
	}
	{
		methodName := StringNameFromStr("get_datetime_string_from_unix_time")
		defer methodName.Destroy()
		ptrsForTime.fnGetDatetimeStringFromUnixTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311239925))
	}
	{
		methodName := StringNameFromStr("get_date_string_from_unix_time")
		defer methodName.Destroy()
		ptrsForTime.fnGetDateStringFromUnixTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_time_string_from_unix_time")
		defer methodName.Destroy()
		ptrsForTime.fnGetTimeStringFromUnixTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_datetime_dict_from_datetime_string")
		defer methodName.Destroy()
		ptrsForTime.fnGetDatetimeDictFromDatetimeString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3253569256))
	}
	{
		methodName := StringNameFromStr("get_datetime_string_from_datetime_dict")
		defer methodName.Destroy()
		ptrsForTime.fnGetDatetimeStringFromDatetimeDict = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1898123706))
	}
	{
		methodName := StringNameFromStr("get_unix_time_from_datetime_dict")
		defer methodName.Destroy()
		ptrsForTime.fnGetUnixTimeFromDatetimeDict = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3021115443))
	}
	{
		methodName := StringNameFromStr("get_unix_time_from_datetime_string")
		defer methodName.Destroy()
		ptrsForTime.fnGetUnixTimeFromDatetimeString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("get_offset_string_from_offset_minutes")
		defer methodName.Destroy()
		ptrsForTime.fnGetOffsetStringFromOffsetMinutes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_datetime_dict_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetDatetimeDictFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205769976))
	}
	{
		methodName := StringNameFromStr("get_date_dict_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetDateDictFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205769976))
	}
	{
		methodName := StringNameFromStr("get_time_dict_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetTimeDictFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205769976))
	}
	{
		methodName := StringNameFromStr("get_datetime_string_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetDatetimeStringFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1136425492))
	}
	{
		methodName := StringNameFromStr("get_date_string_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetDateStringFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1162154673))
	}
	{
		methodName := StringNameFromStr("get_time_string_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetTimeStringFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1162154673))
	}
	{
		methodName := StringNameFromStr("get_time_zone_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetTimeZoneFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_unix_time_from_system")
		defer methodName.Destroy()
		ptrsForTime.fnGetUnixTimeFromSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_ticks_msec")
		defer methodName.Destroy()
		ptrsForTime.fnGetTicksMsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_ticks_usec")
		defer methodName.Destroy()
		ptrsForTime.fnGetTicksUsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type Time struct {
	Object
}

func (me *Time) BaseClass() string {
	return "Time"
}

func NewTime() *Time {
	str := StringNameFromStr("Time") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Time{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TimeMonth int

const (
	TimeMonthMonthJanuary   TimeMonth = 1
	TimeMonthMonthFebruary  TimeMonth = 2
	TimeMonthMonthMarch     TimeMonth = 3
	TimeMonthMonthApril     TimeMonth = 4
	TimeMonthMonthMay       TimeMonth = 5
	TimeMonthMonthJune      TimeMonth = 6
	TimeMonthMonthJuly      TimeMonth = 7
	TimeMonthMonthAugust    TimeMonth = 8
	TimeMonthMonthSeptember TimeMonth = 9
	TimeMonthMonthOctober   TimeMonth = 10
	TimeMonthMonthNovember  TimeMonth = 11
	TimeMonthMonthDecember  TimeMonth = 12
)

type TimeWeekday int

const (
	TimeWeekdayWeekdaySunday    TimeWeekday = 0
	TimeWeekdayWeekdayMonday    TimeWeekday = 1
	TimeWeekdayWeekdayTuesday   TimeWeekday = 2
	TimeWeekdayWeekdayWednesday TimeWeekday = 3
	TimeWeekdayWeekdayThursday  TimeWeekday = 4
	TimeWeekdayWeekdayFriday    TimeWeekday = 5
	TimeWeekdayWeekdaySaturday  TimeWeekday = 6
)

func (me *Time) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Time) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Time) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Time) GetDatetimeDictFromUnixTime(unix_time_val int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&unix_time_val)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDatetimeDictFromUnixTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDateDictFromUnixTime(unix_time_val int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&unix_time_val)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDateDictFromUnixTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetTimeDictFromUnixTime(unix_time_val int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&unix_time_val)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetTimeDictFromUnixTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDatetimeStringFromUnixTime(unix_time_val int64, use_space bool) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val), gdc.ConstTypePtr(&use_space)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&unix_time_val)
	pinner.Pin(&use_space)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDatetimeStringFromUnixTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDateStringFromUnixTime(unix_time_val int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&unix_time_val)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDateStringFromUnixTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetTimeStringFromUnixTime(unix_time_val int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&unix_time_val)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetTimeStringFromUnixTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDatetimeDictFromDatetimeString(datetime String, weekday bool) Dictionary {
	cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr(), gdc.ConstTypePtr(&weekday)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&weekday)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDatetimeDictFromDatetimeString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDatetimeStringFromDatetimeDict(datetime Dictionary, use_space bool) String {
	cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr(), gdc.ConstTypePtr(&use_space)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&use_space)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDatetimeStringFromDatetimeDict), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetUnixTimeFromDatetimeDict(datetime Dictionary) int64 {
	cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetUnixTimeFromDatetimeDict), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Time) GetUnixTimeFromDatetimeString(datetime String) int64 {
	cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetUnixTimeFromDatetimeString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Time) GetOffsetStringFromOffsetMinutes(offset_minutes int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset_minutes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&offset_minutes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetOffsetStringFromOffsetMinutes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDatetimeDictFromSystem(utc bool) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&utc)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDatetimeDictFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDateDictFromSystem(utc bool) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&utc)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDateDictFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetTimeDictFromSystem(utc bool) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&utc)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetTimeDictFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDatetimeStringFromSystem(utc bool, use_space bool) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc), gdc.ConstTypePtr(&use_space)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&utc)
	pinner.Pin(&use_space)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDatetimeStringFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetDateStringFromSystem(utc bool) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&utc)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetDateStringFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetTimeStringFromSystem(utc bool) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&utc)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetTimeStringFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetTimeZoneFromSystem() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetTimeZoneFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Time) GetUnixTimeFromSystem() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetUnixTimeFromSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Time) GetTicksMsec() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetTicksMsec), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Time) GetTicksUsec() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTime.fnGetTicksUsec), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
