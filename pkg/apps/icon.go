package apps

import (
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/services"
	"CampusRecruitment/pkg/types/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

func IconFileUpload(file *multipart.FileHeader) (interface{}, error) {
	if file.Size > consts.IconFileSize {
		return nil, errors.ErrFileTooLarge.WithStatus(http.StatusRequestEntityTooLarge)
	}

	fp, err := file.Open()
	if err != nil {
		return nil, errors.ErrFileIoError.WithCause(err)
	}

	content, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, errors.ErrFileIoError.WithCause(err)
	}

	storage := services.GetStorage()
	savePath, err := services.SaveIcon(storage, file.Filename, content)
	if err != nil {
		return nil, err
	}
	urlPath := "api/v1/icons?path=" + savePath
	fmt.Println(savePath)
	resp := gin.H{"path": savePath, "urlPath": urlPath}
	return resp, nil
}

func ReadIconFile(c *ctx.Context, w io.Writer, path string) error {
	storage := services.GetStorage()
	path = strings.TrimLeft(path, "/")
	data, err := storage.Read(path)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		c.Logger().Warnf("write response error: %v", err)
	}
	return nil
}
