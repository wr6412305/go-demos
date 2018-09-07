package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/simplejia/utils"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/simplejia/clog"
	"github.com/simplejia/namecli/api"
)

const (
	OSSGeneralDir   = "xnggeneral"
	OSSEndpointProd = "oss-cn-shenzhen-internal.aliyuncs.com"
	OSSEndpointTest = "oss-cn-shenzhen.aliyuncs.com"
	OSSBucket       = "xngstatic"
	OSSAccessKey    = "LTAIvRE6ZPgt5Msy"
	OSSSecretKey    = "31sQI8GxbU0ujm9KsVhF0V7aBQffi0"
)

var (
	OSSEndpoint string
)

func UploadFileToOSS(filePath, endpoint, accessKey, secretKey, bucket string) (err error) {
	client, err := oss.New(endpoint, accessKey, secretKey)
	if err != nil {
		return
	}

	bk, err := client.Bucket(bucket)
	if err != nil {
		return
	}

	objectKey := fmt.Sprintf("%s/%s", OSSGeneralDir, filePath)
	err = bk.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		return
	}

	return
}

func getFileList(dirPath string) (err error) {
	fun := "getFileList"

	err = filepath.Walk(dirPath, func(filePath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		re := regexp.MustCompile(`^log\d+\.\d+\.\d+\.\d+\+main_\d+$`)
		if !re.MatchString(f.Name()) {
			return nil
		}

		tarFilePath := filePath + "_" + utils.LocalIp + ".tar.gz"
		if err := exec.Command("tar", "-zcvf", tarFilePath, filePath).Run(); err != nil {
			return err
		}
		defer func() {
			if err := exec.Command("rm", tarFilePath).Run(); err != nil {
				clog.Error("%s rm err: %v, req: %v", fun, err, tarFilePath)
			}
		}()

		if err := UploadFileToOSS(tarFilePath, OSSEndpoint, OSSAccessKey, OSSSecretKey, OSSBucket); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	var dirname string
	flag.StringVar(&dirname, "dir", "", "log directory name")
	var env string
	flag.StringVar(&env, "env", "prod", "specify env")
	flag.Parse()

	clog.AddrFunc = func() (string, error) {
		return api.Name("clog.srv.ns")
	}
	clog.Init("tools", "", 14, 3)

	clog.Info("upload_files_to_oss")

	if env == "prod" {
		OSSEndpoint = OSSEndpointProd
	} else {
		OSSEndpoint = OSSEndpointTest
	}

	if err := getFileList(dirname); err != nil {
		clog.Error("getFileList err: %v, req: %v", err, dirname)
		return
	}
}
