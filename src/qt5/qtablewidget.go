package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qtablewidget.h
// dst-file: /src/widgets/qtablewidget.go
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
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange(int top, int left, int bottom, int right);
extern void* dector_ZN26QTableWidgetSelectionRangeC1Eiiii(int arg0, int arg1, int arg2, int arg3);
extern void _ZN26QTableWidgetSelectionRangeC1Eiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  int QTableWidgetSelectionRange::columnCount();
extern void demth_ZNK26QTableWidgetSelectionRange11columnCountEv(void* qthis);
  // proto:  int QTableWidgetSelectionRange::rowCount();
extern void demth_ZNK26QTableWidgetSelectionRange8rowCountEv(void* qthis);
  // proto:  int QTableWidgetSelectionRange::leftColumn();
extern void demth_ZNK26QTableWidgetSelectionRange10leftColumnEv(void* qthis);
  // proto:  void QTableWidgetSelectionRange::~QTableWidgetSelectionRange();
extern void _ZN26QTableWidgetSelectionRangeD0Ev(void* qthis);
  // proto:  int QTableWidgetSelectionRange::topRow();
extern void demth_ZNK26QTableWidgetSelectionRange6topRowEv(void* qthis);
  // proto:  int QTableWidgetSelectionRange::rightColumn();
extern void demth_ZNK26QTableWidgetSelectionRange11rightColumnEv(void* qthis);
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange();
extern void* dector_ZN26QTableWidgetSelectionRangeC1Ev();
extern void _ZN26QTableWidgetSelectionRangeC1Ev(void* qthis);
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange(const QTableWidgetSelectionRange & other);
extern void* dector_ZN26QTableWidgetSelectionRangeC1ERKS_(void* arg0);
extern void _ZN26QTableWidgetSelectionRangeC1ERKS_(void* qthis, void* arg0);
  // proto:  int QTableWidgetSelectionRange::bottomRow();
extern void demth_ZNK26QTableWidgetSelectionRange9bottomRowEv(void* qthis);
  // proto:  void QTableWidget::setColumnCount(int columns);
extern void _ZN12QTableWidget14setColumnCountEi(void* qthis, int arg0);
  // proto:  void QTableWidget::~QTableWidget();
extern void _ZN12QTableWidgetD0Ev(void* qthis);
  // proto:  QList<QTableWidgetItem *> QTableWidget::selectedItems();
extern void _ZNK12QTableWidget13selectedItemsEv(void* qthis);
  // proto:  bool QTableWidget::isSortingEnabled();
extern void _ZNK12QTableWidget16isSortingEnabledEv(void* qthis);
  // proto:  const QMetaObject * QTableWidget::metaObject();
extern void _ZNK12QTableWidget10metaObjectEv(void* qthis);
  // proto:  void QTableWidget::QTableWidget(const QTableWidget & );
extern void* dector_ZN12QTableWidgetC1ERKS_(void* arg0);
extern void _ZN12QTableWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTableWidget::closePersistentEditor(QTableWidgetItem * item);
extern void _ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  void QTableWidget::setHorizontalHeaderLabels(const QStringList & labels);
extern void _ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList(void* qthis, void* arg0);
  // proto:  void QTableWidget::setItemSelected(const QTableWidgetItem * item, bool select);
extern void _ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb(void* qthis, void* arg0, bool arg1);
  // proto:  QTableWidgetItem * QTableWidget::takeItem(int row, int column);
extern void _ZN12QTableWidget8takeItemEii(void* qthis, int arg0, int arg1);
  // proto:  void QTableWidget::removeCellWidget(int row, int column);
extern void demth_ZN12QTableWidget16removeCellWidgetEii(void* qthis, int arg0, int arg1);
  // proto:  void QTableWidget::setVerticalHeaderItem(int row, QTableWidgetItem * item);
extern void _ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem(void* qthis, int arg0, void* arg1);
  // proto:  QRect QTableWidget::visualItemRect(const QTableWidgetItem * item);
extern void _ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  QTableWidgetItem * QTableWidget::currentItem();
extern void _ZNK12QTableWidget11currentItemEv(void* qthis);
  // proto:  int QTableWidget::row(const QTableWidgetItem * item);
extern void _ZNK12QTableWidget3rowEPK16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  void QTableWidget::removeRow(int row);
extern void _ZN12QTableWidget9removeRowEi(void* qthis, int arg0);
  // proto:  void QTableWidget::setItemPrototype(const QTableWidgetItem * item);
