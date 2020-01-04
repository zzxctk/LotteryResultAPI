//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: cs

package MyGame

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for MonsterStorage service
type MonsterStorageClient interface{
  Store(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* Weapon, error)  
}

type monsterStorageClient struct {
  cc *grpc.ClientConn
}

func NewMonsterStorageClient(cc *grpc.ClientConn) MonsterStorageClient {
  return &monsterStorageClient{cc}
}

func (c *monsterStorageClient) Store(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* Weapon, error) {
  out := new(Weapon)
  err := grpc.Invoke(ctx, "/MyGame.MonsterStorage/Store", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for MonsterStorage service
type MonsterStorageServer interface {
  Store(context.Context, *Monster) (*flatbuffers.Builder, error)  
}

func RegisterMonsterStorageServer(s *grpc.Server, srv MonsterStorageServer) {
  s.RegisterService(&_MonsterStorage_serviceDesc, srv)
}

func _MonsterStorage_Store_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(Monster)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(MonsterStorageServer).Store(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/MyGame.MonsterStorage/Store",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(MonsterStorageServer).Store(ctx, req.(* Monster))
  }
  return interceptor(ctx, in, info, handler)
}


var _MonsterStorage_serviceDesc = grpc.ServiceDesc{
  ServiceName: "MyGame.MonsterStorage",
  HandlerType: (*MonsterStorageServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Store",
      Handler: _MonsterStorage_Store_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}
