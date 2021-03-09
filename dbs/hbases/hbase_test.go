package hbases

import (
	"fmt"
	"testing"
	"time"

	"github.com/popeyeio/gohbase/gen/hbase"
	"github.com/popeyeio/gohbase/pool"
)

func NewHbasePool() pool.Pool {
	opts := []pool.Option{
		pool.WithAddrs("hb-bp1m205ju7p5l24ey-001.hbase.rds.aliyuncs.com:9099"),
		pool.WithUpdatePickerInterval(time.Second * 10),
		pool.WithSocketTimeout(time.Second * 5),
		pool.WithMaxActive(8),
		pool.WithMaxIdle(8),
		pool.WithIdleTimeout(time.Second * 5),
		pool.WithCleanUpInterval(time.Second * 30),
		pool.WithBlockMode(true),
	}
	return pool.NewPool(opts...)
}

//func TestHbaseConnectionFromConfig2(t *testing.T) {
//	hbasePool := NewHbasePool()
//	defer hbasePool.Close()
//
//	hbaseClient, err := hbasePool.Get()
//	if err != nil {
//		fmt.Printf("Get error - %v\n", err)
//		return
//	}
//	defer hbaseClient.Close()
//
//	events, err := ScanEvents(hbaseClient, "1", "2")
//	if err != nil {
//		fmt.Printf("ScanEvents error - %v\n", err)
//		return
//	}
//	for _, e := range events {
//		fmt.Printf("events:%v\n", e)
//	}
//}

func TestHbaseConnectionFromConfig3(t *testing.T) {
	hbasePool := NewHbasePool()
	defer hbasePool.Close()

	hbaseClient, err := hbasePool.Get()
	if err != nil {
		fmt.Printf("Get error - %v\n", err)
		return
	}
	defer hbaseClient.Close()

	events, err := getEvents(hbaseClient, "1", "2")
	if err != nil {
		fmt.Printf("ScanEvents error - %v\n", err)
		return
	}
	for _, e := range events {
		fmt.Printf("events:%v\n", e)
	}
}

type Event struct {
	EventTime string
	Publisher string
}

func (e *Event) ParseFromHbase(r *hbase.TRowResult_) {
	e.EventTime = GetColumn(r, "info:day")
	e.Publisher = GetColumn(r, "info:day")
}

func GetColumn(r *hbase.TRowResult_, c string) (v string) {
	if tcell := r.Columns[c]; tcell != nil {
		v = string(tcell.Value)
	}
	return
}

func ScanEvents(cli pool.Client, startRow, stopRow string) ([]*Event, error) {
	tscan := &hbase.TScan{
		StartRow: hbase.Text(startRow),
		StopRow:  hbase.Text(stopRow),
	}

	scanID, err := cli.ScannerOpenWithScan("recommend_samh:profile_user", tscan, nil)
	if err != nil {
		return nil, err
	}
	defer cli.ScannerClose(scanID)

	var events []*Event
	for {
		results, err := cli.ScannerGetList(scanID, 2)
		if err != nil {
			return nil, err
		}
		if len(results) == 0 {
			return events, nil
		}

		fmt.Println(results)
		fmt.Println(len(results))
	}
}

func getEvents(cli pool.Client, startRow, stopRow string) ([]*Event, error) {

	scanID, err := cli.GetRowWithColumns("recommend_samh:profile_user", "00001466fbb954f0:comic", []string{"info:click_n7d_day"}, nil)
	if err != nil {
		return nil, err
	}

	//var events []*Event
	fmt.Println(len(scanID))
	for _, s := range scanID {
		fmt.Println(GetColumn(s, "info:click_n7d_day"))
	}
	return nil, err
}