extern void _ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  void QTableWidget::QTableWidget(int rows, int columns, QWidget * parent);
extern void* dector_ZN12QTableWidgetC1EiiP7QWidget(int arg0, int arg1, void* arg2);
extern void _ZN12QTableWidgetC1EiiP7QWidget(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  int QTableWidget::visualRow(int logicalRow);
extern void _ZNK12QTableWidget9visualRowEi(void* qthis, int arg0);
  // proto:  void QTableWidget::setCellWidget(int row, int column, QWidget * widget);
extern void _ZN12QTableWidget13setCellWidgetEiiP7QWidget(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QTableWidget::openPersistentEditor(QTableWidgetItem * item);
extern void _ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  int QTableWidget::columnCount();
extern void _ZNK12QTableWidget11columnCountEv(void* qthis);
  // proto:  int QTableWidget::currentRow();
extern void _ZNK12QTableWidget10currentRowEv(void* qthis);
  // proto:  void QTableWidget::setCurrentItem(QTableWidgetItem * item);
extern void _ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  QWidget * QTableWidget::cellWidget(int row, int column);
extern void _ZNK12QTableWidget10cellWidgetEii(void* qthis, int arg0, int arg1);
  // proto:  void QTableWidget::setSortingEnabled(bool enable);
extern void _ZN12QTableWidget17setSortingEnabledEb(void* qthis, bool arg0);
  // proto:  void QTableWidget::setItem(int row, int column, QTableWidgetItem * item);
extern void _ZN12QTableWidget7setItemEiiP16QTableWidgetItem(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QTableWidgetItem * QTableWidget::horizontalHeaderItem(int column);
extern void _ZNK12QTableWidget20horizontalHeaderItemEi(void* qthis, int arg0);
  // proto:  void QTableWidget::editItem(QTableWidgetItem * item);
extern void _ZN12QTableWidget8editItemEP16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  QList<QTableWidgetSelectionRange> QTableWidget::selectedRanges();
extern void _ZNK12QTableWidget14selectedRangesEv(void* qthis);
  // proto:  int QTableWidget::currentColumn();
extern void _ZNK12QTableWidget13currentColumnEv(void* qthis);
  // proto:  void QTableWidget::removeColumn(int column);
extern void _ZN12QTableWidget12removeColumnEi(void* qthis, int arg0);
  // proto:  void QTableWidget::setRangeSelected(const QTableWidgetSelectionRange & range, bool select);
extern void _ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb(void* qthis, void* arg0, bool arg1);
  // proto:  int QTableWidget::column(const QTableWidgetItem * item);
extern void _ZNK12QTableWidget6columnEPK16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  bool QTableWidget::isItemSelected(const QTableWidgetItem * item);
extern void _ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem(void* qthis, void* arg0);
  // proto:  QTableWidgetItem * QTableWidget::takeVerticalHeaderItem(int row);
extern void _ZN12QTableWidget22takeVerticalHeaderItemEi(void* qthis, int arg0);
  // proto:  void QTableWidget::insertRow(int row);
extern void _ZN12QTableWidget9insertRowEi(void* qthis, int arg0);
  // proto:  int QTableWidget::rowCount();
extern void _ZNK12QTableWidget8rowCountEv(void* qthis);
  // proto:  QTableWidgetItem * QTableWidget::item(int row, int column);
extern void _ZNK12QTableWidget4itemEii(void* qthis, int arg0, int arg1);
  // proto:  void QTableWidget::QTableWidget(QWidget * parent);
extern void* dector_ZN12QTableWidgetC1EP7QWidget(void* arg0);
extern void _ZN12QTableWidgetC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QTableWidget::setVerticalHeaderLabels(const QStringList & labels);
extern void _ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList(void* qthis, void* arg0);
  // proto:  const QTableWidgetItem * QTableWidget::itemPrototype();
extern void _ZNK12QTableWidget13itemPrototypeEv(void* qthis);
  // proto:  QTableWidgetItem * QTableWidget::itemAt(const QPoint & p);
extern void _ZNK12QTableWidget6itemAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QTableWidget::clearContents();
extern void _ZN12QTableWidget13clearContentsEv(void* qthis);
  // proto:  QTableWidgetItem * QTableWidget::itemAt(int x, int y);
extern void demth_ZNK12QTableWidget6itemAtEii(void* qthis, int arg0, int arg1);
  // proto:  void QTableWidget::setCurrentCell(int row, int column);
extern void _ZN12QTableWidget14setCurrentCellEii(void* qthis, int arg0, int arg1);
  // proto:  void QTableWidget::setRowCount(int rows);
extern void _ZN12QTableWidget11setRowCountEi(void* qthis, int arg0);
  // proto:  void QTableWidget::setHorizontalHeaderItem(int column, QTableWidgetItem * item);
extern void _ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem(void* qthis, int arg0, void* arg1);
  // proto:  int QTableWidget::visualColumn(int logicalColumn);
extern void _ZNK12QTableWidget12visualColumnEi(void* qthis, int arg0);
  // proto:  QTableWidgetItem * QTableWidget::takeHorizontalHeaderItem(int column);
extern void _ZN12QTableWidget24takeHorizontalHeaderItemEi(void* qthis, int arg0);
  // proto:  QTableWidgetItem * QTableWidget::verticalHeaderItem(int row);
extern void _ZNK12QTableWidget18verticalHeaderItemEi(void* qthis, int arg0);
  // proto:  void QTableWidget::clear();
extern void _ZN12QTableWidget5clearEv(void* qthis);
  // proto:  void QTableWidget::insertColumn(int column);
extern void _ZN12QTableWidget12insertColumnEi(void* qthis, int arg0);
  // proto:  QColor QTableWidgetItem::backgroundColor();
extern void demth_ZNK16QTableWidgetItem15backgroundColorEv(void* qthis);
  // proto:  QVariant QTableWidgetItem::data(int role);
extern void _ZNK16QTableWidgetItem4dataEi(void* qthis, int arg0);
  // proto:  void QTableWidgetItem::setSelected(bool select);
extern void demth_ZN16QTableWidgetItem11setSelectedEb(void* qthis, bool arg0);
  // proto:  void QTableWidgetItem::setStatusTip(const QString & statusTip);
extern void demth_ZN16QTableWidgetItem12setStatusTipERK7QString(void* qthis, void* arg0);
  // proto:  QColor QTableWidgetItem::textColor();
extern void demth_ZNK16QTableWidgetItem9textColorEv(void* qthis);
  // proto:  void QTableWidgetItem::~QTableWidgetItem();
extern void _ZN16QTableWidgetItemD0Ev(void* qthis);
  // proto:  QString QTableWidgetItem::text();
extern void demth_ZNK16QTableWidgetItem4textEv(void* qthis);
  // proto:  void QTableWidgetItem::setSizeHint(const QSize & size);
extern void demth_ZN16QTableWidgetItem11setSizeHintERK5QSize(void* qthis, void* arg0);
  // proto:  QBrush QTableWidgetItem::foreground();
extern void demth_ZNK16QTableWidgetItem10foregroundEv(void* qthis);
  // proto:  int QTableWidgetItem::type();
extern void demth_ZNK16QTableWidgetItem4typeEv(void* qthis);
  // proto:  int QTableWidgetItem::column();
extern void demth_ZNK16QTableWidgetItem6columnEv(void* qthis);
  // proto:  void QTableWidgetItem::setTextAlignment(int alignment);
extern void demth_ZN16QTableWidgetItem16setTextAlignmentEi(void* qthis, int arg0);
  // proto:  QFont QTableWidgetItem::font();
extern void demth_ZNK16QTableWidgetItem4fontEv(void* qthis);
  // proto:  QIcon QTableWidgetItem::icon();
extern void demth_ZNK16QTableWidgetItem4iconEv(void* qthis);
  // proto:  void QTableWidgetItem::write(QDataStream & out);
extern void _ZNK16QTableWidgetItem5writeER11QDataStream(void* qthis, void* arg0);
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QTableWidgetItem & other);
extern void* dector_ZN16QTableWidgetItemC1ERKS_(void* arg0);
extern void _ZN16QTableWidgetItemC1ERKS_(void* qthis, void* arg0);
  // proto:  QBrush QTableWidgetItem::background();
extern void demth_ZNK16QTableWidgetItem10backgroundEv(void* qthis);
  // proto:  void QTableWidgetItem::setIcon(const QIcon & icon);
extern void demth_ZN16QTableWidgetItem7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QString & text, int type);
extern void* dector_ZN16QTableWidgetItemC1ERK7QStringi(void* arg0, int arg1);
extern void _ZN16QTableWidgetItemC1ERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  QString QTableWidgetItem::statusTip();
extern void demth_ZNK16QTableWidgetItem9statusTipEv(void* qthis);
  // proto:  QTableWidgetItem * QTableWidgetItem::clone();
extern void _ZNK16QTableWidgetItem5cloneEv(void* qthis);
  // proto:  void QTableWidgetItem::QTableWidgetItem(int type);
extern void* dector_ZN16QTableWidgetItemC1Ei(int arg0);
extern void _ZN16QTableWidgetItemC1Ei(void* qthis, int arg0);
  // proto:  void QTableWidgetItem::setWhatsThis(const QString & whatsThis);
extern void demth_ZN16QTableWidgetItem12setWhatsThisERK7QString(void* qthis, void* arg0);
  // proto:  QSize QTableWidgetItem::sizeHint();
extern void demth_ZNK16QTableWidgetItem8sizeHintEv(void* qthis);
  // proto:  void QTableWidgetItem::setForeground(const QBrush & brush);
extern void demth_ZN16QTableWidgetItem13setForegroundERK6QBrush(void* qthis, void* arg0);
  // proto:  int QTableWidgetItem::row();
extern void demth_ZNK16QTableWidgetItem3rowEv(void* qthis);
  // proto:  void QTableWidgetItem::setData(int role, const QVariant & value);
extern void _ZN16QTableWidgetItem7setDataEiRK8QVariant(void* qthis, int arg0, void* arg1);
  // proto:  QTableWidget * QTableWidgetItem::tableWidget();
extern void demth_ZNK16QTableWidgetItem11tableWidgetEv(void* qthis);
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QIcon & icon, const QString & text, int type);
extern void* dector_ZN16QTableWidgetItemC1ERK5QIconRK7QStringi(void* arg0, void* arg1, int arg2);
extern void _ZN16QTableWidgetItemC1ERK5QIconRK7QStringi(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  int QTableWidgetItem::textAlignment();
extern void demth_ZNK16QTableWidgetItem13textAlignmentEv(void* qthis);
  // proto:  void QTableWidgetItem::read(QDataStream & in);
extern void _ZN16QTableWidgetItem4readER11QDataStream(void* qthis, void* arg0);
  // proto:  QString QTableWidgetItem::toolTip();
extern void demth_ZNK16QTableWidgetItem7toolTipEv(void* qthis);
  // proto:  bool QTableWidgetItem::isSelected();
extern void demth_ZNK16QTableWidgetItem10isSelectedEv(void* qthis);
  // proto:  void QTableWidgetItem::setBackgroundColor(const QColor & color);
extern void demth_ZN16QTableWidgetItem18setBackgroundColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QTableWidgetItem::setBackground(const QBrush & brush);
extern void demth_ZN16QTableWidgetItem13setBackgroundERK6QBrush(void* qthis, void* arg0);
  // proto:  void QTableWidgetItem::setFont(const QFont & font);
extern void demth_ZN16QTableWidgetItem7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  void QTableWidgetItem::setTextColor(const QColor & color);
extern void demth_ZN16QTableWidgetItem12setTextColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QTableWidgetItem::setText(const QString & text);
extern void demth_ZN16QTableWidgetItem7setTextERK7QString(void* qthis, void* arg0);
  // proto:  QString QTableWidgetItem::whatsThis();
extern void demth_ZNK16QTableWidgetItem9whatsThisEv(void* qthis);
  // proto:  void QTableWidgetItem::setToolTip(const QString & toolTip);
extern void demth_ZN16QTableWidgetItem10setToolTipERK7QString(void* qthis, void* arg0);
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

// class sizeof(QTableWidgetSelectionRange)=16
type QTableWidgetSelectionRange struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTableWidget)=1
type QTableWidget struct {
  /*qbase*/ QTableView;
  qclsinst uint64 /* *mut c_void*/;
//  _itemDoubleClicked QTableWidget_itemDoubleClicked_signal;
//  _cellEntered QTableWidget_cellEntered_signal;
//  _itemClicked QTableWidget_itemClicked_signal;
//  _currentItemChanged QTableWidget_currentItemChanged_signal;
//  _itemEntered QTableWidget_itemEntered_signal;
//  _itemPressed QTableWidget_itemPressed_signal;
//  _cellClicked QTableWidget_cellClicked_signal;
//  _itemSelectionChanged QTableWidget_itemSelectionChanged_signal;
//  _cellChanged QTableWidget_cellChanged_signal;
//  _itemActivated QTableWidget_itemActivated_signal;
//  _cellActivated QTableWidget_cellActivated_signal;
//  _itemChanged QTableWidget_itemChanged_signal;
//  _currentCellChanged QTableWidget_currentCellChanged_signal;
//  _cellDoubleClicked QTableWidget_cellDoubleClicked_signal;
//  _cellPressed QTableWidget_cellPressed_signal;
}

