package StartTimer

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"time"
)

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}


func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	result := "starttime"
	iStartTime := makeTimestamp()
	context.SetOutput(result, iStartTime)
	return true, nil
}

