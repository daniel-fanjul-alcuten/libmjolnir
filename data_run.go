package mjölnir

import (
	"io/ioutil"
	"log"
	"os/exec"
	"syscall"
)

type commandRun interface {
	exec() string
	args() [][]string
	rargs() []string
	inputs() []string
	outputs() []string
	run() ([]byte, []byte, error)
}

func (mjölnir *Mjölnir) runCommandRun(id string, run commandRun) (bool, []byte, []byte, int, error) {

	var key *commandKey
	if mjölnir.DataCache != nil {

		inputs := make([]string, len(run.inputs()))
		for i, input := range run.inputs() {
			data, err := ioutil.ReadFile(input)
			if err != nil {
				return true, nil, nil, 0, err
			}
			inputs[i] = hashFile(data)
		}

		key = &commandKey{id, run.exec(), run.args(), inputs}

		value, err := getCommand(mjölnir.DataCache, key)
		if err != nil {
			return true, nil, nil, 0, err
		}

		if value != nil {

			datas := make([][]byte, len(value.Files))
			for i, hash := range value.Files {

				data, err := getFile(mjölnir.DataCache, hash)
				if err != nil {
					return true, value.Stdout, value.Stderr, value.Status, err
				}

				if data == nil {
					datas = nil
					break
				}

				datas[i] = data
			}

			if datas != nil {
				for i, data := range datas {
					err = ioutil.WriteFile(run.outputs()[i], data, 0666)
					if err != nil {
						return true, value.Stdout, value.Stderr, value.Status, err
					}
				}
				return true, value.Stdout, value.Stderr, value.Status, nil
			}
		}
	}

	stdout, stderr, err := run.run()
	status := 0
	if exiterr, ok := err.(*exec.ExitError); ok {
		sys := exiterr.ProcessState.Sys()
		if s, ok := sys.(syscall.WaitStatus); ok {
			status = s.ExitStatus()
			err = nil
		}
	}
	if err != nil {
		return false, stdout, stderr, status, err
	}

	if mjölnir.DataCache != nil {

		var files []string
		if status == 0 {
			files = make([]string, len(run.outputs()))
			for i, file := range run.outputs() {

				data, err := ioutil.ReadFile(file)
				if err != nil {
					return false, stdout, stderr, status, err
				}

				hash := hashFile(data)
				err = setFile(mjölnir.DataCache, hash, data)
				if err != nil {
					return false, stdout, stderr, status, err
				}

				files[i] = hash
			}
		} else {
			files = []string{}
		}

		value := &commandValue{stdout, stderr, files, status}
		err = setCommand(mjölnir.DataCache, key, value)
		return false, stdout, stderr, status, err
	}

	return false, stdout, stderr, status, nil
}

func (mjölnir *Mjölnir) runCommandRunPrint(id string, run commandRun) (bool, int, error) {

	hit, stdout, stderr, status, err := mjölnir.runCommandRun(id, run)

	if status != 0 || (!hit && mjölnir.Verbose > 0) || (hit && mjölnir.Verbose > 1) {
		var symbol string
		if hit {
			symbol = "-"
		} else {
			symbol = "+"
		}
		line := []interface{}{symbol, run.exec()}
		for _, arg := range run.rargs() {
			line = append(line, arg)
		}
		log.Println(line...)
		log.Print(string(stdout))
		log.Print(string(stderr))
	}

	return hit, status, err
}
