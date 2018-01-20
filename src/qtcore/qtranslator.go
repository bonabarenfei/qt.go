//  header block begin
// /usr/include/qt/QtCore/qtranslator.h
// #include <qtranslator.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
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
type QTranslator struct {
	*QObject
}

func (this *QTranslator) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQTranslatorFromPointer(cthis unsafe.Pointer) *QTranslator {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTranslator{bcthis0}
}

// /usr/include/qt/QtCore/qtranslator.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTranslator) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtranslator.h:58
// index:0
// Public
// void QTranslator(class QObject *)
func NewQTranslator(parent unsafe.Pointer) *QTranslator {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTranslatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtranslator.h:59
// index:0
// Public virtual
// void ~QTranslator()
func DeleteQTranslator(*QTranslator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:61
// index:0
// Public virtual
// QString translate(const char *, const char *, const char *, int)
func (this *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) interface{} {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(sourceText)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator9translateEPKcS1_S1_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtranslator.h:64
// index:0
// Public virtual
// bool isEmpty()
func (this *QTranslator) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtranslator.h:66
// index:0
// Public
// bool load(const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QTranslator) Load(filename *QString, directory *QString, search_delimiters *QString, suffix *QString) interface{} {
	var convArg0 = filename.GetCthis()
	var convArg1 = directory.GetCthis()
	var convArg2 = search_delimiters.GetCthis()
	var convArg3 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QStringS2_S2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtranslator.h:70
// index:1
// Public
// bool load(const class QLocale &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QTranslator) Load_1(locale *QLocale, filename *QString, prefix *QString, directory *QString, suffix *QString) interface{} {
	var convArg0 = locale.GetCthis()
	var convArg1 = filename.GetCthis()
	var convArg2 = prefix.GetCthis()
	var convArg3 = directory.GetCthis()
	var convArg4 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtranslator.h:75
// index:2
// Public
// bool load(const uchar *, int, const class QString &)
func (this *QTranslator) Load_2(data unsafe.Pointer, len int, directory *QString) interface{} {
	var convArg2 = directory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadEPKhiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data, &len, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
