package mjölnir

import (
	. "github.com/daniel-fanjul-alcuten/libmjolnir/utils"
)

type builder struct {
	*Mjölnir
	*TempDir
}

func (builder *builder) compile_file(ofiles map[*CFile]string, file *CFile) (string, error) {
	ofile, ok := ofiles[file]
	if !ok {
		ifile, err := builder.preprocess(file)
		if err != nil {
			return "", err
		}
		ofile, err = builder.compile(file, ifile)
		if err != nil {
			return "", err
		}
		ofiles[file] = ofile
	}
	return ofile, nil
}

func (builder *builder) compile_files_map(ofiles map[*CFile]string, mofiles map[string]bool, files map[*CFile]bool) error {
	for file := range files {
		err := builder.compile_files_map(ofiles, mofiles, file.depends)
		if err != nil {
			return err
		}
		ofile, err := builder.compile_file(ofiles, file)
		if err != nil {
			return err
		}
		mofiles[ofile] = true
	}
	return nil
}

func (builder *builder) compile_files_slice(ofiles map[*CFile]string, files map[*CFile]bool) ([]string, error) {
	mofiles := make(map[string]bool, len(files))
	err := builder.compile_files_map(ofiles, mofiles, files)
	if err != nil {
		return nil, err
	}
	lofiles := make([]string, 0, len(mofiles))
	for ofile := range mofiles {
		lofiles = append(lofiles, ofile)
	}
	return lofiles, nil
}

func (builder *builder) build() error {

	if builder.Mjölnir.LocalCache != nil {
		builder.Mjölnir.LocalCache.read()
		defer builder.Mjölnir.LocalCache.write()
	}

	ofiles := make(map[*CFile]string)

	for _, library := range builder.Mjölnir.cLibraries {
		lofiles, err := builder.compile_files_slice(ofiles, library.files)
		if err != nil {
			return err
		}
		err = builder.linkLibrary(library, lofiles...)
		if err != nil {
			return err
		}
	}

	for _, executable := range builder.Mjölnir.cExecutables {
		lofiles, err := builder.compile_files_slice(ofiles, executable.files)
		if err != nil {
			return err
		}
		err = builder.linkExecutable(executable, lofiles...)
		if err != nil {
			return err
		}
	}

	return nil
}
