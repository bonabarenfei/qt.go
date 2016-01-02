package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qtabbar.h
// dst-file: /src/widgets/qtabbar.go
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
  // proto:  bool QTabBar::usesScrollButtons();
extern void _ZNK7QTabBar17usesScrollButtonsEv(void* qthis);
  // proto:  bool QTabBar::autoHide();
extern void _ZNK7QTabBar8autoHideEv(void* qthis);
  // proto:  QString QTabBar::tabToolTip(int index);
extern void _ZNK7QTabBar10tabToolTipEi(void* qthis, int arg0);
  // proto:  bool QTabBar::expanding();
extern void _ZNK7QTabBar9expandingEv(void* qthis);
  // proto:  void QTabBar::setDocumentMode(bool set);
extern void _ZN7QTabBar15setDocumentModeEb(void* qthis, bool arg0);
  // proto:  int QTabBar::count();
extern void _ZNK7QTabBar5countEv(void* qthis);
  // proto:  void QTabBar::setChangeCurrentOnDrag(bool change);
extern void _ZN7QTabBar22setChangeCurrentOnDragEb(void* qthis, bool arg0);
  // proto:  QIcon QTabBar::tabIcon(int index);
extern void _ZNK7QTabBar7tabIconEi(void* qthis, int arg0);
  // proto:  QSize QTabBar::minimumSizeHint();
extern void _ZNK7QTabBar15minimumSizeHintEv(void* qthis);
  // proto:  void QTabBar::setTabsClosable(bool closable);
extern void _ZN7QTabBar15setTabsClosableEb(void* qthis, bool arg0);
  // proto:  bool QTabBar::changeCurrentOnDrag();
extern void _ZNK7QTabBar19changeCurrentOnDragEv(void* qthis);
  // proto:  void QTabBar::setTabWhatsThis(int index, const QString & text);
extern void _ZN7QTabBar15setTabWhatsThisEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  const QMetaObject * QTabBar::metaObject();
extern void _ZNK7QTabBar10metaObjectEv(void* qthis);
  // proto:  void QTabBar::QTabBar(const QTabBar & );
extern void* dector_ZN7QTabBarC1ERKS_(void* arg0);
extern void _ZN7QTabBarC1ERKS_(void* qthis, void* arg0);
  // proto:  int QTabBar::insertTab(int index, const QIcon & icon, const QString & text);