// class sizeof(QTableWidgetItem)=1
type QTableWidgetItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange(int top, int left, int bottom, int right);
func NewQTableWidgetSelectionRange(args ...interface{}) QTableWidgetSelectionRange {
  return QTableWidgetSelectionRange{}
}

  // proto:  int QTableWidgetSelectionRange::columnCount();
func (this *QTableWidgetSelectionRange) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange11columnCountEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "columnCount", args)
  }

}

  // proto:  int QTableWidgetSelectionRange::rowCount();
func (this *QTableWidgetSelectionRange) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange8rowCountEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rowCount", args)
  }

}

  // proto:  int QTableWidgetSelectionRange::leftColumn();
func (this *QTableWidgetSelectionRange) leftColumn(args ...interface{}) () {
  // leftColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange10leftColumnEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "leftColumn", args)
  }

}

  // proto:  void QTableWidgetSelectionRange::~QTableWidgetSelectionRange();
func (this *QTableWidgetSelectionRange) FreeQTableWidgetSelectionRange(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "~QTableWidgetSelectionRange", args)
  }

}

  // proto:  int QTableWidgetSelectionRange::topRow();
func (this *QTableWidgetSelectionRange) topRow(args ...interface{}) () {
  // topRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange6topRowEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "topRow", args)
  }

}

  // proto:  int QTableWidgetSelectionRange::rightColumn();
