package s3go

import (
    "launchpad.net/goamz/s3"
    "launchpad.net/goamz/aws"
    "io/ioutil"
)

func Put(bucketName string, path string, file string, region aws.Region) {
    auth := ConnectS3()
    contType := "binary/octet-stream"
    Perms := s3.ACL("private")
    s := s3.New(auth, region)
    bucket := s.Bucket(bucketName)

    data, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err.Error())
    }

    err = bucket.Put(path, data, contType, Perms)
    if err != nil {
        panic(err.Error())
    }
}
