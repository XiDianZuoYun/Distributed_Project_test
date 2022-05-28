package paxos

import "reflect"

type BallotNumber struct {
	proposal_id_ int64
	node_id_     int64
}

func CreateBallotNumber(proposal_id int64, node_id int64) *BallotNumber {
	return &BallotNumber{
		proposal_id_: proposal_id,
		node_id_:     node_id,
	}
}

func Equal(bn_1 *BallotNumber, bn_2 *BallotNumber) bool {
	return reflect.DeepEqual(bn_1, bn_2)
}

func Greater(bn_1 *BallotNumber, bn_2 *BallotNumber) bool {
	if bn_1.proposal_id_ > bn_2.proposal_id_ {
		return true
	}
	if bn_1.proposal_id_ == bn_2.proposal_id_ {
		return bn_1.node_id_ > bn_2.node_id_
	}
	return false
}

func NotSmaller(bn_1 *BallotNumber, bn_2 *BallotNumber) bool {
	return Equal(bn_1, bn_2) || Greater(bn_2, bn_2)
}

func (bn *BallotNumber) Reset() {
	bn.node_id_ = 0
	bn.proposal_id_ = 0
}

func (bn *BallotNumber) IsNull() bool {
	return bn.proposal_id_ == 0
}

// 允许广播或者定点发送信息的IO结构，非常简单
type NetHelper struct {
	
}

func (n* NetHelper) 