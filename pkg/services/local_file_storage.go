package services

import (
	"CampusRecruitment/pkg/config"
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/infra/log"
	"CampusRecruitment/pkg/types/errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FileStorage interface {

	// Read 读取文件内容，如果文件不存在则返回 error types.ErrFileNotExists
	// path 参数为相对路径，不带 "/" 前缀
	Read(path string) ([]byte, error)

	// Write 全量写入文件内容，如果文件不存在则创建
	// path 参数为相对路径，不带 "/" 前缀
	Write(path string, data []byte) (int, error)

	IsExists(path string) (bool, error)
}

type LocalFileStorage struct {
	FileName       string `json:"fileName"`       // 文件名称
	FilePath       string `json:"filePath"`       // 图标本地存储路径
	FilePathPreFix string `json:"filePathPreFix"` // 图片存储路径前缀 例如 /var/registry-storage/uploads
}

func (s *LocalFileStorage) Read(filepath string) ([]byte, error) {
	bs, err := ioutil.ReadFile(path.Join(s.FilePathPreFix, filepath))
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil, errors.ErrFileNotExists
		} else {
			log.Get().WithField("file", filepath).Warnf("read file error: %v", err)
			return nil, errors.ErrFileIoError
		}
	}
	return bs, err
}

func (s *LocalFileStorage) Write(filepath string, data []byte) (int, error) {
	fullPath := s.fullPath(filepath)
	dirPath := path.Dir(fullPath)
	isExist, err := s.IsExists(dirPath)
	if err != nil {
		return 0, err
	}
	if !isExist {
		if err = s.mkdir(dirPath); err != nil {
			return 0, err
		}
	}
	err = os.WriteFile(fullPath, data, 0666) // nolint:gosec
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

func (s *LocalFileStorage) mkdir(filePath string) error {
	return os.MkdirAll(filePath, os.ModePerm)
}

func (s *LocalFileStorage) IsExists(filepath string) (bool, error) {
	fullPath := s.fullPath(filepath)
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *LocalFileStorage) fullPath(filePath string) string {
	return path.Join(s.FilePathPreFix, filePath)
}

func GetStorage() FileStorage {
	return &LocalFileStorage{
		FilePathPreFix: filepath.Join(config.Get().Storage, consts.UploadDirName),
	}
}
