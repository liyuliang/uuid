package services

import (
	"time"
	"errors"
)

func Get() (id int64, err error) {

	timeout := time.After(3 * time.Second)

wait:
	for {

		select {

		case id = <-ids:
			err = nil
			break wait

		case <-timeout:
			err = errors.New("uuid generate failed, timeout")
			break wait
		}
	}

	return id, err

}

