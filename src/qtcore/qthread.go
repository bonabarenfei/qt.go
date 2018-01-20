//  header block begin
// /usr/include/qt/QtCore/qthread.h
// #include <qthread.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QThread struct {
	*QObject
}

func (this *QThread) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQThreadFromPointer(cthis unsafe.Pointer) *QThread {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QThread{bcthis0}
}

// /usr/include/qt/QtCore/qthread.h:72
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QThread) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:74
// index:0
// Public static
// Qt::HANDLE currentThreadId()
func (this *QThread) CurrentThreadId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread15currentThreadIdEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QThread_CurrentThreadId() {
	var nilthis *QThread
	nilthis.CurrentThreadId()
}

// /usr/include/qt/QtCore/qthread.h:75
// index:0
// Public static
// QThread * currentThread()
func (this *QThread) CurrentThread() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread13currentThreadEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QThread_CurrentThread() {
	var nilthis *QThread
	nilthis.CurrentThread()
}

// /usr/include/qt/QtCore/qthread.h:76
// index:0
// Public static
// int idealThreadCount()
func (this *QThread) IdealThreadCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread16idealThreadCountEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QThread_IdealThreadCount() {
	var nilthis *QThread
	nilthis.IdealThreadCount()
}

// /usr/include/qt/QtCore/qthread.h:77
// index:0
// Public static
// void yieldCurrentThread()
func (this *QThread) YieldCurrentThread() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread18yieldCurrentThreadEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QThread_YieldCurrentThread() {
	var nilthis *QThread
	nilthis.YieldCurrentThread()
}

// /usr/include/qt/QtCore/qthread.h:79
// index:0
// Public
// void QThread(class QObject *)
func NewQThread(parent unsafe.Pointer) *QThread {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThreadC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQThreadFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qthread.h:80
// index:0
// Public virtual
// void ~QThread()
func DeleteQThread(*QThread) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThreadD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:96
// index:0
// Public
// void setPriority(enum QThread::Priority)
func (this *QThread) SetPriority(priority int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread11setPriorityENS_8PriorityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:97
// index:0
// Public
// QThread::Priority priority()
func (this *QThread) Priority() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread8priorityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:99
// index:0
// Public
// bool isFinished()
func (this *QThread) IsFinished() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread10isFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:100
// index:0
// Public
// bool isRunning()
func (this *QThread) IsRunning() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:102
// index:0
// Public
// void requestInterruption()
func (this *QThread) RequestInterruption() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread19requestInterruptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:103
// index:0
// Public
// bool isInterruptionRequested()
func (this *QThread) IsInterruptionRequested() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread23isInterruptionRequestedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:105
// index:0
// Public
// void setStackSize(uint)
func (this *QThread) SetStackSize(stackSize uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread12setStackSizeEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stackSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:106
// index:0
// Public
// uint stackSize()
func (this *QThread) StackSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9stackSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:108
// index:0
// Public
// void exit(int)
func (this *QThread) Exit(retcode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4exitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &retcode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:110
// index:0
// Public
// QAbstractEventDispatcher * eventDispatcher()
func (this *QThread) EventDispatcher() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread15eventDispatcherEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:111
// index:0
// Public
// void setEventDispatcher(class QAbstractEventDispatcher *)
func (this *QThread) SetEventDispatcher(eventDispatcher unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), eventDispatcher)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:113
// index:0
// Public virtual
// bool event(class QEvent *)
func (this *QThread) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:114
// index:0
// Public
// int loopLevel()
func (this *QThread) LoopLevel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9loopLevelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:134
// index:0
// Public
// void start(enum QThread::Priority)
func (this *QThread) Start(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5startENS_8PriorityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:135
// index:0
// Public
// void terminate()
func (this *QThread) Terminate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread9terminateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:136
// index:0
// Public
// void quit()
func (this *QThread) Quit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4quitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:140
// index:0
// Public
// bool wait(unsigned long)
func (this *QThread) Wait(time uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4waitEm", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &time)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:142
// index:0
// Public static
// void sleep(unsigned long)
func (this *QThread) Sleep(arg0 uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5sleepEm", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThread_Sleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Sleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:143
// index:0
// Public static
// void msleep(unsigned long)
func (this *QThread) Msleep(arg0 uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread6msleepEm", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThread_Msleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Msleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:144
// index:0
// Public static
// void usleep(unsigned long)
func (this *QThread) Usleep(arg0 uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread6usleepEm", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThread_Usleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Usleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:151
// index:0
// Protected virtual
// void run()
func (this *QThread) Run() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread3runEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:152
// index:0
// Protected
// int exec()
func (this *QThread) Exec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4execEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthread.h:154
// index:0
// Protected static
// void setTerminationEnabled(_Bool)
func (this *QThread) SetTerminationEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread21setTerminationEnabledEb", ffiqt.FFI_TYPE_POINTER, enabled)
	gopp.ErrPrint(err, rv)
}
func QThread_SetTerminationEnabled(enabled bool) {
	var nilthis *QThread
	nilthis.SetTerminationEnabled(enabled)
}

//  body block end
