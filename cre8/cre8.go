package cre8

const version = "1.0.0"

// New เพื่อตรวจสอบ folder path ที่ทำงานอยู่ให้ปัจจุบัน โดยจะคืนค่ากลับมาเป็น string
func (c *Cre8) New(rootPath string) error {
	folderConfig := initFolders{
		rootPath:    rootPath,
		folderNames: []string{"models", "views", "controllers", "data", "public", "static", "middleware", "logs", "tmp"},
	}

	if err := c.Init(folderConfig); err != nil {
		return err
	}

	return nil
}

func (c *Cre8) Init(fd initFolders) error {
	rootPath := fd.rootPath
	for _, folderName := range fd.folderNames {
		err := c.createFolderIfNotExist(rootPath + "/" + folderName)
		if err != nil {
			return err
		}
	}
	return nil
}
