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

Future work
--------------

These are the features I would like to implement someday:

1. To compile C++ files.
1. -D flag.
1. It should be concurrent, run several commands at the same time.
1. To support memcached.
1. To cache memcached hits in the disk cache.
1. To support distributed compilation.
1. To improve disk cache: set max size, set ttl.
1. To improve performance of disk cache.
1. More options to the preprocessor, compiler, ar and linker.

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
