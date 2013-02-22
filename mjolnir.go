package mjölnir

import (
	. "github.com/daniel-fanjul-alcuten/libmjolnir/utils"
)

type Mjölnir struct {
	Verbose      int
	Unclean      bool
	Backend      string
	LocalCache   *localCache
	DataCache    dataCache
	cFiles       map[string]*CFile
	cLibraries   map[string]*CLibrary
	cExecutables map[string]*CExecutable
}

func NewMjölnir() *Mjölnir {
	return &Mjölnir{
		cFiles:       make(map[string]*CFile),
		cLibraries:   make(map[string]*CLibrary),
		cExecutables: make(map[string]*CExecutable)}
}

func (mjölnir *Mjölnir) preprocessor() string {
	if mjölnir.Backend == "" {
		return "cc"
	}
	return mjölnir.Backend
}

func (mjölnir *Mjölnir) compiler() string {
	return mjölnir.preprocessor()
}

func (mjölnir *Mjölnir) libraryLinker() string {
	return "ar"
}

func (mjölnir *Mjölnir) executableLinker() string {
	return mjölnir.compiler()
}

func (mjölnir *Mjölnir) Build() error {

	tempDir := NewTempDir()
	if !mjölnir.Unclean {
		defer tempDir.Remove()
	}

	builder := &builder{mjölnir, tempDir}
	return builder.build()
}
