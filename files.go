package mjölnir

type CFile struct {
	path     string
	std      string
	includes map[string]bool
	depends  map[*CFile]bool
}

func (mjölnir *Mjölnir) CFile(path string) *CFile {
	file, ok := mjölnir.cFiles[path]
	if !ok {
		file = &CFile{path: path,
			includes: make(map[string]bool),
			depends:  make(map[*CFile]bool)}
		mjölnir.cFiles[path] = file
	}
	return file
}

func (f *CFile) SetStd(std string) *CFile {
	f.std = std
	return f
}

func (f *CFile) Includes(paths ...string) *CFile {
	for _, path := range paths {
		f.includes[path] = true
	}
	return f
}

func (f *CFile) BuildAllIncludes(includes map[string]bool) {
	for include := range f.includes {
		includes[include] = true
	}
	for depend := range f.depends {
		depend.BuildAllIncludes(includes)
	}
}

func (f *CFile) AllIncludes() []string {
	m := make(map[string]bool)
	f.BuildAllIncludes(m)
	includes := make([]string, 0, len(m))
	for include := range m {
		includes = append(includes, include)
	}
	return includes
}

func (f *CFile) DependsOn(files ...*CFile) *CFile {
	for _, file := range files {
		f.depends[file] = true
	}
	return f
}
