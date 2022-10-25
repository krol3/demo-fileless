package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
)

// the constant values below are valid for x86_64
const (
	mfdCloexec                   uintptr = 1 // 0x0001 - 1 in base 10
	SYS_MEMFD_CREATE_LINUX_AMD64 uintptr = 319
)

func runFromMemory(displayName string, filePathArgs []string) error {
	fdName := "" //  unsafe.Pointer for unsafe memory access
	fd, _, e1 := syscall.Syscall(SYS_MEMFD_CREATE_LINUX_AMD64, uintptr(unsafe.Pointer(&fdName)), mfdCloexec, 0)
	fmt.Println("---> Create the memory file descriptor: ", fd)

	if e1 != 0 {
		e := os.NewSyscallError("memfd_create", e1)
		return e
	}

	filePath := filePathArgs[0]
	fmt.Println("---> Reading ELF: ", filePath)
	buffer, _ := ioutil.ReadFile(filePath)

	fmt.Printf("---> Writing ELF file from [%s] in the memory file descriptor [%d] \n", filePath, fd)
	_, _ = syscall.Write(int(fd), buffer)

	fdPath := fmt.Sprintf("/proc/self/fd/%d", fd)
	fmt.Printf("---> Executes execve( pathname=%s, argv[]={%s}, nil)\n", fdPath, displayName)
	// # int execve(const char *pathname, char *const argv[],char *const envp[]);
	execErr := syscall.Exec(fdPath, []string{displayName}, nil)
	if execErr != nil {
		panic(execErr)
	}

	return nil
}

func main() {
	lenArgs := len(os.Args)
	if lenArgs < 3 {
		fmt.Println("Usage: process_name elf_binary")
		os.Exit(1)
	}

	fmt.Println("Usage: process_name elf_binary")
	runFromMemory(os.Args[1], os.Args[2:])
}