extern void _ZN7QTabBar9insertTabEiRK5QIconRK7QString(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  void QTabBar::setTabIcon(int index, const QIcon & icon);
extern void _ZN7QTabBar10setTabIconEiRK5QIcon(void* qthis, int arg0, void* arg1);
  // proto:  bool QTabBar::isMovable();
extern void _ZNK7QTabBar9isMovableEv(void* qthis);
  // proto:  void QTabBar::setExpanding(bool enabled);
extern void _ZN7QTabBar12setExpandingEb(void* qthis, bool arg0);
  // proto:  void QTabBar::removeTab(int index);
extern void _ZN7QTabBar9removeTabEi(void* qthis, int arg0);
  // proto:  void QTabBar::setTabEnabled(int index, bool );
extern void _ZN7QTabBar13setTabEnabledEib(void* qthis, int arg0, bool arg1);
  // proto:  bool QTabBar::isTabEnabled(int index);
extern void _ZNK7QTabBar12isTabEnabledEi(void* qthis, int arg0);
  // proto:  void QTabBar::setCurrentIndex(int index);
extern void _ZN7QTabBar15setCurrentIndexEi(void* qthis, int arg0);
  // proto:  QRect QTabBar::tabRect(int index);
extern void _ZNK7QTabBar7tabRectEi(void* qthis, int arg0);
  // proto:  bool QTabBar::tabsClosable();
extern void _ZNK7QTabBar12tabsClosableEv(void* qthis);
  // proto:  void QTabBar::setMovable(bool movable);
extern void _ZN7QTabBar10setMovableEb(void* qthis, bool arg0);
  // proto:  void QTabBar::setAutoHide(bool hide);
extern void _ZN7QTabBar11setAutoHideEb(void* qthis, bool arg0);
  // proto:  QSize QTabBar::iconSize();
extern void _ZNK7QTabBar8iconSizeEv(void* qthis);
  // proto:  QString QTabBar::tabText(int index);
extern void _ZNK7QTabBar7tabTextEi(void* qthis, int arg0);
  // proto:  QString QTabBar::tabWhatsThis(int index);
extern void _ZNK7QTabBar12tabWhatsThisEi(void* qthis, int arg0);
  // proto:  bool QTabBar::documentMode();
extern void _ZNK7QTabBar12documentModeEv(void* qthis);
  // proto:  int QTabBar::tabAt(const QPoint & pos);
extern void _ZNK7QTabBar5tabAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QTabBar::setTabData(int index, const QVariant & data);
extern void _ZN7QTabBar10setTabDataEiRK8QVariant(void* qthis, int arg0, void* arg1);
  // proto:  void QTabBar::QTabBar(QWidget * parent);
extern void* dector_ZN7QTabBarC1EP7QWidget(void* arg0);
extern void _ZN7QTabBarC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QColor QTabBar::tabTextColor(int index);
extern void _ZNK7QTabBar12tabTextColorEi(void* qthis, int arg0);
  // proto:  void QTabBar::~QTabBar();
extern void _ZN7QTabBarD0Ev(void* qthis);
  // proto:  int QTabBar::insertTab(int index, const QString & text);
extern void _ZN7QTabBar9insertTabEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  int QTabBar::addTab(const QString & text);
extern void _ZN7QTabBar6addTabERK7QString(void* qthis, void* arg0);
  // proto:  int QTabBar::addTab(const QIcon & icon, const QString & text);
extern void _ZN7QTabBar6addTabERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QTabBar::setTabToolTip(int index, const QString & tip);
extern void _ZN7QTabBar13setTabToolTipEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QTabBar::setTabTextColor(int index, const QColor & color);
extern void _ZN7QTabBar15setTabTextColorEiRK6QColor(void* qthis, int arg0, void* arg1);
  // proto:  void QTabBar::moveTab(int from, int to);
extern void _ZN7QTabBar7moveTabEii(void* qthis, int arg0, int arg1);
  // proto:  QVariant QTabBar::tabData(int index);
extern void _ZNK7QTabBar7tabDataEi(void* qthis, int arg0);
  // proto:  bool QTabBar::drawBase();
extern void _ZNK7QTabBar8drawBaseEv(void* qthis);
  // proto:  int QTabBar::currentIndex();
extern void _ZNK7QTabBar12currentIndexEv(void* qthis);
  // proto:  void QTabBar::setDrawBase(bool drawTheBase);
extern void _ZN7QTabBar11setDrawBaseEb(void* qthis, bool arg0);
  // proto:  void QTabBar::setUsesScrollButtons(bool useButtons);
extern void _ZN7QTabBar20setUsesScrollButtonsEb(void* qthis, bool arg0);
  // proto:  QSize QTabBar::sizeHint();
extern void _ZNK7QTabBar8sizeHintEv(void* qthis);
  // proto:  void QTabBar::setIconSize(const QSize & size);
extern void _ZN7QTabBar11setIconSizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QTabBar::setTabText(int index, const QString & text);
extern void _ZN7QTabBar10setTabTextEiRK7QString(void* qthis, int arg0, void* arg1);
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

// class sizeof(QTabBar)=1
type QTabBar struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _tabCloseRequested QTabBar_tabCloseRequested_signal;
//  _tabBarDoubleClicked QTabBar_tabBarDoubleClicked_signal;
//  _tabMoved QTabBar_tabMoved_signal;
//  _tabBarClicked QTabBar_tabBarClicked_signal;
//  _currentChanged QTabBar_currentChanged_signal;
}

  // proto:  bool QTabBar::usesScrollButtons();
func (this *QTabBar) usesScrollButtons(args ...interface{}) () {
  // usesScrollButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar17usesScrollButtonsEv
  default:
    qtrt.ErrorResolve("QTabBar", "usesScrollButtons", args)
  }

}

  // proto:  bool QTabBar::autoHide();
func (this *QTabBar) autoHide(args ...interface{}) () {
  // autoHide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8autoHideEv
  default:
    qtrt.ErrorResolve("QTabBar", "autoHide", args)
  }

}

  // proto:  QString QTabBar::tabToolTip(int index);
func (this *QTabBar) tabToolTip(args ...interface{}) () {
  // tabToolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar10tabToolTipEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabToolTip", args)
  }

}

  // proto:  bool QTabBar::expanding();
func (this *QTabBar) expanding(args ...interface{}) () {
  // expanding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar9expandingEv
  default:
    qtrt.ErrorResolve("QTabBar", "expanding", args)
  }

}

  // proto:  void QTabBar::setDocumentMode(bool set);
func (this *QTabBar) setDocumentMode(args ...interface{}) () {
  // setDocumentMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setDocumentModeEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setDocumentMode", args)
  }

}

  // proto:  int QTabBar::count();
