package s3go

import (
    "launchpad.net/goamz/s3"
    "launchpad.net/goamz/aws"
)

func ListBucketContents(bucketName string, key string, region aws.Region) (contents []s3.Key) {
    auth := ConnectS3()
    s := s3.New(auth, region)
    bucket := s.Bucket(bucketName)
    data, err := bucket.List(key, "", "", 0)
    if err != nil {
        panic(err.Error())
    }

    return data.Contents
}
