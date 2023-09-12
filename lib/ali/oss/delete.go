package oss

// DeleteObject 删除单个对象，只能删除空文件夹
func DeleteObject(key string) error {
	err := Client.b.DeleteObject(key)
	return err
}
