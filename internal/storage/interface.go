package storage

// Storage 存储接口
//
// 方法：
//   - SaveBlock 保存区块数据
//   - GetBlock 获取区块数据
//   - Close 关闭存储
//   - Iterate 遍历所有区块
type Storage interface {
	SaveBlock(hash string, data []byte) error
	GetBlock(hash string) ([]byte, error)
	Close() error
	Iterate(callback func(hash string, data []byte) error) error
}
