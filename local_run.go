package mjölnir

import (
	"fmt"
	"io"
	"os"
	"time"
)

type localCommandRun interface {
	commandRun
	key() string
	depends() ([]string, error)
}

func (mjölnir *Mjölnir) runLocalCommandRun(run localCommandRun) (int, error) {

	var id string
	var stime time.Time

	if mjölnir.LocalCache != nil {
		cdata, ok := mjölnir.LocalCache.data[run.key()]
		if ok {
			for _, file := range cdata.Depends {

				info, err := os.Stat(file)
				if err != nil {
					return 0, err
				}

				ftime := info.ModTime()
				if !cdata.Time.After(ftime) {
					ok = false
					break
				}
			}
		}

		if ok {
			id = cdata.Id
			stime = cdata.Time
		}
	}

	if id == "" && mjölnir.DataCache != nil {

		f, err := os.Open("/dev/urandom")
		if err != nil {
			return 0, err
		}

		b := make([]byte, 16)
		_, err = io.ReadFull(f, b)
		if err != nil {
			f.Close()
			return 0, err
		}

		err = f.Close()
		if err != nil {
			return 0, err
		}

		id = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
		stime = time.Now()
	}

	hit, status, err := mjölnir.runCommandRunPrint(id, run)
	if err != nil {
		return status, err
	}

	if mjölnir.LocalCache != nil && !hit {
		depends, err := run.depends()
		if err != nil {
			return status, err
		}
		mjölnir.LocalCache.data[run.key()] = lcdata{id, stime, depends}
	}
	return status, nil
}