func (this *QTabBar) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar5countEv
  default:
    qtrt.ErrorResolve("QTabBar", "count", args)
  }

}

  // proto:  void QTabBar::setChangeCurrentOnDrag(bool change);
func (this *QTabBar) setChangeCurrentOnDrag(args ...interface{}) () {
  // setChangeCurrentOnDrag(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar22setChangeCurrentOnDragEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setChangeCurrentOnDrag", args)
  }

}

  // proto:  QIcon QTabBar::tabIcon(int index);
func (this *QTabBar) tabIcon(args ...interface{}) () {
  // tabIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabIconEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabIcon", args)
  }

}

  // proto:  QSize QTabBar::minimumSizeHint();
func (this *QTabBar) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QTabBar", "minimumSizeHint", args)
  }

}

  // proto:  void QTabBar::setTabsClosable(bool closable);
func (this *QTabBar) setTabsClosable(args ...interface{}) () {
  // setTabsClosable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setTabsClosableEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabsClosable", args)
  }

}

  // proto:  bool QTabBar::changeCurrentOnDrag();
func (this *QTabBar) changeCurrentOnDrag(args ...interface{}) () {
  // changeCurrentOnDrag()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar19changeCurrentOnDragEv
  default:
    qtrt.ErrorResolve("QTabBar", "changeCurrentOnDrag", args)
  }

}

  // proto:  void QTabBar::setTabWhatsThis(int index, const QString & text);
func (this *QTabBar) setTabWhatsThis(args ...interface{}) () {
  // setTabWhatsThis(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setTabWhatsThisEiRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabWhatsThis", args)
  }

}

  // proto:  const QMetaObject * QTabBar::metaObject();
func (this *QTabBar) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar10metaObjectEv
  default:
    qtrt.ErrorResolve("QTabBar", "metaObject", args)
  }

}

  // proto:  void QTabBar::QTabBar(const QTabBar & );
func NewQTabBar(args ...interface{}) QTabBar {
  return QTabBar{}
}

  // proto:  int QTabBar::insertTab(int index, const QIcon & icon, const QString & text);
func (this *QTabBar) insertTab(args ...interface{}) () {
  // insertTab(int, const class QIcon &, const class QString &)
  // insertTab(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar9insertTabEiRK5QIconRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN7QTabBar9insertTabEiRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "insertTab", args)
  }

}

  // proto:  void QTabBar::setTabIcon(int index, const QIcon & icon);
func (this *QTabBar) setTabIcon(args ...interface{}) () {
  // setTabIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar10setTabIconEiRK5QIcon
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabIcon", args)
  }

}

  // proto:  bool QTabBar::isMovable();
func (this *QTabBar) isMovable(args ...interface{}) () {
  // isMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar9isMovableEv
  default:
    qtrt.ErrorResolve("QTabBar", "isMovable", args)
  }

}

  // proto:  void QTabBar::setExpanding(bool enabled);
func (this *QTabBar) setExpanding(args ...interface{}) () {
  // setExpanding(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar12setExpandingEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setExpanding", args)
  }

}

  // proto:  void QTabBar::removeTab(int index);
func (this *QTabBar) removeTab(args ...interface{}) () {
  // removeTab(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar9removeTabEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "removeTab", args)
  }

}

  // proto:  void QTabBar::setTabEnabled(int index, bool );
func (this *QTabBar) setTabEnabled(args ...interface{}) () {
  // setTabEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar13setTabEnabledEib
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabEnabled", args)
  }

}

  // proto:  bool QTabBar::isTabEnabled(int index);
func (this *QTabBar) isTabEnabled(args ...interface{}) () {
  // isTabEnabled(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12isTabEnabledEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "isTabEnabled", args)
  }

}

  // proto:  void QTabBar::setCurrentIndex(int index);
func (this *QTabBar) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setCurrentIndexEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setCurrentIndex", args)
  }

}

  // proto:  QRect QTabBar::tabRect(int index);
func (this *QTabBar) tabRect(args ...interface{}) () {
  // tabRect(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabRectEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabRect", args)
  }

}

  // proto:  bool QTabBar::tabsClosable();
func (this *QTabBar) tabsClosable(args ...interface{}) () {
  // tabsClosable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12tabsClosableEv
  default:
    qtrt.ErrorResolve("QTabBar", "tabsClosable", args)
  }

}

  // proto:  void QTabBar::setMovable(bool movable);
func (this *QTabBar) setMovable(args ...interface{}) () {
  // setMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar10setMovableEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setMovable", args)
  }

}

  // proto:  void QTabBar::setAutoHide(bool hide);
