package vanish

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertFileExists(t *testing.T, path string) {
	_, err := os.Stat(path)
	assert.Nilf(t, err, "error calling 'stat' on path %s: %s", path, err)
}

func assertFileDoesNotExists(t *testing.T, path string) {
	_, err := os.Stat(path)
	assert.NotNilf(t, err, "Path %s should not exist", path)
}

func TestFile(t *testing.T) {
	var filename string

	err := File(func(name string) {
		fi, err := os.Stat(name)
		assert.Nil(t, err)
		assert.True(t, fi.Mode().IsRegular())

		filename = name
	})
	assert.Nil(t, err)

	assertFileDoesNotExists(t, filename)
}

func TestFileIn(t *testing.T) {
	err := Dir(func(dir string) {
		err := FileIn(dir, func(filename string) {
			assert.True(t, strings.HasPrefix(filename, dir))
			assertFileExists(t, filename)
		})
		assert.Nil(t, err)
	})
	assert.Nil(t, err)
}

func TestEmptyDir(t *testing.T) {
	var dirname string

	err := Dir(func(name string) {
		fi, err := os.Stat(name)
		assert.Nil(t, err)
		assert.True(t, fi.Mode().IsDir())

		dirname = name
	})
	assert.Nil(t, err)

	assertFileDoesNotExists(t, dirname)
}

func TestDirNotEmpty(t *testing.T) {
	var dirname string

	assert.Nil(t, Dir(func(name string) {
		_, err := os.CreateTemp(name, "")
		assert.Nil(t, err)

		dirname = name
	}))

	assertFileDoesNotExists(t, dirname)
}

func TestDirIn(t *testing.T) {
	err := Dir(func(parent string) {
		err := DirIn(parent, func(dir string) {
			assertFileExists(t, dir)
			assert.True(t, strings.HasPrefix(dir, parent))
		})
		assert.Nil(t, err)
	})
	assert.Nil(t, err)
}

func TestEnv(t *testing.T) {
	key := "VANISH_TEST_ENV_42XYZ"
	val := "foob&ar+/xqye$@"
	assert.Nil(t, os.Setenv(key, val))

	err := Env(func() {
		assert.Equal(t, val, os.Getenv(key))
		assert.Nil(t, os.Setenv(key, "yolofoo"))
	})
	assert.Nil(t, err)

	assert.Equal(t, val, os.Getenv(key))
}

func ExampleFile() {
	File(func(name string) {
		// 'name' is the name of a temporary file that’ll be deleted at the end
		// of the function. You can thus use it for your tests without worrying
		// about it
	})
}

func ExampleDir() {
	Dir(func(name string) {
		// Dir works like File, but it creates a directory instead. Use it as
		// you wish, it’ll be deleted at the end of this function.
	})
}

func ExampleEnv() {
	os.Setenv("VANISH", "xyz")

	Env(func() {
		// Env saves the environment before calling this function, ensuring
		// it’s restored as it was at its end.

		os.Setenv("VANISH", "foo")

		fmt.Printf("VANISH (1) = %s\n", os.Getenv("VANISH"))
	})

	fmt.Printf("VANISH (2) = %s\n", os.Getenv("VANISH"))

	// Output:
	// VANISH (1) = foo
	// VANISH (2) = xyz
}
