package s3go

import (
    "fmt"
    "launchpad.net/goamz/s3"
    "launchpad.net/goamz/aws"
    "os"
)

func ListBucketContents(bucketName string, region aws.Region) {
    auth := ConnectS3()
    s := s3.New(auth, region)
    bucket := s.Bucket(bucketName)
    data, err := bucket.List("", "", "", 0)
    if err != nil {
        panic(err.Error())
    }

    for key := range data.Contents {
        fmt.Printf("s3://%s/%s\n", bucketName, data.Contents[key].Key)
    }
}

func ConnectS3() (auth aws.Auth)  {
    // The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
    auth, err := aws.EnvAuth()
    if err != nil {
      panic(err.Error())
    }
    return auth
}
