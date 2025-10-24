package storage

import (
	"go.etcd.io/bbolt"
)

// blocksBucket 区块存储桶名称
const blocksBucket = "blocks"

// BoltDB Bolt数据库实现
//
// 字段：
//   - db Bolt数据库实例
//
// 方法：
//   - SaveBlock 保存区块数据
//   - GetBlock 获取区块数据
//   - Close 关闭存储
//   - Iterate 遍历所有区块
type BoltDB struct {
	db *bbolt.DB
}

// NewBoltDB 创建一个新的BoltDB实例
//
// 参数：
//   - path Bolt数据库文件路径
//
// 返回值：
//   - *BoltDB BoltDB实例
//   - error 错误信息
func NewBoltDB(path string) (*BoltDB, error) {
	// 打开或创建Bolt数据库文件
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	// 创建区块存储桶
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	})
	if err != nil {
		return nil, err
	}
	return &BoltDB{db: db}, nil
}

// SaveBlock 保存区块数据
//
// 参数：
//   - hash 区块哈希值
//   - data 区块数据
//
// 返回值：
//   - error 错误信息
func (b *BoltDB) SaveBlock(hash string, data []byte) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		return bucket.Put([]byte(hash), data)
	})
}

// GetBlock 获取区块数据
//
// 参数：
//   - hash 区块哈希值
//
// 返回值：
//   - []byte 区块数据
//   - error 错误信息
func (b *BoltDB) GetBlock(hash string) ([]byte, error) {
	var data []byte
	// 从数据库中获取区块数据
	err := b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Close 关闭存储
//
// 返回值：
//   - error 错误信息
func (b *BoltDB) Close() error {
	// 关闭数据库连接
	return b.db.Close()
}

// Iterate 遍历所有区块
//
// 参数：
//   - callback 遍历回调函数
//
// 返回值：
//   - error 错误信息
func (b *BoltDB) Iterate(callback func(hash string, data []byte) error) error {
	// 遍历数据库中的所有区块
	return b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		return bucket.ForEach(func(hash, data []byte) error {
			callback(string(hash), data)
			return nil
		})
	})
}
