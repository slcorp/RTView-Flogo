
package UpdateCache

import (
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"time"
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
)

// Constants
const (
	command = "Target"
	spStartTime = "StartTime"
	params  = "params"
	result  = "result"
)

type MyActivity struct {
	metadata *activity.Metadata
}

type slPost struct {
	Metadata []slColDef `json:"metadata"`
	Data []map[string]interface{} `json:"data"`
	cols int
}

type slMetadata struct {
	Metadata []slColDef `json:"metadata"`
}

type slColDef struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (d slColDef) string() string {
	return "\"name\":\"" + d.Name + "\",\"type:\""+d.Type
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func  (m *slPost) addCol (n string, t string) {
	var colDef slColDef
	colDef.Name = n
	colDef.Type = t
	iLast := m.cols;
	if (cap(m.Metadata) == 0) {
		m.Metadata = make([]slColDef, 5, 100)
	}
	m.Metadata[iLast] = colDef
	m.cols += 1
}

// Post info contained in row data to rtview via the url
func (rowData *slPost) postRowData(url string) error {
  json, _ := json.Marshal(rowData)
  
  fmt.Println("Sending: ", url)
  fmt.Println("Post:\n", string(json))
  
  resp, err := http.Post(url, "plain/text", bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    fmt.Println("response:\n", string(body))
	return nil
}

func (a *MyActivity) updatePerformance(sURL string,  sFlowName string, iMillis int64) {
	var post slPost
	post.addCol("hostName", "string")
	post.addCol("flowName", "string")
	post.addCol("duration", "int")
	
	sHostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	
	var d map[string]interface{}
	d = make(map[string]interface{})
	d["hostName"] = sHostName
	d["flowName"] = sFlowName
	d["duration"] = iMillis

	var d2 []map[string]interface{}
	d2 = make([]map[string]interface{}, 1,1)
	d2[0] = d
	
	post.Data = d2
	bolD, _ := json.Marshal(post)

	post.postRowData(sURL)
	fmt.Println(string(bolD))	
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	sTarget := context.GetInput(command).(string)
	fmt.Println(sTarget)
	
	iVal := context.GetInput(spStartTime).(int64)
	sName := context.ActivityHost().Name()
	iEndVal := makeTimestamp() -iVal
	
    a.updatePerformance(sTarget, sName, iEndVal)

	return true, nil
}

