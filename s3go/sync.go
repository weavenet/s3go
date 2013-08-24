package s3go

import (
    "os"
    "path/filepath"
    "strings"
    "launchpad.net/goamz/aws"
)

type SyncPair struct {
    Source string
    Target string
    Auth aws.Auth
}

func (s *SyncPair) Sync() bool {
    if s.validPair() {
        if validS3Url(s.Source) {
           loadFiles("/etc/")
           s.syncS3ToDir()
        } else {
           s.syncDirToS3()
        }
    }
    return false
}

func loadFile(path string, f os.FileInfo, err error) error {
  println("%s", path)
  return nil
} 

func loadFiles(path string) {
  filepath.Walk(path, loadFile)
}

func (s *SyncPair) syncDirToS3() bool {
    return true
}

func (s *SyncPair) syncS3ToDir() bool {
    return true
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
