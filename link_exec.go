package mjölnir

import (
	"bytes"
	"os"
	"os/exec"
)

type linkExecutable struct {
	mjölnir    *Mjölnir
	executable *CExecutable
	ofiles     []string
}

func (run *linkExecutable) exec() string {
	return run.mjölnir.executableLinker()
}

func (run *linkExecutable) args() [][]string {
	return [][]string{}
}

func (run *linkExecutable) rargs() []string {
	args := []string{}
	args = append(args, "-o", run.executable.path)
	args = append(args, run.ofiles...)
	return args
}

func (run *linkExecutable) inputs() []string {
	return run.ofiles
}

func (run *linkExecutable) outputs() []string {
	return []string{run.executable.path}
}

func (run *linkExecutable) run() ([]byte, []byte, error) {
	cmd := exec.Command(run.exec(), run.rargs()...)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

func (builder *builder) linkExecutable(executable *CExecutable, ofiles ...string) error {
	run := &linkExecutable{builder.Mjölnir, executable, ofiles}
	err := builder.run(run)
	if err != nil {
		return err
	}
	return os.Chmod(executable.path, 0755)
}
