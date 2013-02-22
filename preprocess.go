package mjölnir

import (
	"bytes"
	"os/exec"
)

type preprocess struct {
	mjölnir *Mjölnir
	file    *CFile
	ifile   string
	mfile   string
}

func (run *preprocess) exec() string {
	return run.mjölnir.preprocessor()
}

func (run *preprocess) args() [][]string {
	args := [][]string{}
	args = append(args, []string{"-E"})
	for _, include := range run.file.AllIncludes() {
		args = append(args, []string{"-I" + include})
	}
	return args
}

func (run *preprocess) rargs() []string {
	args := []string{}
	for _, arg := range run.args() {
		args = append(args, arg...)
	}
	args = append(args, "-o", run.ifile)
	args = append(args, "-MMD", "-MF", run.mfile)
	args = append(args, run.file.path)
	return args
}

func (run *preprocess) inputs() []string {
	return []string{}
}

func (run *preprocess) outputs() []string {
	return []string{run.ifile}
}

func (run *preprocess) run() ([]byte, []byte, error) {
	cmd := exec.Command(run.exec(), run.rargs()...)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

func (run *preprocess) key() string {
	return run.file.path
}

func (run *preprocess) depends() ([]string, error) {
	return parseMFile(run.mfile)
}

func (builder *builder) preprocess(file *CFile) (string, error) {

	ifile, err := builder.NewFileName()
	if err != nil {
		return "", err
	}

	mfile, err := builder.NewFileName()
	if err != nil {
		return "", err
	}

	run := &preprocess{builder.Mjölnir, file, ifile, mfile}
	err = builder.run(run)
	return ifile, err
}
