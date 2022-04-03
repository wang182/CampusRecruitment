package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/openpgp"
)

func init() {
	uuid.SetRand(rand.Reader)
}

func GPGKeyId(pubKey string) (string, error) {
	reader := bytes.NewBufferString(pubKey)
	entityList, err := openpgp.ReadArmoredKeyRing(reader)
	if err != nil {
		return "", err
	}
	for _, e := range entityList {
		return e.PrimaryKey.KeyIdString(), nil
	}
	return "", fmt.Errorf("not found")
}

func ScanFileLines(path string, fn func(line string) error) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		if err := fn(scanner.Text()); err != nil {
			return err
		}
	}
	return nil
}

func FileSha256(path string) (string, error) {
	sha := sha256.New()
	fp, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	if _, err := io.Copy(sha, fp); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha.Sum(nil)), nil
}

func ReadCheckSums(checksumFile string) (map[string]string, error) {
	rs := make(map[string]string)
	err := ScanFileLines(checksumFile, func(line string) error {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return fmt.Errorf("invalid checksum line: '%s'", line)
		}
		checksum, filename := fields[0], fields[1]
		rs[filename] = checksum
		return nil
	})
	return rs, err
}

func FileExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func ProviderPackageName(typ, version, os, arch string) string {
	return fmt.Sprintf("terraform-provider-%s_%s_%s_%s.zip", typ, version, os, arch)
}

func WriteFileAsJson(file string, obj interface{}) error {
	if bs, err := json.Marshal(obj); err != nil {
		return err
	} else {
		if err := os.WriteFile(file, bs, 0644); err != nil {
			return err
		}
	}
	return nil
}

// 计算文件md5值
func Md5(content []byte) string {
	h := md5.New()
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}

func GenId() string {
	id := [16]byte(uuid.New())
	rv := base64.RawStdEncoding.EncodeToString(id[:])
	rv = strings.ReplaceAll(rv, "+", "")
	rv = strings.ReplaceAll(rv, "/", "")
	return rv
}

func GenUUID() string {
	return uuid.New().String()
}

func IsHostname(h string) bool {
	return strings.Contains(h, ".")
}

// 将 git tag 转为 version
// example: v0.1.1 => 0.1.1
func TagToVersion(tag string) string {
	return strings.TrimLeft(tag, "v")
}

func StrInArr(s string, arr ...string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func Recover(fn func(), recovderFunc func(interface{})) {
	defer func() {
		if r := recover(); r != nil {
			recovderFunc(r)
		}
	}()

	fn()
}

//去掉url路径结尾的'/'
func UrlPathTrimSuffix(address string) string {
	return strings.TrimSuffix(address, "/")
}

// LimitOffset2Page
// offset 必须为 limit 的整数倍，否则会 panic
// page 从 1 开始
func LimitOffset2Page(limit int, offset int) (page int) {
	if limit <= 0 {
		return 1
	}

	if offset%limit != 0 {
		panic(fmt.Errorf("LimitOffset2Page: offset(%d) %% limit(%d) != 0", offset, limit))
	}
	return (offset / limit) + 1
}

// PageSize2Offset page 从 1 开始
func PageSize2Offset(page int, pageSize int) (offset int) {
	if page <= 1 {
		return 0
	}
	return (page - 1) * pageSize
}

// GenQueryURL url拼接
func GenQueryURL(address string, path string, params url.Values) string {
	address = UrlPathTrimSuffix(address)
	if params != nil {
		return fmt.Sprintf("%s%s?%s", address, path, params.Encode())
	} else {
		return fmt.Sprintf("%s%s", address, path)
	}
}

func JoinURL(address string, elems ...string) string {
	return fmt.Sprintf("%s/%s",
		strings.TrimRight(address, "/"),
		strings.TrimLeft(path.Join(elems...), "/"))
}