func (this *QTableWidgetSelectionRange) rightColumn(args ...interface{}) () {
  // rightColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange11rightColumnEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rightColumn", args)
  }

}

  // proto:  int QTableWidgetSelectionRange::bottomRow();
func (this *QTableWidgetSelectionRange) bottomRow(args ...interface{}) () {
  // bottomRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange9bottomRowEv
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "bottomRow", args)
  }

}

  // proto:  void QTableWidget::setColumnCount(int columns);
func (this *QTableWidget) setColumnCount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setColumnCountEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setColumnCount", args)
  }

}

  // proto:  void QTableWidget::~QTableWidget();
func (this *QTableWidget) FreeQTableWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidget", "~QTableWidget", args)
  }

}

  // proto:  QList<QTableWidgetItem *> QTableWidget::selectedItems();
func (this *QTableWidget) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13selectedItemsEv
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedItems", args)
  }

}

  // proto:  bool QTableWidget::isSortingEnabled();
func (this *QTableWidget) isSortingEnabled(args ...interface{}) () {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget16isSortingEnabledEv
  default:
    qtrt.ErrorResolve("QTableWidget", "isSortingEnabled", args)
  }

}

  // proto:  const QMetaObject * QTableWidget::metaObject();
func (this *QTableWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QTableWidget", "metaObject", args)
  }

}

  // proto:  void QTableWidget::QTableWidget(const QTableWidget & );
