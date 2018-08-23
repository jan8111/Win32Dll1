package main
import (
	"fmt"
	"syscall"
	"unsafe"
)


func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func Lib_add(a, b int) {
	lib := syscall.NewLazyDLL("Win32Dll1.dll")
	add := lib.NewProc("sj")

	ret, _, err := add.Call()
	if err != nil {
		fmt.Println("lib.dll运算结果为:", ret)
	}

}

func DllTestDef_add(a, b int) {
	DllTestDef, _ := syscall.LoadLibrary("Win32Dll1.dll")
	fmt.Println("+++++++syscall.LoadLibrary:", DllTestDef, "+++++++")
	defer syscall.FreeLibrary(DllTestDef)
	add, err := syscall.GetProcAddress(DllTestDef, "add")
	fmt.Println("GetProcAddress", add)

	ret, _, err := syscall.Syscall(add,
		2,
		IntPtr(a),
		IntPtr(b),
		0)
	if err != nil {
		fmt.Println("DllTestDef.dll运算结果为:", ret)
	}

}

func DllTestDef_add2(a, b int) {
	DllTestDef := syscall.MustLoadDLL("Win32Dll1.dll")
	add := DllTestDef.MustFindProc("add")

	fmt.Println("+++++++MustFindProc：", add, "+++++++")
	ret, _, err := add.Call(IntPtr(a), IntPtr(b))
	if err != nil {
		fmt.Println("DllTestDef的运算结果为:", ret)
	}
}

func main() {
	Lib_add(4, 5)
	//DllTestDef_add(4, 5)
	//DllTestDef_add2(4, 5)
}