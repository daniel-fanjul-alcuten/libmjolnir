package mjölnir

type CLibrary struct {
	path  string
	files map[*CFile]bool
}

func (mjölnir *Mjölnir) CLibrary(path string) *CLibrary {
	library, ok := mjölnir.cLibraries[path]
	if !ok {
		library = &CLibrary{path, make(map[*CFile]bool)}
		mjölnir.cLibraries[path] = library
	}
	return library
}

func (l *CLibrary) Link(files ...*CFile) {
	for _, file := range files {
		l.files[file] = true
	}
}
