package s3go

import (
    "os"
    "strings"
)

type SyncPair struct {
    Source string
    Target string
}

func (s *SyncPair) Sync() bool {
    if s.validPair() {
        return true
    }
    if validS3Url(s.Source) {
       s.syncS3ToDir()
    } else {
       s.syncDirToS3()
    }
    return false
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
