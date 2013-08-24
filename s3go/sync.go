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
    return false
}

func (s *SyncPair) validPair() bool {
     if exists(s.Source) == false && exists(s.Target) == false {
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

func exists(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return false
}
