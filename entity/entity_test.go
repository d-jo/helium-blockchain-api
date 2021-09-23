package entity

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestNewID(t *testing.T) {
	id := NewID()

	if len(id.String()) != 36 {
		t.Logf("failed length test")
		t.FailNow()
	}
}

func TestParseID(t *testing.T) {
	id_str := "c639bef2-ff6d-4cf9-9ea5-b65bc7dbc01f"

	uid, err := ParseID(id_str)

	if err != nil {
		t.Logf("failed to parse uuid: %s", err.Error())
		t.FailNow()
	}

	if uid.String() != id_str {
		t.FailNow()
	}
}

func TestParseIDFail(t *testing.T) {
	id_str := "c639bef2-ff6d-4cf9-9ea5-b5bc7dbc01f"

	_, err := ParseID(id_str)

	if err == nil {
		t.Logf("failed fail to parse id")
		t.FailNow()
	}
}

func TestTimeMarshal(t *testing.T) {
	unixTime := time.Now().Unix()

	timeObj := Time(unixTime)

	marshaled, err := json.Marshal(timeObj)

	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	if string(marshaled) != strconv.Quote(strconv.FormatInt(unixTime, 10)) {
		t.FailNow()
	}
}

func TestTimeUnmarshal(t *testing.T) {
	//tValue := "2021-09-19T15:53:42.000000Z"
	utc, _ := time.LoadLocation("UTC")
	tNow := time.Now().In(utc) //.Unix()
	tNowUnix := tNow.Unix()
	//tv := time.Unix(tNow, 0)
	tBuf := tNow.Format("2006-01-02T15:04:05.999999Z")
	log.Println(tBuf)
	tValue := fmt.Sprintf(`{
		"time": "%s"
	}`, tBuf)
	type tTarget struct {
		TS Time `json:"time"`
	}

	tt := &tTarget{}

	err := json.Unmarshal([]byte(tValue), &tt)
	//err := (&tTarget).UnmarshalJSON([]byte(tValue))

	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	log.Println(tt.TS)
	log.Println(int64(tt.TS) - tNowUnix)
	if int64(tt.TS) != tNowUnix {
		t.Logf("time unmarshal does not match %d != %d", int64(tt.TS), tNowUnix)
		t.FailNow()
	}

}

func TestTimeUnmarshalFail(t *testing.T) {
	tBuf := "sdfnkbj" //tNow.Format("2006-01-02T15:04:05.999999Z")
	log.Println(tBuf)
	tValue := fmt.Sprintf(`{
		"time": "%s"
	}`, tBuf)
	type tTarget struct {
		TS Time `json:"time"`
	}

	tt := &tTarget{}

	err := json.Unmarshal([]byte(tValue), &tt)
	//err := (&tTarget).UnmarshalJSON([]byte(tValue))

	if err == nil {
		t.Logf(err.Error())
		t.FailNow()
	}

}
