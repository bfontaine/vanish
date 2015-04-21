package vanish

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
