package firing

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"relay/line"
	"strconv"
	"time"
)

type RecvDataer interface {
	Set(postMap map[string]string)
	GetValueMap(re *http.Request, target ...string) (map[string]string, error)
	CallCommand(token string)
}

type recvData struct {
	dataTime     time.Time
	email        string
	affiliations string
	group        int
	formatImg    int
	cnnModel     int
}

func (r *recvData) Set(postMap map[string]string) {
	//Mon Jan 16 14:35:34 GMT+08:00 2023
	timeObj, err := time.Parse("2006-01-02 15:04:05 -0700", postMap["time"])
	if err != nil {
		fmt.Println(err)
	}
	r.dataTime = timeObj
	r.email = postMap["email"]
	r.affiliations = postMap["aff"]

	floatGroup, _ := strconv.ParseFloat(postMap["group"], 32)
	r.group = int(floatGroup)

	floatFormat, _ := strconv.ParseFloat(postMap["format"], 32)
	r.formatImg = int(floatFormat)

	floatCnn, _ := strconv.ParseFloat(postMap["cnn"], 32)
	r.cnnModel = int(floatCnn)

}
func (r *recvData) GetValueMap(re *http.Request, target ...string) (map[string]string, error) {
	err := re.ParseForm()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bufMap := make(map[string]string)
	for _, v := range target {
		bufMap[v] = re.FormValue(v)
	}
	return bufMap, nil
}
func (r *recvData) CallCommand(token string) {
	out, err := exec.Command("D:\\OLO_CNN\\iGetG_OLO_Form.exe", r.email, fmt.Sprint(r.group), fmt.Sprint(r.formatImg), fmt.Sprint(r.cnnModel)).Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
	var line line.Liner = line.NewNotify()
	ymdhis := r.dataTime.Format("2006-01-02 15:04:05")
	line.Config(token)
	msg := fmt.Sprintf("\n觸發時間:%s\n電子郵件:%s\n來源機構:%s", ymdhis, r.email, r.affiliations)
	rep, err := line.Trigger(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rep)

}
func New() *recvData {
	return &recvData{}
}
