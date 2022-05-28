package paxos

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
type ProposerState struct {
	proposal_id_ int64 // paxos协议中的proposal_id
	value_ string // 本次协议需要确定的值
	highest_other_proposal_id_ int64 // 其他node最大的proposal_id
	highest_other_pre_accept_ballot_ BallotNumber // 是否存在其他node完成了accept
}

func CreateProposerState(node_id int64) *ProposerState {
	return &ProposerState{
		proposal_id_: 0,
		value_: "",
		highest_other_proposal_id_: 0,
		highest_other_pre_accept_ballot_: *CreateBallotNumber(0,node_id),
	}
}

// ProposerState接收新的值，需要bn更大
func AddPreAcceptValue (p *ProposerState)(value string,bn* BallotNumber){
	if(Greater(bn,&p.highest_other_pre_accept_ballot_)){
		p.highest_other_pre_accept_ballot_ = *bn
		p.value_ = value
	}
}

type Proposer struct {
	proposer_state_ ProposerState
	server_cnt_ int32 // 服务器数量 todo(改为读取配置文件)
	is_preparing_ bool // 
	is_accepting_ bool
	logger_ zap.Logger// 日志模块
}

func CreateProposer (server_cnt int32,node_id int64) *Proposer {
	// 日志相关设置
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	file, _ := os.OpenFile("/tmp/paxos_log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	sync_file := zapcore.AddSync(file)
    sync_console := zapcore.AddSync(os.Stderr)
	sync := zapcore.NewMultiWriteSyncer(sync_console,sync_file)
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	return &Proposer{
		server_cnt_: server_cnt,
		is_preparing_: false,
		is_accepting_:  false,
		proposer_state_: *CreateProposerState(node_id),
		logger_: zap.New(core),
	}
}

