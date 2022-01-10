package utils

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

func PkgError(err error) error {
	if err == nil {
		return err
	}

	id := xid.New().String()
	Logger.Error(fmt.Errorf("ErrorID: %s error: %s", id, err.Error()))

	return errors.New(fmt.Sprintf("LogID: %s 系统更新中 请稍候再试", id))
}
