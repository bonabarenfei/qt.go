//  header block begin
// /usr/include/qt/QtCore/qbuffer.h
// #include <qbuffer.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 57
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
type QBuffer struct {
	*QIODevice
}

func (this *QBuffer) GetCthis() unsafe.Pointer {
	return this.QIODevice.GetCthis()
}
func NewQBufferFromPointer(cthis unsafe.Pointer) *QBuffer {
	bcthis0 := NewQIODeviceFromPointer(cthis)
	return &QBuffer{bcthis0}
}

// /usr/include/qt/QtCore/qbuffer.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QBuffer) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:60
// index:0
// Public
// void QBuffer(class QObject *)
func NewQBuffer(parent unsafe.Pointer) *QBuffer {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:61
// index:1
// Public
// void QBuffer(class QByteArray *, class QObject *)
func NewQBuffer_1(buf unsafe.Pointer, parent unsafe.Pointer) *QBuffer {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP10QByteArrayP7QObject", ffiqt.FFI_TYPE_VOID, cthis, buf, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:66
// index:0
// Public virtual
// void ~QBuffer()
func DeleteQBuffer(*QBuffer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:68
// index:0
// Public
// QByteArray & buffer()
func (this *QBuffer) Buffer() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer6bufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:69
// index:1
// Public
// const QByteArray & buffer()
func (this *QBuffer) Buffer_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer6bufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:70
// index:0
// Public
// void setBuffer(class QByteArray *)
func (this *QBuffer) SetBuffer(a unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer9setBufferEP10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:72
// index:0
// Public
// void setData(const class QByteArray &)
func (this *QBuffer) SetData(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:73
// index:1
// Public inline
// void setData(const char *, int)
func (this *QBuffer) SetData_1(data string, len int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:74
// index:0
// Public
// const QByteArray & data()
func (this *QBuffer) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:78
// index:0
// Public virtual
// void close()
func (this *QBuffer) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:79
// index:0
// Public virtual
// qint64 size()
func (this *QBuffer) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:80
// index:0
// Public virtual
// qint64 pos()
func (this *QBuffer) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:81
// index:0
// Public virtual
// bool seek(qint64)
func (this *QBuffer) Seek(off int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer4seekEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &off)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:82
// index:0
// Public virtual
// bool atEnd()
func (this *QBuffer) AtEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:83
// index:0
// Public virtual
// bool canReadLine()
func (this *QBuffer) CanReadLine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:87
// index:0
// Protected virtual
// void connectNotify(const class QMetaMethod &)
func (this *QBuffer) ConnectNotify(arg0 *QMetaMethod) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:88
// index:0
// Protected virtual
// void disconnectNotify(const class QMetaMethod &)
func (this *QBuffer) DisconnectNotify(arg0 *QMetaMethod) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:90
// index:0
// Protected virtual
// qint64 readData(char *, qint64)
func (this *QBuffer) ReadData(data string, maxlen int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbuffer.h:91
// index:0
// Protected virtual
// qint64 writeData(const char *, qint64)
func (this *QBuffer) WriteData(data string, len int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
