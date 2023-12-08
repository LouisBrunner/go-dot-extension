// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LineEdit struct {
  obj gdc.ObjectPtr
}

func (me *LineEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LineEdit) BaseClass() string {
  return "LineEdit"
}

type LineEditMenuItems int
const (
  LineEditMenuCut LineEditMenuItems = 0
  LineEditMenuCopy LineEditMenuItems = 1
  LineEditMenuPaste LineEditMenuItems = 2
  LineEditMenuClear LineEditMenuItems = 3
  LineEditMenuSelectAll LineEditMenuItems = 4
  LineEditMenuUndo LineEditMenuItems = 5
  LineEditMenuRedo LineEditMenuItems = 6
  LineEditMenuSubmenuTextDir LineEditMenuItems = 7
  LineEditMenuDirInherited LineEditMenuItems = 8
  LineEditMenuDirAuto LineEditMenuItems = 9
  LineEditMenuDirLtr LineEditMenuItems = 10
  LineEditMenuDirRtl LineEditMenuItems = 11
  LineEditMenuDisplayUcc LineEditMenuItems = 12
  LineEditMenuSubmenuInsertUcc LineEditMenuItems = 13
  LineEditMenuInsertLrm LineEditMenuItems = 14
  LineEditMenuInsertRlm LineEditMenuItems = 15
  LineEditMenuInsertLre LineEditMenuItems = 16
  LineEditMenuInsertRle LineEditMenuItems = 17
  LineEditMenuInsertLro LineEditMenuItems = 18
  LineEditMenuInsertRlo LineEditMenuItems = 19
  LineEditMenuInsertPdf LineEditMenuItems = 20
  LineEditMenuInsertAlm LineEditMenuItems = 21
  LineEditMenuInsertLri LineEditMenuItems = 22
  LineEditMenuInsertRli LineEditMenuItems = 23
  LineEditMenuInsertFsi LineEditMenuItems = 24
  LineEditMenuInsertPdi LineEditMenuItems = 25
  LineEditMenuInsertZwj LineEditMenuItems = 26
  LineEditMenuInsertZwnj LineEditMenuItems = 27
  LineEditMenuInsertWj LineEditMenuItems = 28
  LineEditMenuInsertShy LineEditMenuItems = 29
  LineEditMenuMax LineEditMenuItems = 30
)

type LineEditVirtualKeyboardType int
const (
  LineEditKeyboardTypeDefault LineEditVirtualKeyboardType = 0
  LineEditKeyboardTypeMultiline LineEditVirtualKeyboardType = 1
  LineEditKeyboardTypeNumber LineEditVirtualKeyboardType = 2
  LineEditKeyboardTypeNumberDecimal LineEditVirtualKeyboardType = 3
  LineEditKeyboardTypePhone LineEditVirtualKeyboardType = 4
  LineEditKeyboardTypeEmailAddress LineEditVirtualKeyboardType = 5
  LineEditKeyboardTypePassword LineEditVirtualKeyboardType = 6
  LineEditKeyboardTypeUrl LineEditVirtualKeyboardType = 7
)
