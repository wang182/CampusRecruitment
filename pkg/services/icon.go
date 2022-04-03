package services

import (
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/infra/log"
	"CampusRecruitment/pkg/types/errors"
	"CampusRecruitment/pkg/utils"
	"path"
	"path/filepath"
)

func SaveIcon(storage FileStorage, filename string, content []byte) (string, error) {
	md5 := utils.Md5(content)
	ext := filepath.Ext(filename)
	if ext == "" {
		ext = ".png"
	}
	filePath := path.Join(consts.IconFilePrefix, md5[:2], md5[2:4], md5+ext)

	// 文件己存在则直接返回
	isExist, err := storage.IsExists(filePath)
	if err != nil {
		log.Get().Warnf("check file exists error: %v", err)
	} else if isExist {
		return filePath, nil
	}
	_, err = storage.Write(filePath, content)
	if err != nil {
		return "", errors.ErrFileIoError.WithCause(err)
	}
	return filePath, nil
}
