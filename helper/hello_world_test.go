/**
nama file utk unittest harus diakhiri _test
nama func utk unit test hrs diawali Test
func hrs ada parameter (t *testing.T) dan engga ada return value nya

di terminal buat jalanin pake 
go test => jalanin smua file test dan func
go test -v => jalanin smua file test dan func sambil ditunjukin func yg dijalanin apa aj
go test -v -run nama_func => buat jalanin 1 func aj (yg dirun sesuai dg awalan yg dikasi / prefix)
go test ./... => run testing di top level dir (smua package dirun testing nya)

t.Fail = masih menjalankan code di bawahnya (dlm 1 func testing)
t.FailNow = langsung brenti di FailNow, tp msh lanjut ke func testing slanjutnya
recommend:
t.Error(args) = sama kek Fail tp dikasi log error nya (args)
t.Fatal(args) = sama kek FailNow tp ada log error (args)

assert menggunakan Fail
require menggunakan FailNow

untuk mengskip unit test bs pake skip test (biasanya ketika di os tertentu)
*/

package helper

import (
	"testing"
	"fmt"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "runtime"
)

func TestSkip(t *testing.T) {
  if runtime.GOOS == "linux" {
    //fmt.Println("TestSkip done")  --> just checking
    t.Skip("can't run on linux")
  }

  result := HelloWorld("Eko")
  require.Equal(t, "Hello eko", result, "Result must be 'Hello eko'")
}

func TestHelloWorldRequire(t *testing.T) {
  result := HelloWorld("Eko")
  require.Equal(t, "Hello eko", result, "Result must be 'Hello eko'")
  fmt.Println("TestHelloWorldAssert w/ require done")
}

func TestHelloWorldAssert(t *testing.T) {
  result := HelloWorld("Eko")
  assert.Equal(t, "Hello eko", result, "Result must be 'Hello eko'")
  fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("eko") // kalo var hanya dipakai di file ini aj pake := blh

	if result != "Hello eko" {
		// error
		//t.Fail()
		t.Error("result must be 'Hello eko'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldKk(t *testing.T) {
	result := HelloWorld("Kk") // kalo var hanya dipakai di file ini aj pake := blh

	if result != "Hello Kk" {
		// error
		//t.FailNow()
		t.Fatal("result must be 'Hello Kk'")
	}

	fmt.Println("TestHelloWorldKk Done")
}