func NewQTableWidget(args ...interface{}) QTableWidget {
  return QTableWidget{}
}

  // proto:  void QTableWidget::closePersistentEditor(QTableWidgetItem * item);
func (this *QTableWidget) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "closePersistentEditor", args)
  }

}

  // proto:  void QTableWidget::setHorizontalHeaderLabels(const QStringList & labels);
func (this *QTableWidget) setHorizontalHeaderLabels(args ...interface{}) () {
  // setHorizontalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderLabels", args)
  }

}

  // proto:  void QTableWidget::setItemSelected(const QTableWidgetItem * item, bool select);
func (this *QTableWidget) setItemSelected(args ...interface{}) () {
  // setItemSelected(const class QTableWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemSelected", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::takeItem(int row, int column);
func (this *QTableWidget) takeItem(args ...interface{}) () {
  // takeItem(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget8takeItemEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "takeItem", args)
  }

}

  // proto:  void QTableWidget::removeCellWidget(int row, int column);
func (this *QTableWidget) removeCellWidget(args ...interface{}) () {
  // removeCellWidget(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16removeCellWidgetEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "removeCellWidget", args)
  }

}

  // proto:  void QTableWidget::setVerticalHeaderItem(int row, QTableWidgetItem * item);
func (this *QTableWidget) setVerticalHeaderItem(args ...interface{}) () {
  // setVerticalHeaderItem(int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderItem", args)
  }

}

  // proto:  QRect QTableWidget::visualItemRect(const QTableWidgetItem * item);
func (this *QTableWidget) visualItemRect(args ...interface{}) () {
  // visualItemRect(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "visualItemRect", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::currentItem();
func (this *QTableWidget) currentItem(args ...interface{}) () {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget11currentItemEv
  default:
    qtrt.ErrorResolve("QTableWidget", "currentItem", args)
  }

}

  // proto:  int QTableWidget::row(const QTableWidgetItem * item);
func (this *QTableWidget) row(args ...interface{}) () {
  // row(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget3rowEPK16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "row", args)
  }

}

  // proto:  void QTableWidget::removeRow(int row);
func (this *QTableWidget) removeRow(args ...interface{}) () {
  // removeRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget9removeRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "removeRow", args)
  }

}

  // proto:  void QTableWidget::setItemPrototype(const QTableWidgetItem * item);
func (this *QTableWidget) setItemPrototype(args ...interface{}) () {
  // setItemPrototype(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemPrototype", args)
  }

}

  // proto:  int QTableWidget::visualRow(int logicalRow);
