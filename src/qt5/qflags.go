package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qflags.h
// dst-file: /src/core/qflags.go
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
  // proto:  void QIncompatibleFlag::QIncompatibleFlag(int i);
extern void* dector_ZN17QIncompatibleFlagC1Ei(int arg0);
extern void _ZN17QIncompatibleFlagC1Ei(void* qthis, int arg0);
  // proto:  void QFlag::QFlag(ushort ai);
extern void* dector_ZN5QFlagC1Et(unsigned short arg0);
extern void _ZN5QFlagC1Et(void* qthis, unsigned short arg0);
  // proto:  void QFlag::QFlag(int ai);
extern void* dector_ZN5QFlagC1Ei(int arg0);
extern void _ZN5QFlagC1Ei(void* qthis, int arg0);
  // proto:  void QFlag::QFlag(short ai);
extern void* dector_ZN5QFlagC1Es(short arg0);
extern void _ZN5QFlagC1Es(void* qthis, short arg0);
  // proto:  void QFlag::QFlag(uint ai);
extern void* dector_ZN5QFlagC1Ej(unsigned int arg0);
extern void _ZN5QFlagC1Ej(void* qthis, unsigned int arg0);
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

// class sizeof(QIncompatibleFlag)=4
type QIncompatibleFlag struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFlag)=4
type QFlag struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QIncompatibleFlag::QIncompatibleFlag(int i);
func NewQIncompatibleFlag(args ...interface{}) QIncompatibleFlag {
  return QIncompatibleFlag{}
}

  // proto:  void QFlag::QFlag(ushort ai);
func NewQFlag(args ...interface{}) QFlag {
  return QFlag{}
}

// <= body block end

