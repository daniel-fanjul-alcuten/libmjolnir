package mjölnir

import (
	"bytes"
	"os/exec"
)

type compile struct {
	builder *builder
	file    *CFile
	ifile   string
	ofile   string
}

func (run *compile) exec() string {
	return run.builder.compiler()
}

func (run *compile) args() [][]string {
	args := [][]string{}
	args = append(args, []string{"-c"})
	if len(run.file.std) > 0 {
		args = append(args, []string{"-std=" + run.file.std})
	}
	return args
}

func (run *compile) rargs() []string {
	args := []string{}
	for _, arg := range run.args() {
		args = append(args, arg...)
	}
	args = append(args, "-o", run.ofile)
	args = append(args, "-x", "cpp-output", run.ifile)
	return args
}

func (run *compile) inputs() []string {
	return []string{run.ifile}
}

func (run *compile) outputs() []string {
	return []string{run.ofile}
}

func (run *compile) run() ([]byte, []byte, error) {
	cmd := exec.Command(run.exec(), run.rargs()...)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

func (builder *builder) compile(file *CFile, ifile string) (string, error) {

	ofile, err := builder.NewFileName()
	if err != nil {
		return "", err
	}

	run := &compile{builder, file, ifile, ofile}
	err = builder.run(run)
	return ofile, err
}
