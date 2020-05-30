package etcdv3

import (
	"context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// ServiceRegister 创建租约注册服务
type ServiceRegister struct {
	cli     *clientv3.Client // etcd client
	leaseID clientv3.LeaseID // 租约ID
	// 租约keepalieve相应chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string // key
	weight        string // value
}

// NewServiceRegister 新建注册服务
func NewServiceRegister(endpoints []string, addr, weigit string, lease int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	ser := &ServiceRegister{
		cli:    cli,
		key:    "/" + schema + "/" + addr,
		weight: weigit,
	}

	// 申请租约并绑定key, 设置时间 keepalive
	if err := ser.putKeyWithLease(lease); err != nil {
		return nil, err
	}

	return ser, nil
}

// 申请租约并绑定key
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	// 设置租约时间
	resp, err := s.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	// 注册服务并绑定租约
	_, err = s.cli.Put(context.Background(), s.key, s.weight, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}

	// 设置续租 定期发送需求请求
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	s.keepAliveChan = leaseRespChan
	log.Printf("Put key:%s  weight:%s  success!", s.key, s.weight)
	return nil
}

// ListenLeaseRespChan 监听 续租情况
func (s *ServiceRegister) ListenLeaseRespChan() {
	for leaseKeepResp := range s.keepAliveChan {
		log.Println("续约成功", leaseKeepResp)
	}
	log.Println("关闭续租")
}

// Close 注销服务
func (s *ServiceRegister) Close() error {
	// 撤销租约, 会同时删除与租约绑定的所有key
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Println("撤销租约")
	return s.cli.Close()
}

// func main() {
// 	var endpoints = []string{"117.51.148.112:2379"}
// 	ser, err := NewServiceRegister(endpoints, "/web/node1", "localhost:8080", 5)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// 监听续租相应chan
// 	go ser.ListenLeaseRespChan()
// 	select {}
// }
