// Thanks to https://stackoverflow.com/questions/33516053/windows-encrypted-rdp-passwords-in-golang
package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func bytesToBlob(bytes []byte) *windows.DataBlob {
	blob := &windows.DataBlob{Size: uint32(len(bytes))}
	if len(bytes) > 0 {
		blob.Data = &bytes[0]
	}
	return blob
}

func Decrypt(data []byte) ([]byte, error) {

	out := windows.DataBlob{}
	var outName *uint16

	err := windows.CryptUnprotectData(bytesToBlob(data), &outName, nil, 0, nil, 0, &out)
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt DPAPI protected data: %w", err)
	}
	ret := make([]byte, out.Size)
	copy(ret, unsafe.Slice(out.Data, out.Size))

	windows.LocalFree(windows.Handle(unsafe.Pointer(out.Data)))
	windows.LocalFree(windows.Handle(unsafe.Pointer(outName)))

	return ret, nil
}
