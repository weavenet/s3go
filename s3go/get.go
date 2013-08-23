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
