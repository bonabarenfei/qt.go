package qtqml

// /usr/include/qt/QtQml/qqmldebug.h
// #include <qqmldebug.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlDebuggingEnabler struct {
	*qtrt.CObject
}

func (this *QQmlDebuggingEnabler) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlDebuggingEnabler) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlDebuggingEnablerFromPointer(cthis unsafe.Pointer) *QQmlDebuggingEnabler {
	return &QQmlDebuggingEnabler{&qtrt.CObject{cthis}}
}
func (*QQmlDebuggingEnabler) NewFromPointer(cthis unsafe.Pointer) *QQmlDebuggingEnabler {
	return NewQQmlDebuggingEnablerFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmldebug.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlDebuggingEnabler(_Bool)
func NewQQmlDebuggingEnabler(printWarning bool) *QQmlDebuggingEnabler {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQmlDebuggingEnablerC2Eb", ffiqt.FFI_TYPE_POINTER, printWarning)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlDebuggingEnablerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmldebug.h:65
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setServices(const QStringList &)
func (this *QQmlDebuggingEnabler) SetServices(services *qtcore.QStringList) {
	var convArg0 = services.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler11setServicesERK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QQmlDebuggingEnabler_SetServices(services *qtcore.QStringList) {
	var nilthis *QQmlDebuggingEnabler
	nilthis.SetServices(services)
}

// /usr/include/qt/QtQml/qqmldebug.h:67
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool startTcpDebugServer(int, enum QQmlDebuggingEnabler::StartMode, const QString &)
func (this *QQmlDebuggingEnabler) StartTcpDebugServer(port int, mode int, hostName *qtcore.QString) bool {
	var convArg2 = hostName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler19startTcpDebugServerEiNS_9StartModeERK7QString", ffiqt.FFI_TYPE_POINTER, port, mode, convArg2)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QQmlDebuggingEnabler_StartTcpDebugServer(port int, mode int, hostName *qtcore.QString) bool {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.StartTcpDebugServer(port, mode, hostName)
	return rv
}

// /usr/include/qt/QtQml/qqmldebug.h:69
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool connectToLocalDebugger(const QString &, enum QQmlDebuggingEnabler::StartMode)
func (this *QQmlDebuggingEnabler) ConnectToLocalDebugger(socketFileName *qtcore.QString, mode int) bool {
	var convArg0 = socketFileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQmlDebuggingEnabler22connectToLocalDebuggerERK7QStringNS_9StartModeE", ffiqt.FFI_TYPE_POINTER, convArg0, mode)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QQmlDebuggingEnabler_ConnectToLocalDebugger(socketFileName *qtcore.QString, mode int) bool {
	var nilthis *QQmlDebuggingEnabler
	rv := nilthis.ConnectToLocalDebugger(socketFileName, mode)
	return rv
}

type QQmlDebuggingEnabler__StartMode = int

const QQmlDebuggingEnabler__DoNotWaitForClient QQmlDebuggingEnabler__StartMode = 0
const QQmlDebuggingEnabler__WaitForClient QQmlDebuggingEnabler__StartMode = 1

//  body block end