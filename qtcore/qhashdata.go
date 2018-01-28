package qtcore

// /usr/include/qt/QtCore/qhash.h
// #include <qhash.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QHashData struct {
	*qtrt.CObject
}

func (this *QHashData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHashData) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQHashDataFromPointer(cthis unsafe.Pointer) *QHashData {
	return &QHashData{&qtrt.CObject{cthis}}
}
func (*QHashData) NewFromPointer(cthis unsafe.Pointer) *QHashData {
	return NewQHashDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qhash.h:84
// index:0
// Public
// void * allocateNode(int)
func (this *QHashData) AllocateNode(nodeAlign int) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData12allocateNodeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), nodeAlign)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qhash.h:85
// index:0
// Public
// void freeNode(void *)
func (this *QHashData) FreeNode(node unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData8freeNodeEPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), node)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:86
// index:0
// Public
// QHashData * detach_helper(void (*)(struct QHashData::Node *, void *), void (*)(struct QHashData::Node *), int, int)
func (this *QHashData) Detach_helper(node_duplicate unsafe.Pointer /*666*/, node_delete unsafe.Pointer /*666*/, nodeSize int, nodeAlign int) *QHashData /*777 QHashData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData13detach_helperEPFvPNS_4NodeEPvEPFvS1_Eii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), node_duplicate, node_delete, nodeSize, nodeAlign)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQHashDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qhash.h:88
// index:0
// Public
// bool willGrow()
func (this *QHashData) WillGrow() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData8willGrowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qhash.h:89
// index:0
// Public
// void hasShrunk()
func (this *QHashData) HasShrunk() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData9hasShrunkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:90
// index:0
// Public
// void rehash(int)
func (this *QHashData) Rehash(hint int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData6rehashEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:91
// index:0
// Public
// void free_helper(void (*)(struct QHashData::Node *))
func (this *QHashData) Free_helper(node_delete unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData11free_helperEPFvPNS_4NodeEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), node_delete)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhash.h:92
// index:0
// Public
// QHashData::Node * firstNode()
func (this *QHashData) FirstNode() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHashData9firstNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

//  body block end
