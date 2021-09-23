package entity

import (
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	uid, err := uuid.Parse(id)
	return ID(uid), err
}

type Time int64

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strconv.FormatInt(int64(t), 10))), nil
}

func (t *Time) UnmarshalJSON(s []byte) error {
	tstr := string(s)
	tstr = tstr[1 : len(tstr)-1]
	log.Println(tstr)

	utc, _ := time.LoadLocation("UTC")

	tint, err := time.ParseInLocation("2006-01-02T15:04:05.999999Z", tstr, utc)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	*(*int64)(t) = tint.Unix()

	return nil
}

//func (t *Time) Time() time.Time {
//	return time.Time(*t).UTC()
//}
