// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type TimeMonth int
const (
  TimeMonthJanuary TimeMonth = 1
  TimeMonthFebruary TimeMonth = 2
  TimeMonthMarch TimeMonth = 3
  TimeMonthApril TimeMonth = 4
  TimeMonthMay TimeMonth = 5
  TimeMonthJune TimeMonth = 6
  TimeMonthJuly TimeMonth = 7
  TimeMonthAugust TimeMonth = 8
  TimeMonthSeptember TimeMonth = 9
  TimeMonthOctober TimeMonth = 10
  TimeMonthNovember TimeMonth = 11
  TimeMonthDecember TimeMonth = 12
)

type TimeWeekday int
const (
  TimeWeekdaySunday TimeWeekday = 0
  TimeWeekdayMonday TimeWeekday = 1
  TimeWeekdayTuesday TimeWeekday = 2
  TimeWeekdayWednesday TimeWeekday = 3
  TimeWeekdayThursday TimeWeekday = 4
  TimeWeekdayFriday TimeWeekday = 5
  TimeWeekdaySaturday TimeWeekday = 6
)
