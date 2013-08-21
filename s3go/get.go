package s3go

import (
    "launchpad.net/goamz/s3"
    "launchpad.net/goamz/aws"
    "io/ioutil"
    "os"
)

func Get(file string, bucketName string, path string, region aws.Region) {
    auth := ConnectS3()
    s := s3.New(auth, region)
    bucket := s.Bucket(bucketName)

    data, err := bucket.Get(path)
    if err != nil {
        panic(err.Error())
    }
    perms := os.FileMode(0644)
    ioutil.WriteFile(file, data, perms)
}
