package jsons

/*
有关json读取保存的封装
*/
import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/cihub/seelog"
)

type JsonStruct struct {
}

func SaveJsonFile(filename string, saveData string) (bool, error) {
	//将string保存为文件
	err := ioutil.WriteFile(filename, []byte(saveData), os.ModeAppend)
	if err != nil {
		seelog.Debug("SaveFile fail")
		return false, err
	}
	return true, nil
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (self *JsonStruct) JsonLoad(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, v)
	if err != nil {
		return
	}
}
