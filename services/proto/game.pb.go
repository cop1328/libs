// Code generated by protoc-gen-go.
// source: game.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type Game struct {
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto1.CompactTextString(m) }
func (*Game) ProtoMessage()    {}

type Game_Nil struct {
}

func (m *Game_Nil) Reset()         { *m = Game_Nil{} }
func (m *Game_Nil) String() string { return proto1.CompactTextString(m) }
func (*Game_Nil) ProtoMessage()    {}

type Game_Packet struct {
	Uid     int32  `protobuf:"varint,1,opt" json:"Uid,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,proto3" json:"Content,omitempty"`
}

func (m *Game_Packet) Reset()         { *m = Game_Packet{} }
func (m *Game_Packet) String() string { return proto1.CompactTextString(m) }
func (*Game_Packet) ProtoMessage()    {}

func init() {
}

// Client API for GameService service

type GameServiceClient interface {
	Packet(ctx context.Context, opts ...grpc.CallOption) (GameService_PacketClient, error)
	Notify(ctx context.Context, in *Game_Nil, opts ...grpc.CallOption) (GameService_NotifyClient, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) Packet(ctx context.Context, opts ...grpc.CallOption) (GameService_PacketClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GameService_serviceDesc.Streams[0], c.cc, "/proto.GameService/Packet", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServicePacketClient{stream}
	return x, nil
}

type GameService_PacketClient interface {
	Send(*Game_Packet) error
	Recv() (*Game_Packet, error)
	grpc.ClientStream
}

type gameServicePacketClient struct {
	grpc.ClientStream
}

func (x *gameServicePacketClient) Send(m *Game_Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServicePacketClient) Recv() (*Game_Packet, error) {
	m := new(Game_Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) Notify(ctx context.Context, in *Game_Nil, opts ...grpc.CallOption) (GameService_NotifyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GameService_serviceDesc.Streams[1], c.cc, "/proto.GameService/Notify", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceNotifyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameService_NotifyClient interface {
	Recv() (*Game_Packet, error)
	grpc.ClientStream
}

type gameServiceNotifyClient struct {
	grpc.ClientStream
}

func (x *gameServiceNotifyClient) Recv() (*Game_Packet, error) {
	m := new(Game_Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GameService service

type GameServiceServer interface {
	Packet(GameService_PacketServer) error
	Notify(*Game_Nil, GameService_NotifyServer) error
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_Packet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).Packet(&gameServicePacketServer{stream})
}

type GameService_PacketServer interface {
	Send(*Game_Packet) error
	Recv() (*Game_Packet, error)
	grpc.ServerStream
}

type gameServicePacketServer struct {
	grpc.ServerStream
}

func (x *gameServicePacketServer) Send(m *Game_Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServicePacketServer) Recv() (*Game_Packet, error) {
	m := new(Game_Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GameService_Notify_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Game_Nil)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServiceServer).Notify(m, &gameServiceNotifyServer{stream})
}

type GameService_NotifyServer interface {
	Send(*Game_Packet) error
	grpc.ServerStream
}

type gameServiceNotifyServer struct {
	grpc.ServerStream
}

func (x *gameServiceNotifyServer) Send(m *Game_Packet) error {
	return x.ServerStream.SendMsg(m)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Packet",
			Handler:       _GameService_Packet_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Notify",
			Handler:       _GameService_Notify_Handler,
			ServerStreams: true,
		},
	},
}