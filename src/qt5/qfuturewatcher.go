package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qfuturewatcher.h
// dst-file: /src/core/qfuturewatcher.go
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
  // proto:  bool QFutureWatcherBase::isRunning();
extern void _ZNK18QFutureWatcherBase9isRunningEv(void* qthis);
  // proto:  void QFutureWatcherBase::setPaused(bool paused);
extern void _ZN18QFutureWatcherBase9setPausedEb(void* qthis, bool arg0);
  // proto:  int QFutureWatcherBase::progressMinimum();
extern void _ZNK18QFutureWatcherBase15progressMinimumEv(void* qthis);
  // proto:  void QFutureWatcherBase::resume();
extern void _ZN18QFutureWatcherBase6resumeEv(void* qthis);
  // proto:  const QMetaObject * QFutureWatcherBase::metaObject();
extern void _ZNK18QFutureWatcherBase10metaObjectEv(void* qthis);
  // proto:  bool QFutureWatcherBase::isFinished();
extern void _ZNK18QFutureWatcherBase10isFinishedEv(void* qthis);
  // proto:  int QFutureWatcherBase::progressMaximum();
extern void _ZNK18QFutureWatcherBase15progressMaximumEv(void* qthis);
  // proto:  bool QFutureWatcherBase::event(QEvent * event);
extern void _ZN18QFutureWatcherBase5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  bool QFutureWatcherBase::isCanceled();
extern void _ZNK18QFutureWatcherBase10isCanceledEv(void* qthis);
  // proto:  int QFutureWatcherBase::progressValue();
extern void _ZNK18QFutureWatcherBase13progressValueEv(void* qthis);
  // proto:  bool QFutureWatcherBase::isStarted();
extern void _ZNK18QFutureWatcherBase9isStartedEv(void* qthis);
  // proto:  void QFutureWatcherBase::setPendingResultsLimit(int limit);
extern void _ZN18QFutureWatcherBase22setPendingResultsLimitEi(void* qthis, int arg0);
  // proto:  void QFutureWatcherBase::cancel();
extern void _ZN18QFutureWatcherBase6cancelEv(void* qthis);
  // proto:  bool QFutureWatcherBase::isPaused();
extern void _ZNK18QFutureWatcherBase8isPausedEv(void* qthis);
  // proto:  void QFutureWatcherBase::pause();
extern void _ZN18QFutureWatcherBase5pauseEv(void* qthis);
  // proto:  QString QFutureWatcherBase::progressText();
extern void _ZNK18QFutureWatcherBase12progressTextEv(void* qthis);
  // proto:  void QFutureWatcherBase::QFutureWatcherBase(QObject * parent);
extern void* dector_ZN18QFutureWatcherBaseC1EP7QObject(void* arg0);
extern void _ZN18QFutureWatcherBaseC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QFutureWatcherBase::togglePaused();
extern void _ZN18QFutureWatcherBase12togglePausedEv(void* qthis);
  // proto:  void QFutureWatcherBase::waitForFinished();
extern void _ZN18QFutureWatcherBase15waitForFinishedEv(void* qthis);
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

// class sizeof(QFutureWatcherBase)=1
type QFutureWatcherBase struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _progressRangeChanged QFutureWatcherBase_progressRangeChanged_signal;
//  _resumed QFutureWatcherBase_resumed_signal;
//  _progressValueChanged QFutureWatcherBase_progressValueChanged_signal;
//  _started QFutureWatcherBase_started_signal;
//  _resultReadyAt QFutureWatcherBase_resultReadyAt_signal;
//  _resultsReadyAt QFutureWatcherBase_resultsReadyAt_signal;
//  _paused QFutureWatcherBase_paused_signal;
//  _canceled QFutureWatcherBase_canceled_signal;
//  _finished QFutureWatcherBase_finished_signal;
//  _progressTextChanged QFutureWatcherBase_progressTextChanged_signal;
}

  // proto:  bool QFutureWatcherBase::isRunning();
func (this *QFutureWatcherBase) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isRunningEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isRunning", args)
  }

}

  // proto:  void QFutureWatcherBase::setPaused(bool paused);
func (this *QFutureWatcherBase) setPaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase9setPausedEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPaused", args)
  }

}

  // proto:  int QFutureWatcherBase::progressMinimum();
func (this *QFutureWatcherBase) progressMinimum(args ...interface{}) () {
  // progressMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMinimumEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMinimum", args)
  }

}

  // proto:  void QFutureWatcherBase::resume();
func (this *QFutureWatcherBase) resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6resumeEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "resume", args)
  }

}

  // proto:  const QMetaObject * QFutureWatcherBase::metaObject();
func (this *QFutureWatcherBase) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10metaObjectEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "metaObject", args)
  }

}

  // proto:  bool QFutureWatcherBase::isFinished();
func (this *QFutureWatcherBase) isFinished(args ...interface{}) () {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isFinishedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isFinished", args)
  }

}

  // proto:  int QFutureWatcherBase::progressMaximum();
func (this *QFutureWatcherBase) progressMaximum(args ...interface{}) () {
  // progressMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMaximumEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMaximum", args)
  }

}

  // proto:  bool QFutureWatcherBase::event(QEvent * event);
func (this *QFutureWatcherBase) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase5eventEP6QEvent
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "event", args)
  }

}

  // proto:  bool QFutureWatcherBase::isCanceled();
func (this *QFutureWatcherBase) isCanceled(args ...interface{}) () {
  // isCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isCanceledEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isCanceled", args)
  }

}

  // proto:  int QFutureWatcherBase::progressValue();
func (this *QFutureWatcherBase) progressValue(args ...interface{}) () {
  // progressValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase13progressValueEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressValue", args)
  }

}

  // proto:  bool QFutureWatcherBase::isStarted();
func (this *QFutureWatcherBase) isStarted(args ...interface{}) () {
  // isStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isStartedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isStarted", args)
  }

}

  // proto:  void QFutureWatcherBase::setPendingResultsLimit(int limit);
func (this *QFutureWatcherBase) setPendingResultsLimit(args ...interface{}) () {
  // setPendingResultsLimit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase22setPendingResultsLimitEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPendingResultsLimit", args)
  }

}

  // proto:  void QFutureWatcherBase::cancel();
func (this *QFutureWatcherBase) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6cancelEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "cancel", args)
  }

}

  // proto:  bool QFutureWatcherBase::isPaused();
func (this *QFutureWatcherBase) isPaused(args ...interface{}) () {
  // isPaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase8isPausedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isPaused", args)
  }

}

  // proto:  void QFutureWatcherBase::pause();
func (this *QFutureWatcherBase) pause(args ...interface{}) () {
  // pause()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase5pauseEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "pause", args)
  }

}

  // proto:  QString QFutureWatcherBase::progressText();
func (this *QFutureWatcherBase) progressText(args ...interface{}) () {
  // progressText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase12progressTextEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressText", args)
  }

}

  // proto:  void QFutureWatcherBase::QFutureWatcherBase(QObject * parent);
func NewQFutureWatcherBase(args ...interface{}) QFutureWatcherBase {
  return QFutureWatcherBase{}
}

  // proto:  void QFutureWatcherBase::togglePaused();
func (this *QFutureWatcherBase) togglePaused(args ...interface{}) () {
  // togglePaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase12togglePausedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "togglePaused", args)
  }

}

  // proto:  void QFutureWatcherBase::waitForFinished();
func (this *QFutureWatcherBase) waitForFinished(args ...interface{}) () {
  // waitForFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase15waitForFinishedEv
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "waitForFinished", args)
  }

}

// <= body block end

