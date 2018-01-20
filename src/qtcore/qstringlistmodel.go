//  header block begin
// /usr/include/qt/QtCore/qstringlistmodel.h
// #include <qstringlistmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QStringListModel struct {
	*QAbstractListModel
}

func (this *QStringListModel) GetCthis() unsafe.Pointer {
	return this.QAbstractListModel.GetCthis()
}
func NewQStringListModelFromPointer(cthis unsafe.Pointer) *QStringListModel {
	bcthis0 := NewQAbstractListModelFromPointer(cthis)
	return &QStringListModel{bcthis0}
}

// /usr/include/qt/QtCore/qstringlistmodel.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStringListModel) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:55
// index:0
// Public
// void QStringListModel(class QObject *)
func NewQStringListModel(parent unsafe.Pointer) *QStringListModel {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringlistmodel.h:56
// index:1
// Public
// void QStringListModel(const class QStringList &, class QObject *)
func NewQStringListModel_1(strings *QStringList, parent unsafe.Pointer) *QStringListModel {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = strings.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModelC2ERK11QStringListP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringlistmodel.h:58
// index:0
// Public virtual
// int rowCount(const class QModelIndex &)
func (this *QStringListModel) RowCount(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:59
// index:0
// Public virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QStringListModel) Sibling(row int, column int, idx *QModelIndex) interface{} {
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:61
// index:0
// Public virtual
// QVariant data(const class QModelIndex &, int)
func (this *QStringListModel) Data(index *QModelIndex, role int) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:62
// index:0
// Public virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QStringListModel) SetData(index *QModelIndex, value *QVariant, role int) interface{} {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:64
// index:0
// Public virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QStringListModel) Flags(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:66
// index:0
// Public virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QStringListModel) InsertRows(row int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:67
// index:0
// Public virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QStringListModel) RemoveRows(row int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:69
// index:0
// Public virtual
// void sort(int, Qt::SortOrder)
func (this *QStringListModel) Sort(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:71
// index:0
// Public
// QStringList stringList()
func (this *QStringListModel) StringList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel10stringListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringlistmodel.h:72
// index:0
// Public
// void setStringList(const class QStringList &)
func (this *QStringListModel) SetStringList(strings *QStringList) {
	var convArg0 = strings.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStringListModel13setStringListERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:74
// index:0
// Public virtual
// Qt::DropActions supportedDropActions()
func (this *QStringListModel) SupportedDropActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QStringListModel20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
