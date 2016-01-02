package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qhash.h
// dst-file: /src/core/qhash.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QHashData::hasShrunk();
extern void _ZN9QHashData9hasShrunkEv(void* qthis);
  // proto:  void * QHashData::allocateNode(int nodeAlign);
extern void _ZN9QHashData12allocateNodeEi(void* qthis, int arg0);
  // proto:  bool QHashData::willGrow();
extern void _ZN9QHashData8willGrowEv(void* qthis);
  // proto:  void QHashData::rehash(int hint);
extern void _ZN9QHashData6rehashEi(void* qthis, int arg0);
  // proto:  void QHashData::freeNode(void * node);
extern void _ZN9QHashData8freeNodeEPv(void* qthis, void* arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QHashDummyValue)=1
type QHashDummyValue struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHashData)=1
type QHashData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QHashData::hasShrunk();
func (this *QHashData) hasShrunk(args ...interface{}) () {
  // hasShrunk()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData9hasShrunkEv
  default:
    qtrt.ErrorResolve("QHashData", "hasShrunk", args)
  }

}

  // proto:  void * QHashData::allocateNode(int nodeAlign);
func (this *QHashData) allocateNode(args ...interface{}) () {
  // allocateNode(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData12allocateNodeEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QHashData", "allocateNode", args)
  }

}

  // proto:  bool QHashData::willGrow();
func (this *QHashData) willGrow(args ...interface{}) () {
  // willGrow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData8willGrowEv
  default:
    qtrt.ErrorResolve("QHashData", "willGrow", args)
  }

}

  // proto:  void QHashData::rehash(int hint);
func (this *QHashData) rehash(args ...interface{}) () {
  // rehash(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData6rehashEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QHashData", "rehash", args)
  }

}

  // proto:  void QHashData::freeNode(void * node);
func (this *QHashData) freeNode(args ...interface{}) () {
  // freeNode(void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QHashData8freeNodeEPv
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QHashData", "freeNode", args)
  }

}

// <= body block end

