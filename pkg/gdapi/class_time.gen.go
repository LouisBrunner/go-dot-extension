// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Time struct {
  obj gdc.ObjectPtr
}

func (me *Time) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Time) BaseClass() string {
  return "Time"
}



// Enums

type TimeMonth int
const (
  TimeMonthMonthJanuary TimeMonth = 1
  TimeMonthMonthFebruary TimeMonth = 2
  TimeMonthMonthMarch TimeMonth = 3
  TimeMonthMonthApril TimeMonth = 4
  TimeMonthMonthMay TimeMonth = 5
  TimeMonthMonthJune TimeMonth = 6
  TimeMonthMonthJuly TimeMonth = 7
  TimeMonthMonthAugust TimeMonth = 8
  TimeMonthMonthSeptember TimeMonth = 9
  TimeMonthMonthOctober TimeMonth = 10
  TimeMonthMonthNovember TimeMonth = 11
  TimeMonthMonthDecember TimeMonth = 12
)

type TimeWeekday int
const (
  TimeWeekdayWeekdaySunday TimeWeekday = 0
  TimeWeekdayWeekdayMonday TimeWeekday = 1
  TimeWeekdayWeekdayTuesday TimeWeekday = 2
  TimeWeekdayWeekdayWednesday TimeWeekday = 3
  TimeWeekdayWeekdayThursday TimeWeekday = 4
  TimeWeekdayWeekdayFriday TimeWeekday = 5
  TimeWeekdayWeekdaySaturday TimeWeekday = 6
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

func  (me *Time) GetDatetimeDictFromUnixTime(unix_time_val int, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDateDictFromUnixTime(unix_time_val int, )  {
  panic("TODO: implement")
}

func  (me *Time) GetTimeDictFromUnixTime(unix_time_val int, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDatetimeStringFromUnixTime(unix_time_val int, use_space bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDateStringFromUnixTime(unix_time_val int, )  {
  panic("TODO: implement")
}

func  (me *Time) GetTimeStringFromUnixTime(unix_time_val int, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDatetimeDictFromDatetimeString(datetime String, weekday bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDatetimeStringFromDatetimeDict(datetime Dictionary, use_space bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetUnixTimeFromDatetimeDict(datetime Dictionary, )  {
  panic("TODO: implement")
}

func  (me *Time) GetUnixTimeFromDatetimeString(datetime String, )  {
  panic("TODO: implement")
}

func  (me *Time) GetOffsetStringFromOffsetMinutes(offset_minutes int, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDatetimeDictFromSystem(utc bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDateDictFromSystem(utc bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetTimeDictFromSystem(utc bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDatetimeStringFromSystem(utc bool, use_space bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetDateStringFromSystem(utc bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetTimeStringFromSystem(utc bool, )  {
  panic("TODO: implement")
}

func  (me *Time) GetTimeZoneFromSystem()  {
  panic("TODO: implement")
}

func  (me *Time) GetUnixTimeFromSystem()  {
  panic("TODO: implement")
}

func  (me *Time) GetTicksMsec()  {
  panic("TODO: implement")
}

func  (me *Time) GetTicksUsec()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