func (this *QTableWidget) visualRow(args ...interface{}) () {
  // visualRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget9visualRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "visualRow", args)
  }

}

  // proto:  void QTableWidget::setCellWidget(int row, int column, QWidget * widget);
func (this *QTableWidget) setCellWidget(args ...interface{}) () {
  // setCellWidget(int, int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget13setCellWidgetEiiP7QWidget
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setCellWidget", args)
  }

}

  // proto:  void QTableWidget::openPersistentEditor(QTableWidgetItem * item);
func (this *QTableWidget) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "openPersistentEditor", args)
  }

}

  // proto:  int QTableWidget::columnCount();
func (this *QTableWidget) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget11columnCountEv
  default:
    qtrt.ErrorResolve("QTableWidget", "columnCount", args)
  }

}

  // proto:  int QTableWidget::currentRow();
func (this *QTableWidget) currentRow(args ...interface{}) () {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10currentRowEv
  default:
    qtrt.ErrorResolve("QTableWidget", "currentRow", args)
  }

}

  // proto:  void QTableWidget::setCurrentItem(QTableWidgetItem * item);
func (this *QTableWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QTableWidgetItem *, class QItemSelectionModel::SelectionFlags)
  // setCurrentItem(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
  vtys[0][1] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentItem", args)
  }

}

  // proto:  QWidget * QTableWidget::cellWidget(int row, int column);
func (this *QTableWidget) cellWidget(args ...interface{}) () {
  // cellWidget(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10cellWidgetEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "cellWidget", args)
  }

}

  // proto:  void QTableWidget::setSortingEnabled(bool enable);
func (this *QTableWidget) setSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget17setSortingEnabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setSortingEnabled", args)
  }

}

  // proto:  void QTableWidget::setItem(int row, int column, QTableWidgetItem * item);
func (this *QTableWidget) setItem(args ...interface{}) () {
  // setItem(int, int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget7setItemEiiP16QTableWidgetItem
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setItem", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::horizontalHeaderItem(int column);
func (this *QTableWidget) horizontalHeaderItem(args ...interface{}) () {
  // horizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget20horizontalHeaderItemEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "horizontalHeaderItem", args)
  }

}

  // proto:  void QTableWidget::editItem(QTableWidgetItem * item);
func (this *QTableWidget) editItem(args ...interface{}) () {
  // editItem(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget8editItemEP16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "editItem", args)
  }

}

  // proto:  QList<QTableWidgetSelectionRange> QTableWidget::selectedRanges();
func (this *QTableWidget) selectedRanges(args ...interface{}) () {
  // selectedRanges()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14selectedRangesEv
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedRanges", args)
  }

}

  // proto:  int QTableWidget::currentColumn();
func (this *QTableWidget) currentColumn(args ...interface{}) () {
  // currentColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13currentColumnEv
  default:
    qtrt.ErrorResolve("QTableWidget", "currentColumn", args)
  }

}

  // proto:  void QTableWidget::removeColumn(int column);
func (this *QTableWidget) removeColumn(args ...interface{}) () {
  // removeColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget12removeColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "removeColumn", args)
  }

}

  // proto:  void QTableWidget::setRangeSelected(const QTableWidgetSelectionRange & range, bool select);
func (this *QTableWidget) setRangeSelected(args ...interface{}) () {
  // setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetSelectionRange{}) // "const QTableWidgetSelectionRange &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb
    var arg0 = args[0].(QTableWidgetSelectionRange).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setRangeSelected", args)
  }

}

  // proto:  int QTableWidget::column(const QTableWidgetItem * item);
func (this *QTableWidget) column(args ...interface{}) () {
  // column(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget6columnEPK16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "column", args)
  }

}

  // proto:  bool QTableWidget::isItemSelected(const QTableWidgetItem * item);
func (this *QTableWidget) isItemSelected(args ...interface{}) () {
  // isItemSelected(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "isItemSelected", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::takeVerticalHeaderItem(int row);
func (this *QTableWidget) takeVerticalHeaderItem(args ...interface{}) () {
  // takeVerticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget22takeVerticalHeaderItemEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "takeVerticalHeaderItem", args)
  }

}

  // proto:  void QTableWidget::insertRow(int row);
func (this *QTableWidget) insertRow(args ...interface{}) () {
  // insertRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget9insertRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "insertRow", args)
  }

}

  // proto:  int QTableWidget::rowCount();
func (this *QTableWidget) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget8rowCountEv
  default:
    qtrt.ErrorResolve("QTableWidget", "rowCount", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::item(int row, int column);
func (this *QTableWidget) item(args ...interface{}) () {
  // item(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget4itemEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "item", args)
  }

}

  // proto:  void QTableWidget::setVerticalHeaderLabels(const QStringList & labels);
