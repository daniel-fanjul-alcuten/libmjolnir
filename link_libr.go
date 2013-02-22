package mjölnir

import (
	"bytes"
	"os/exec"
)

type linkLibrary struct {
	mjölnir *Mjölnir
	library *CLibrary
	ofiles  []string
}

func (run *linkLibrary) exec() string {
	return run.mjölnir.libraryLinker()
}

func (run *linkLibrary) args() [][]string {
	return [][]string{}
}

func (run *linkLibrary) rargs() []string {
	args := []string{}
	args = append(args, "r", "-c", run.library.path)
	args = append(args, run.ofiles...)
	return args
}

func (run *linkLibrary) inputs() []string {
	return run.ofiles
}

func (run *linkLibrary) outputs() []string {
	return []string{run.library.path}
}

func (run *linkLibrary) run() ([]byte, []byte, error) {
	cmd := exec.Command(run.exec(), run.rargs()...)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

func (builder *builder) linkLibrary(library *CLibrary, ofiles ...string) error {
	run := &linkLibrary{builder.Mjölnir, library, ofiles}
	return builder.run(run)
}
