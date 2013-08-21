package s3go

import (
    "launchpad.net/goamz/s3"
    "launchpad.net/goamz/aws"
)

func Del(bucketName string, path string, region aws.Region) {
    auth := ConnectS3()
    s := s3.New(auth, region)
    bucket := s.Bucket(bucketName)

    err := bucket.Del(path)
    if err != nil {
        panic(err.Error())
    }
}
