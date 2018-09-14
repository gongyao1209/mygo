package test

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func FileRead()  {
	name := "mytestfile.txt"
	name = "./tempdir/myfile.txt"

	Ioutil(name)
	OsIoutil(name)
	fileRead(name)
	//BufioRead(name)
}

//文件读取操作
//使用 ioutil
func Ioutil(name string)  {

	if contents, err := ioutil.ReadFile(name); err == nil {
		result := strings.Replace(string(contents), "\n", "", 100)
		fmt.Println(result)
	} else {
		panic(err)
	}
}

//使用 os 读取
func OsIoutil(name string)  {
	if obj, err := os.Open(name); err == nil {
		defer obj.Close()

		if contents, err := ioutil.ReadAll(obj); err == nil {
			result := strings.Replace(string(contents), "\n", "", 100)
			fmt.Println(result)
		}
	}
}

//使用 *File 文件对象 读取文件内容
func fileRead(name string)  {
	if obj, err := os.Open(name); err == nil {
		defer obj.Close()

		buf := make([]byte, 1024)

		if n, err := obj.Read(buf); err == nil {
			fmt.Println(strconv.Itoa(n), "Buf length:", strconv.Itoa(len(buf)))
			result := strings.Replace(string(buf), "\n", "", 100)
			fmt.Println(result)
		}
	} else {
		panic(err)
	}
}

func BufioRead(name string)  {
	if obj, err := os.Open(name); err == nil {
		defer obj.Close()

		reader := bufio.NewReader(obj)
		if res, err := reader.ReadString(byte('@')); err == nil {
			fmt.Println("使用ReadSlince相关方法读取内容：", res)
		} else {
			panic(err)
		}

		buf := make([]byte, 1042)
		if n, err := reader.Read(buf); err == nil {
			fmt.Println(strconv.Itoa(n), "Buf length:", strconv.Itoa(len(buf)))
			result := strings.Replace(string(buf), "\n", "", 100)
			fmt.Println(result)
		} else {
			panic(err)
		}
	}
}

//================================================================================================
func FileWrite()  {
	name := "./tempdir/testwritefile.txt"
	content := "报君黄金台上意, 提携玉龙为君死\n"
	write1Ioutil(name, content)

	content = "hello, world\n"
	write2FileWrite(name, content)

	content = "hello, gongyao\n"
	write3WithIo(name, content)

	write4WithBufio(name, content)
}

//使用ioutil包进行文件写入
func write1Ioutil(name, content string)  {
	data := []byte(content)

	if err := ioutil.WriteFile(name, data, 0644); err == nil {
		fmt.Println("写入文件成功，", content)
	}
}

//使用 os.OpenFile()相关函数打开文件对象
func write2FileWrite(name, content string)  {
	/*
	const (
		O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
		O_WRONLY int = syscall.O_WRONLY // 只写打开文件
		O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
		O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
		O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
		O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
		O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
	)
	 */
	obj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		os.Exit(2)
	}
	defer obj.Close()

	if _, err := obj.WriteString(content); err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.",content)
	}

	contents := []byte(content)
	if _, err := obj.Write(contents); err == nil {
		fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.",content)
	}
}

//使用io包中的相关函数写入文件
func write3WithIo(name, content string)  {
	obj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		os.Exit(2)
	}
	defer obj.Close()

	if _, err := io.WriteString(obj, content); err == nil {
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.",content)
	}
}

//使用bufio 包中的相关函数写入文件
func write4WithBufio(name, content string)  {
	obj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		os.Exit(2)
	}
	defer obj.Close()

	write_obj := bufio.NewWriterSize(obj, 4096)

	if _, err := write_obj.WriteString(content); err == nil {
		fmt.Println("Successful appending buffer and flush to file with bufio's Writer obj WriteString method ", content)
	}

	contents := []byte(content)
	if _, err := write_obj.Write(contents); err == nil {
		fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Write method.", content)
		if err := write_obj.Flush(); err != nil {
			panic(err)
		}
		fmt.Println("Successful flush the buffer data to file ", content)
	}
}