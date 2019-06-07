package Sleeper

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"time"
	"math/rand"

)

// Constants
const (
	spMinMillis = "Min"
	spRndMillis = "Rnd"
	params  = "params"
	result  = "result"
)


type MyActivity struct {
	metadata *activity.Metadata
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func sleepFor(iLow int, iRand int) int{
		iSleep := (iLow + rand.Intn(iRand))
		time.Sleep(time.Duration(iSleep) * time.Millisecond)
		return iSleep
}


// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	fMin := context.GetInput(spMinMillis).(float64)
	fRnd := context.GetInput(spRndMillis).(float64)
	
	iMin := int(fMin)
	iRnd := int(fRnd)
	
	sleepFor(iMin, iRnd)

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}


