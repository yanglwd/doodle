package message

import pb_gen "github.com/yanglwd/doodle/pkg/net/message/pb.gen"

func CreateControlMessage(srcId uint32, dstId uint32, msgId uint32, metas ...string) *pb_gen.ControlMessage {
	msg := &pb_gen.ControlMessage{}
	msg.SrcId = srcId
	msg.DstId = dstId
	msg.MessageId = msgId

	msg.MetaDatas = make(map[string]string)
	for i := 0; i < len(metas); i += 2 {
		msg.MetaDatas[metas[i]] = metas[i+1]
	}
	return msg
}