func (this *QTableWidget) setVerticalHeaderLabels(args ...interface{}) () {
  // setVerticalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderLabels", args)
  }

}

  // proto:  const QTableWidgetItem * QTableWidget::itemPrototype();
func (this *QTableWidget) itemPrototype(args ...interface{}) () {
  // itemPrototype()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13itemPrototypeEv
  default:
    qtrt.ErrorResolve("QTableWidget", "itemPrototype", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::itemAt(const QPoint & p);
func (this *QTableWidget) itemAt(args ...interface{}) () {
  // itemAt(const class QPoint &)
  // itemAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget6itemAtERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK12QTableWidget6itemAtEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "itemAt", args)
  }

}

  // proto:  void QTableWidget::clearContents();
func (this *QTableWidget) clearContents(args ...interface{}) () {
  // clearContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget13clearContentsEv
  default:
    qtrt.ErrorResolve("QTableWidget", "clearContents", args)
  }

}

  // proto:  void QTableWidget::setCurrentCell(int row, int column);
func (this *QTableWidget) setCurrentCell(args ...interface{}) () {
  // setCurrentCell(int, int, class QItemSelectionModel::SelectionFlags)
  // setCurrentCell(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setCurrentCellEii6QFlagsIN19QItemSelectionModel13SelectionFlagEE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN12QTableWidget14setCurrentCellEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentCell", args)
  }

}

  // proto:  void QTableWidget::setRowCount(int rows);
func (this *QTableWidget) setRowCount(args ...interface{}) () {
  // setRowCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget11setRowCountEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setRowCount", args)
  }

}

  // proto:  void QTableWidget::setHorizontalHeaderItem(int column, QTableWidgetItem * item);
func (this *QTableWidget) setHorizontalHeaderItem(args ...interface{}) () {
  // setHorizontalHeaderItem(int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderItem", args)
  }

}

  // proto:  int QTableWidget::visualColumn(int logicalColumn);
func (this *QTableWidget) visualColumn(args ...interface{}) () {
  // visualColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget12visualColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "visualColumn", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::takeHorizontalHeaderItem(int column);
func (this *QTableWidget) takeHorizontalHeaderItem(args ...interface{}) () {
  // takeHorizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget24takeHorizontalHeaderItemEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "takeHorizontalHeaderItem", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidget::verticalHeaderItem(int row);
func (this *QTableWidget) verticalHeaderItem(args ...interface{}) () {
  // verticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget18verticalHeaderItemEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "verticalHeaderItem", args)
  }

}

  // proto:  void QTableWidget::clear();
func (this *QTableWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget5clearEv
  default:
    qtrt.ErrorResolve("QTableWidget", "clear", args)
  }

}

  // proto:  void QTableWidget::insertColumn(int column);
func (this *QTableWidget) insertColumn(args ...interface{}) () {
  // insertColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget12insertColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidget", "insertColumn", args)
  }

}

  // proto:  QColor QTableWidgetItem::backgroundColor();
func (this *QTableWidgetItem) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem15backgroundColorEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "backgroundColor", args)
  }

}

  // proto:  QVariant QTableWidgetItem::data(int role);
func (this *QTableWidgetItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4dataEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "data", args)
  }

}

  // proto:  void QTableWidgetItem::setSelected(bool select);
func (this *QTableWidgetItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem11setSelectedEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSelected", args)
  }

}

  // proto:  void QTableWidgetItem::setStatusTip(const QString & statusTip);
func (this *QTableWidgetItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setStatusTipERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setStatusTip", args)
  }

}

  // proto:  QColor QTableWidgetItem::textColor();
func (this *QTableWidgetItem) textColor(args ...interface{}) () {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9textColorEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textColor", args)
  }

}

  // proto:  void QTableWidgetItem::~QTableWidgetItem();
func (this *QTableWidgetItem) FreeQTableWidgetItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "~QTableWidgetItem", args)
  }

}

  // proto:  QString QTableWidgetItem::text();
func (this *QTableWidgetItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4textEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "text", args)
  }

}

  // proto:  void QTableWidgetItem::setSizeHint(const QSize & size);
func (this *QTableWidgetItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem11setSizeHintERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSizeHint", args)
  }

}

  // proto:  QBrush QTableWidgetItem::foreground();
func (this *QTableWidgetItem) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10foregroundEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "foreground", args)
  }

}

  // proto:  int QTableWidgetItem::type();
func (this *QTableWidgetItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "type", args)
  }

}

  // proto:  int QTableWidgetItem::column();
