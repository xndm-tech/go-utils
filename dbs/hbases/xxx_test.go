package hbases

import (
	"fmt"
	"testing"
)

func TestHbaseGets(t *testing.T) {
	cli, _ := createClient()
	fmt.Println(connPool)
	t.Log(cli)
	t.Log(ping(cli))
}
