package main

import (
	"github.com/g7r/wtrdirtycache/pkg1"
	"github.com/g7r/wtrdirtycache/pkg2"
)

func main() {
	pkg1.Do(&pkg2.Foo{})
}
