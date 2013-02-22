mjolnir
=======

_[mjolnir](https://github.com/daniel-fanjul-alcuten/mjolnir)_ is a launcher to simplify the test, compilation and execution of programs using libmjolnir.

libmjolnir
==========

_[libmjolnir](https://github.com/daniel-fanjul-alcuten/libmjolnir)_ is a go library to write build systems.

libmjolnir-demo
===============

_[libmjolnir-demo](https://github.com/daniel-fanjul-alcuten/libmjolnir-demo)_ is an example project using libmjolnir.

Configuration
-------------

Many build systems use plain files that are hard to refactor when the source files are refactored. _libmjolnir_ is a go library that allows to write a program that builds the source files, so the build system in go must be compiled before the source files.

Current status
--------------

These are the features currently implemented:

1. It compiles C files.
1. -I flag is passed to the preprocessor.
1. -std flag is passed to the compiler.

Example
-------

<pre>
package foo

import (
  ...
)

type Foo struct {
  Files []*CFile
}

func NewFoo(mjölnir *Mjölnir, basepath string) *Foo {
  path := filepath.Join(basepath, "foo")
  f1 := mjölnir.CFile(filepath.Join(path, "foo1.c")).SetStd("c99").Includes(basepath))
  f2 := mjölnir.CFile(filepath.Join(path, "foo2.c")).SetStd("c99").DependsOn(f1)
  return &Foo{[]*CFile{f1, f2}}
}
</pre>

<pre>
package main

import (
  ...
)

func main() {

  mjölnir := NewMjölnir()
  mjölnir.Verbose = 1
  u, err := user.Current()
  if err != nil {
    log.Fatalln("error:", err)
  }
  path := filepath.Join(os.TempDir(), "disk-cache-"+u.Uid+".mjolnir")
  mjölnir.DataCache = NewDiskDataCache(path)

  basepath, err := os.Getwd()
  if err != nil {
    log.Fatalln("error:", err)
  }

  foo := NewFoo(mjölnir, basepath)
  main := mjölnir.CFile(filepath.Join(basepath, "main.c")).DependsOn(foo.Files...)

  mjölnir.CExecutable("mmain").Link(main)
  if err := mjölnir.Build(); err != nil {
    log.Fatalln("error:", err)
  }
}
</pre>
