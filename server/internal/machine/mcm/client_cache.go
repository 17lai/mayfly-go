package mcm

import (
	"context"
	"mayfly-go/pkg/pool"
)

var (
	poolGroup = pool.NewPoolGroup[*Cli]()
)

// 从缓存中获取客户端信息，不存在则回调获取机器信息函数，并新建。
// @param 机器的授权凭证名
func GetMachineCli(ctx context.Context, authCertName string, getMachine func(string) (*MachineInfo, error)) (*Cli, error) {
	pool, err := poolGroup.GetCachePool(authCertName, func() (*Cli, error) {
		mi, err := getMachine(authCertName)
		if err != nil {
			return nil, err
		}
		mi.Key = authCertName
		return mi.Conn(context.Background())
	})

	if err != nil {
		return nil, err
	}
	// 从连接池中获取一个可用的连接
	return pool.Get(ctx)
}

// 删除指定机器缓存客户端，并关闭客户端连接
func DeleteCli(id uint64) {
	for _, pool := range poolGroup.AllPool() {
		if pool.Stats().TotalConns == 0 {
			continue
		}
		conn, err := pool.Get(context.Background())
		if err != nil {
			continue
		}
		if conn.Info.Id == id {
			pool.Close()
		}
	}
}