func (this *QTabBar) setAutoHide(args ...interface{}) () {
  // setAutoHide(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar11setAutoHideEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setAutoHide", args)
  }

}

  // proto:  QSize QTabBar::iconSize();
func (this *QTabBar) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8iconSizeEv
  default:
    qtrt.ErrorResolve("QTabBar", "iconSize", args)
  }

}

  // proto:  QString QTabBar::tabText(int index);
func (this *QTabBar) tabText(args ...interface{}) () {
  // tabText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabTextEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabText", args)
  }

}

  // proto:  QString QTabBar::tabWhatsThis(int index);
func (this *QTabBar) tabWhatsThis(args ...interface{}) () {
  // tabWhatsThis(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12tabWhatsThisEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabWhatsThis", args)
  }

}

  // proto:  bool QTabBar::documentMode();
func (this *QTabBar) documentMode(args ...interface{}) () {
  // documentMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12documentModeEv
  default:
    qtrt.ErrorResolve("QTabBar", "documentMode", args)
  }

}

  // proto:  int QTabBar::tabAt(const QPoint & pos);
func (this *QTabBar) tabAt(args ...interface{}) () {
  // tabAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar5tabAtERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabAt", args)
  }

}

  // proto:  void QTabBar::setTabData(int index, const QVariant & data);
func (this *QTabBar) setTabData(args ...interface{}) () {
  // setTabData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar10setTabDataEiRK8QVariant
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabData", args)
  }

}

  // proto:  QColor QTabBar::tabTextColor(int index);
func (this *QTabBar) tabTextColor(args ...interface{}) () {
  // tabTextColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12tabTextColorEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabTextColor", args)
  }

}

  // proto:  void QTabBar::~QTabBar();
func (this *QTabBar) FreeQTabBar(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTabBar", "~QTabBar", args)
  }

}

  // proto:  int QTabBar::addTab(const QString & text);
func (this *QTabBar) addTab(args ...interface{}) () {
  // addTab(const class QString &)
  // addTab(const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar6addTabERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN7QTabBar6addTabERK5QIconRK7QString
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "addTab", args)
  }

}

  // proto:  void QTabBar::setTabToolTip(int index, const QString & tip);
func (this *QTabBar) setTabToolTip(args ...interface{}) () {
  // setTabToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar13setTabToolTipEiRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabToolTip", args)
  }

}

  // proto:  void QTabBar::setTabTextColor(int index, const QColor & color);
func (this *QTabBar) setTabTextColor(args ...interface{}) () {
  // setTabTextColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setTabTextColorEiRK6QColor
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabTextColor", args)
  }

}

  // proto:  void QTabBar::moveTab(int from, int to);
func (this *QTabBar) moveTab(args ...interface{}) () {
  // moveTab(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar7moveTabEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "moveTab", args)
  }

}

  // proto:  QVariant QTabBar::tabData(int index);
func (this *QTabBar) tabData(args ...interface{}) () {
  // tabData(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabDataEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "tabData", args)
  }

}

  // proto:  bool QTabBar::drawBase();
func (this *QTabBar) drawBase(args ...interface{}) () {
  // drawBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8drawBaseEv
  default:
    qtrt.ErrorResolve("QTabBar", "drawBase", args)
  }

}

  // proto:  int QTabBar::currentIndex();
func (this *QTabBar) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12currentIndexEv
  default:
    qtrt.ErrorResolve("QTabBar", "currentIndex", args)
  }

}

  // proto:  void QTabBar::setDrawBase(bool drawTheBase);
func (this *QTabBar) setDrawBase(args ...interface{}) () {
  // setDrawBase(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar11setDrawBaseEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setDrawBase", args)
  }

}

  // proto:  void QTabBar::setUsesScrollButtons(bool useButtons);
func (this *QTabBar) setUsesScrollButtons(args ...interface{}) () {
  // setUsesScrollButtons(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar20setUsesScrollButtonsEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setUsesScrollButtons", args)
  }

}

  // proto:  QSize QTabBar::sizeHint();
func (this *QTabBar) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8sizeHintEv
  default:
    qtrt.ErrorResolve("QTabBar", "sizeHint", args)
  }

}

  // proto:  void QTabBar::setIconSize(const QSize & size);
func (this *QTabBar) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar11setIconSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTabBar", "setIconSize", args)
  }

}

  // proto:  void QTabBar::setTabText(int index, const QString & text);
func (this *QTabBar) setTabText(args ...interface{}) () {
  // setTabText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar10setTabTextEiRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTabBar", "setTabText", args)
  }

}

// <= body block end

