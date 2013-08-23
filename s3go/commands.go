package s3go

import (
    "launchpad.net/goamz/s3"
    "io/ioutil"
    "os"
)

func Get(file string, bucket *s3.Bucket, path string) {
    data, err := bucket.Get(path)
    if err != nil {
        panic(err.Error())
    }
    perms := os.FileMode(0644)
    ioutil.WriteFile(file, data, perms)
}

func Put(bucket *s3.Bucket, path string, file string) {
    contType := "binary/octet-stream"
    Perms := s3.ACL("private")

    data, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err.Error())
    }

    err = bucket.Put(path, data, contType, Perms)
    if err != nil {
        panic(err.Error())
    }
}
