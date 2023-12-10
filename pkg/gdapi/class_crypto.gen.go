// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Crypto struct {
  obj gdc.ObjectPtr
}

func (me *Crypto) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Crypto) BaseClass() string {
  return "Crypto"
}



// Enums

func (me *Crypto) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Crypto) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Crypto) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Crypto) GenerateRandomBytes(size int, ) PackedByteArray {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_random_bytes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 47165747) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) GenerateRsa(size int, ) CryptoKey {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_rsa")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1237515462) // FIXME: should cache?
  var ret CryptoKey
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) GenerateSelfSignedCertificate(key CryptoKey, issuer_name String, not_before String, not_after String, ) X509Certificate {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_self_signed_certificate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 947314696) // FIXME: should cache?
  var ret X509Certificate
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(issuer_name.AsCTypePtr()), gdc.ConstTypePtr(not_before.AsCTypePtr()), gdc.ConstTypePtr(not_after.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) Sign(hash_type HashingContextHashType, hash PackedByteArray, key CryptoKey, ) PackedByteArray {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sign")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1673662703) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type), gdc.ConstTypePtr(hash.AsCTypePtr()), gdc.ConstTypePtr(key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) Verify(hash_type HashingContextHashType, hash PackedByteArray, signature PackedByteArray, key CryptoKey, ) bool {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("verify")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2805902225) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type), gdc.ConstTypePtr(hash.AsCTypePtr()), gdc.ConstTypePtr(signature.AsCTypePtr()), gdc.ConstTypePtr(key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) Encrypt(key CryptoKey, plaintext PackedByteArray, ) PackedByteArray {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("encrypt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2361793670) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(plaintext.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) Decrypt(key CryptoKey, ciphertext PackedByteArray, ) PackedByteArray {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("decrypt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2361793670) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(ciphertext.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) HmacDigest(hash_type HashingContextHashType, key PackedByteArray, msg PackedByteArray, ) PackedByteArray {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hmac_digest")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2368951203) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type), gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(msg.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Crypto) ConstantTimeCompare(trusted PackedByteArray, received PackedByteArray, ) bool {
  classNameV := StringNameFromStr("Crypto")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("constant_time_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1024142237) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(trusted.AsCTypePtr()), gdc.ConstTypePtr(received.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties