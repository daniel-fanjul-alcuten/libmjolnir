package mjölnir

type CExecutable struct {
	path  string
	files map[*CFile]bool
}

func (mjölnir *Mjölnir) CExecutable(path string) *CExecutable {
	executable, ok := mjölnir.cExecutables[path]
	if !ok {
		executable = &CExecutable{path, make(map[*CFile]bool)}
		mjölnir.cExecutables[path] = executable
	}
	return executable
}

func (e *CExecutable) Link(files ...*CFile) {
	for _, file := range files {
		e.files[file] = true
	}
}
