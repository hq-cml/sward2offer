package _019_del_back_n_node

import (
	"github.com/hq-cml/sward2offer/common"
	"log"
	"testing"
)

func TestDelBackN(t *testing.T) {
	var l *common.ListNode
	l = l.PushNode(1)
	l = l.PushNode(2)
	l = l.PushNode(3)
	l.Foreach(common.NodePrint)

	//l, err := DelBackN(l, 1)
	//l, err := DelBackN(l, 2)
	//l, err := DelBackN(l, 3)
	l, err := DelBackN(l, 4)
	if err != nil {
		log.Println(err)
	}
	l.Foreach(common.NodePrint)
}
