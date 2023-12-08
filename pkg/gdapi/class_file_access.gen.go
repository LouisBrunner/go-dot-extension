// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileAccess struct {
  obj gdc.ObjectPtr
}

func (me *FileAccess) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FileAccess) BaseClass() string {
  return "FileAccess"
}

type FileAccessModeFlags int
const (
  FileAccessModeFlagsRead FileAccessModeFlags = 1
  FileAccessModeFlagsWrite FileAccessModeFlags = 2
  FileAccessModeFlagsReadWrite FileAccessModeFlags = 3
  FileAccessModeFlagsWriteRead FileAccessModeFlags = 7
)

type FileAccessCompressionMode int
const (
  FileAccessCompressionModeCompressionFastlz FileAccessCompressionMode = 0
  FileAccessCompressionModeCompressionDeflate FileAccessCompressionMode = 1
  FileAccessCompressionModeCompressionZstd FileAccessCompressionMode = 2
  FileAccessCompressionModeCompressionGzip FileAccessCompressionMode = 3
  FileAccessCompressionModeCompressionBrotli FileAccessCompressionMode = 4
)

func  FileAccessOpen(path String, flags FileAccessModeFlags, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessOpenEncrypted(path String, mode_flags FileAccessModeFlags, key PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessOpenEncryptedWithPass(path String, mode_flags FileAccessModeFlags, pass String, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessOpenCompressed(path String, mode_flags FileAccessModeFlags, compression_mode FileAccessCompressionMode, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessGetOpenError() { // TODO: return value
  // TODO: implement
}

func  FileAccessGetFileAsBytes(path String, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessGetFileAsString(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Flush() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetPath() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetPathAbsolute() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) IsOpen() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Seek(position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) SeekEnd(position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetLength() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) EofReached() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Get8() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Get16() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Get32() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Get64() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetFloat() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetDouble() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetReal() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetBuffer(length int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetLine() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetCsvLine(delim String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetAsText(skip_cr bool, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessGetMd5(path String, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessGetSha256(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) IsBigEndian() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) SetBigEndian(big_endian bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetError() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetVar(allow_objects bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Store8(value int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Store16(value int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Store32(value int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Store64(value int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreFloat(value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreDouble(value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreReal(value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreBuffer(buffer PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreLine(line String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreCsvLine(values PackedStringArray, delim String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreString(string String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StoreVar(value Variant, full_objects bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) StorePascalString(string String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) GetPascalString() { // TODO: return value
  // TODO: implement
}

func  (me *FileAccess) Close() { // TODO: return value
  // TODO: implement
}

func  FileAccessFileExists(path String, ) { // TODO: return value
  // TODO: implement
}

func  FileAccessGetModifiedTime(file String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
