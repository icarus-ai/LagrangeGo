// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x929B_0.proto

package oidb

type OidbSvcTrpcTcp0X929B_0_Req struct {
	GroupUin      uint32                                    `protobuf:"varint,1,opt"`
	VoiceId       string                                    `protobuf:"bytes,2,opt"`
	Text          string                                    `protobuf:"bytes,3,opt"`
	ChatType      uint32                                    `protobuf:"varint,4,opt"` // 1 voice,2 song
	ClientMsgInfo *OidbSvcTrpcTcp0X929B_0_Req_ClientMsgInfo `protobuf:"bytes,5,opt"`
	_             [0]func()
}

type OidbSvcTrpcTcp0X929B_0_Rsp struct {
	Field1  uint32   `protobuf:"varint,1,opt"` // 1 complete, 2 wait
	Field2  uint32   `protobuf:"varint,2,opt"` // 319
	Field3  uint32   `protobuf:"varint,3,opt"` // 20
	MsgInfo *MsgInfo `protobuf:"bytes,4,opt"`
	_       [0]func()
}

type OidbSvcTrpcTcp0X929B_0_Req_ClientMsgInfo struct {
	MsgRandom uint32 `protobuf:"varint,1,opt"`
	_         [0]func()
}
