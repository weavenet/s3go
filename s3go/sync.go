package s3go

import (
    "crypto/md5"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "launchpad.net/goamz/aws"
    "launchpad.net/goamz/s3"
)

type SyncPair struct {
    Source string
    Target string
    Auth aws.Auth
}

func (s *SyncPair) Sync() bool {
    if s.validPair() {
        if validS3Url(s.Source) {
           s.syncS3ToDir()
           return true
        } else {
           s.syncDirToS3()
           return true
        }
    }
    fmt.Printf("Path not valid.")
    return false
}

func (s *SyncPair) syncDirToS3() bool {
    sourceFiles := loadLocalFiles(s.Source)
    targetFiles := loadS3Files(s.Target, s.Auth)
    fmt.Printf("Sources:\n")
    for k, _ := range sourceFiles { fmt.Printf("Key %s Value %s\n", k, sourceFiles[k]) }
    fmt.Printf("Targets:\n")
    for k, _ := range targetFiles { fmt.Printf("Key %s Value %s\n", k, targetFiles[k]) }
    return true
}

func (s *SyncPair) syncS3ToDir() bool {
    sourceFiles := loadS3Files(s.Source, s.Auth)
    targetFiles := loadLocalFiles(s.Target)
    fmt.Printf("Sources:\n")
    for k, _ := range sourceFiles { fmt.Printf("Key %s Value %s\n", k, sourceFiles[k]) }
    fmt.Printf("Targets:\n")
    for k, _ := range targetFiles { fmt.Printf("Key %s Value %s\n", k, targetFiles[k]) }
    return true
}

func loadS3Files(url string, auth aws.Auth) map[string]string {
    files := map[string]string{}
          s3url := S3Url{Url: url}
          key := s3url.Key()
          region := aws.USEast
          s := s3.New(auth, region)
          bucket := s.Bucket(s3url.Bucket())
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          data, err := bucket.List(key, "", "", 0)
          if err != nil {
             panic(err.Error())
          }
          for key := range data.Contents {
            md5sum := data.Contents[key].ETag
            files[data.Contents[key].Key] = strings.Trim(md5sum, "\"")
          }
          return files
}

func loadLocalFiles(path string) map[string]string {
    files := map[string]string{}
    filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            relativePath := strings.TrimLeft(filePath, path)

            buf, err := ioutil.ReadFile(filePath)
            if err != nil {
                panic(err)
            }

            hasher := md5.New()
            hasher.Write(buf)
            md5sum := fmt.Sprintf("%x", hasher.Sum(nil))
            files[relativePath] = md5sum
        }
        return nil
    })
    return files
}

func (s *SyncPair) validPair() bool {
     if pathExists(s.Source) == false && pathExists(s.Target) == false {
         return false
     }
     if validS3Url(s.Source) == false && validS3Url(s.Target) == false {
         return false
     }
     return true
}

func validS3Url(path string) bool {
    return strings.HasPrefix(path, "s3://")
}

func pathExists(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return false
}