func (this *QTableWidgetItem) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem6columnEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "column", args)
  }

}

  // proto:  void QTableWidgetItem::setTextAlignment(int alignment);
func (this *QTableWidgetItem) setTextAlignment(args ...interface{}) () {
  // setTextAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem16setTextAlignmentEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextAlignment", args)
  }

}

  // proto:  QFont QTableWidgetItem::font();
func (this *QTableWidgetItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4fontEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "font", args)
  }

}

  // proto:  QIcon QTableWidgetItem::icon();
func (this *QTableWidgetItem) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4iconEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "icon", args)
  }

}

  // proto:  void QTableWidgetItem::write(QDataStream & out);
func (this *QTableWidgetItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5writeER11QDataStream
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "write", args)
  }

}

  // proto:  void QTableWidgetItem::QTableWidgetItem(const QTableWidgetItem & other);
func NewQTableWidgetItem(args ...interface{}) QTableWidgetItem {
  return QTableWidgetItem{}
}

  // proto:  QBrush QTableWidgetItem::background();
func (this *QTableWidgetItem) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10backgroundEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "background", args)
  }

}

  // proto:  void QTableWidgetItem::setIcon(const QIcon & icon);
func (this *QTableWidgetItem) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setIconERK5QIcon
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setIcon", args)
  }

}

  // proto:  QString QTableWidgetItem::statusTip();
func (this *QTableWidgetItem) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9statusTipEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "statusTip", args)
  }

}

  // proto:  QTableWidgetItem * QTableWidgetItem::clone();
func (this *QTableWidgetItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5cloneEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "clone", args)
  }

}

  // proto:  void QTableWidgetItem::setWhatsThis(const QString & whatsThis);
func (this *QTableWidgetItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setWhatsThisERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setWhatsThis", args)
  }

}

  // proto:  QSize QTableWidgetItem::sizeHint();
func (this *QTableWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem8sizeHintEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "sizeHint", args)
  }

}

  // proto:  void QTableWidgetItem::setForeground(const QBrush & brush);
func (this *QTableWidgetItem) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setForegroundERK6QBrush
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setForeground", args)
  }

}

  // proto:  int QTableWidgetItem::row();
func (this *QTableWidgetItem) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem3rowEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "row", args)
  }

}

  // proto:  void QTableWidgetItem::setData(int role, const QVariant & value);
func (this *QTableWidgetItem) setData(args ...interface{}) () {
  // setData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setDataEiRK8QVariant
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setData", args)
  }

}

  // proto:  QTableWidget * QTableWidgetItem::tableWidget();
func (this *QTableWidgetItem) tableWidget(args ...interface{}) () {
  // tableWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem11tableWidgetEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "tableWidget", args)
  }

}

  // proto:  int QTableWidgetItem::textAlignment();
func (this *QTableWidgetItem) textAlignment(args ...interface{}) () {
  // textAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem13textAlignmentEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textAlignment", args)
  }

}

  // proto:  void QTableWidgetItem::read(QDataStream & in);
func (this *QTableWidgetItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem4readER11QDataStream
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "read", args)
  }

}

  // proto:  QString QTableWidgetItem::toolTip();
func (this *QTableWidgetItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem7toolTipEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "toolTip", args)
  }

}

  // proto:  bool QTableWidgetItem::isSelected();
func (this *QTableWidgetItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10isSelectedEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "isSelected", args)
  }

}

  // proto:  void QTableWidgetItem::setBackgroundColor(const QColor & color);
func (this *QTableWidgetItem) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem18setBackgroundColorERK6QColor
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackgroundColor", args)
  }

}

  // proto:  void QTableWidgetItem::setBackground(const QBrush & brush);
func (this *QTableWidgetItem) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setBackgroundERK6QBrush
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackground", args)
  }

}

  // proto:  void QTableWidgetItem::setFont(const QFont & font);
func (this *QTableWidgetItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setFontERK5QFont
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setFont", args)
  }

}

  // proto:  void QTableWidgetItem::setTextColor(const QColor & color);
func (this *QTableWidgetItem) setTextColor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setTextColorERK6QColor
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextColor", args)
  }

}

  // proto:  void QTableWidgetItem::setText(const QString & text);
func (this *QTableWidgetItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setTextERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setText", args)
  }

}

  // proto:  QString QTableWidgetItem::whatsThis();
func (this *QTableWidgetItem) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9whatsThisEv
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "whatsThis", args)
  }

}

  // proto:  void QTableWidgetItem::setToolTip(const QString & toolTip);
func (this *QTableWidgetItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem10setToolTipERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setToolTip", args)
  }

}

// <= body block end

