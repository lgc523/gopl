package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

const (
	endpoint   = ""
	keyId      = ""
	secret     = ""
	bucketName = ""
	keyPrefix  = "http://" + bucketName + "." + endpoint + "/"
)

var dir string

func main() {
	flag.StringVar(&dir, "dir", "", "print this dir files")
	flag.Parse()
	if len(dir) == 0 {
		fmt.Println("input your dir like: -dir xxx")
		os.Exit(1)
	}
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, keyId, secret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	marker := oss.Marker("")
	prefix := oss.Prefix(dir)
	for {
		lor, err := bucket.ListObjects(marker, prefix)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		for _, object := range lor.Objects {
			fmt.Printf("%s %d %s\n", object.LastModified, object.Size, keyPrefix+object.Key)
		}

		prefix = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}
}
