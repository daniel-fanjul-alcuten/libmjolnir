package mjölnir

import (
	"fmt"
)

func (mjölnir *Mjölnir) run(run commandRun) error {

	lrun, ok := run.(localCommandRun)

	var status int
	var err error
	if ok {
		status, err = mjölnir.runLocalCommandRun(lrun)
		if err != nil {
			return err
		}

	} else {
		_, status, err = mjölnir.runCommandRunPrint("", run)
		if err != nil {
			return err
		}
	}

	if status != 0 {
		return fmt.Errorf("%v failed (%v)", run.exec(), status)
	}
	return nil
}
