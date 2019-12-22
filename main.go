// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/user/spy_go/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement pong.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedSpyServer
}

// SayHello implements pong.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SendStep(ctx context.Context, in *pb.StepUpdate) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetValue())
	return &pb.Reply{Message: "Received"}, nil
}

func (s *server) SendAccelerometer(ctx context.Context, in *pb.AccelerometerUpdate) (*pb.Reply, error) {
	log.Printf("Accelerometer: %v %v %v", in.GetX(), in.GetY(), in.GetZ())
	return &pb.Reply{Message: "Received accelerometer"}, nil
}

func (s *server) SendUserAccelerometer(ctx context.Context, in *pb.UserAccelerometerUpdate) (*pb.Reply, error) {
	log.Printf("UserAccelerometer: %v %v %v", in.GetX(), in.GetY(), in.GetZ())
	return &pb.Reply{Message: "Received user accelerometer"}, nil
}

func (s *server) SendGyroscope(ctx context.Context, in *pb.GyroscopeUpdate) (*pb.Reply, error) {
	log.Printf("Gyroscope: %v %v %v", in.GetX(), in.GetY(), in.GetZ())
	return &pb.Reply{Message: "Received gyroscope"}, nil
}

func (s *server) SendDeviceInfo(ctx context.Context, in *pb.DeviceInfo) (*pb.Reply, error) {
	log.Printf("Model: %v", in.GetModel())
	return &pb.Reply{Message: "Received model"}, nil
}

func (s *server) SendBatteryLevel(ctx context.Context, in *pb.BatteryLevel) (*pb.Reply, error) {
	log.Printf("Battery level: %v", in.GetValue())
	return &pb.Reply{Message: "Received battery level"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Server started")
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterSpyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
