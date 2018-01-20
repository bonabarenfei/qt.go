//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QAccessibleTableInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleTableInterface) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQAccessibleTableInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleTableInterface {
	return &QAccessibleTableInterface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessible.h:595
// index:0
// Public virtual
// void ~QAccessibleTableInterface()
func DeleteQAccessibleTableInterface(*QAccessibleTableInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:597
// index:0
// Public pure virtual
// QAccessibleInterface * caption()
func (this *QAccessibleTableInterface) Caption() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface7captionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:598
// index:0
// Public pure virtual
// QAccessibleInterface * summary()
func (this *QAccessibleTableInterface) Summary() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface7summaryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:600
// index:0
// Public pure virtual
// QAccessibleInterface * cellAt(int, int)
func (this *QAccessibleTableInterface) CellAt(row int, column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface6cellAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:601
// index:0
// Public pure virtual
// int selectedCellCount()
func (this *QAccessibleTableInterface) SelectedCellCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface17selectedCellCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:602
// index:0
// Public pure virtual
// QList<QAccessibleInterface *> selectedCells()
func (this *QAccessibleTableInterface) SelectedCells() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface13selectedCellsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:604
// index:0
// Public pure virtual
// QString columnDescription(int)
func (this *QAccessibleTableInterface) ColumnDescription(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface17columnDescriptionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:605
// index:0
// Public pure virtual
// QString rowDescription(int)
func (this *QAccessibleTableInterface) RowDescription(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface14rowDescriptionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:606
// index:0
// Public pure virtual
// int selectedColumnCount()
func (this *QAccessibleTableInterface) SelectedColumnCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface19selectedColumnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:607
// index:0
// Public pure virtual
// int selectedRowCount()
func (this *QAccessibleTableInterface) SelectedRowCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface16selectedRowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:608
// index:0
// Public pure virtual
// int columnCount()
func (this *QAccessibleTableInterface) ColumnCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:609
// index:0
// Public pure virtual
// int rowCount()
func (this *QAccessibleTableInterface) RowCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:610
// index:0
// Public pure virtual
// QList<int> selectedColumns()
func (this *QAccessibleTableInterface) SelectedColumns() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface15selectedColumnsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:611
// index:0
// Public pure virtual
// QList<int> selectedRows()
func (this *QAccessibleTableInterface) SelectedRows() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface12selectedRowsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:612
// index:0
// Public pure virtual
// bool isColumnSelected(int)
func (this *QAccessibleTableInterface) IsColumnSelected(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface16isColumnSelectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:613
// index:0
// Public pure virtual
// bool isRowSelected(int)
func (this *QAccessibleTableInterface) IsRowSelected(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface13isRowSelectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:614
// index:0
// Public pure virtual
// bool selectRow(int)
func (this *QAccessibleTableInterface) SelectRow(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface9selectRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:615
// index:0
// Public pure virtual
// bool selectColumn(int)
func (this *QAccessibleTableInterface) SelectColumn(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface12selectColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:616
// index:0
// Public pure virtual
// bool unselectRow(int)
func (this *QAccessibleTableInterface) UnselectRow(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface11unselectRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:617
// index:0
// Public pure virtual
// bool unselectColumn(int)
func (this *QAccessibleTableInterface) UnselectColumn(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface14unselectColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:619
// index:0
// Public pure virtual
// void modelChange(class QAccessibleTableModelChangeEvent *)
func (this *QAccessibleTableInterface) ModelChange(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface11modelChangeEP32QAccessibleTableModelChangeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
