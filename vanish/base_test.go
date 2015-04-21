package vanish

import (
	"os"
	"testing"

	"github.com/bfontaine/vanish/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func fileExists(path string) bool {
	// we don’t really check for errors here
	_, err := os.Stat(path)
	return err == nil
}

func TestFile(t *testing.T) {
	var filename string

	File(func(name string) {
		filename = name
		assert.True(t, fileExists(filename))
	})

	assert.False(t, fileExists(filename))
}

func TestDir(t *testing.T) {
	var dirname string

	Dir(func(name string) {
		dirname = name
		assert.True(t, fileExists(dirname))
	})

	assert.False(t, fileExists(dirname))
}

func TestEnv(t *testing.T) {
	key := "VANISH_TEST_ENV_42XYZ"
	val := "foob&ar+/xqye$@"
	os.Setenv(key, val)

	Env(func() {
		assert.Equal(t, val, os.Getenv(key))
		os.Setenv(key, "yolofoo")
	})

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
	Env(func() {
		// Env saves the environment before calling this function, ensuring
		// it’s restored as it was at its end.
	})
}
