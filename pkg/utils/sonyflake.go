package utils

import (
	"github.com/sony/sonyflake"

	"math/rand"
	"strconv"
	"time"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings

	st.MachineID = func() (uint16, error) {
		rand.Seed(time.Now().UnixNano())
		return uint16(rand.Intn(9999)), nil
	}

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func SonuFlakeIntID() (uint64, error) {
	return sf.NextID()
}

func SonuFlakeStringID() string {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(int(id))
}
