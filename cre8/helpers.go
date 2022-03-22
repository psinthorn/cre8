package cre8

import "os"

func (c *Cre8) createFolderIfNotExist(folderPath string) error {
	// ในการสร้าง folder เราต้องการทราบอะไรบ้าง
	// 1. folder mode
	const mode = 0755

	// ตรวจสอบ path ของ folder หากมี error ให้ return return error กลับไป หากไม่มี error จึงให้สร้าง folder
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, mode); err != nil {
			return err
		}
	}

	return nil
}
