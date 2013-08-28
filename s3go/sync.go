package s3go

import (
    //"crypto/sha1"
    "io"
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
           s.syncS3ToDir()
           return true
        } else {
           s.syncDirToS3()
           return true
        }
    }
    println("Path not valid.")
    return false
}

func (s *SyncPair) syncDirToS3() bool {
    loadFiles(s.Source)
    return true
}

func (s *SyncPair) syncS3ToDir() bool {
    loadFiles(s.Target)
    return true
}

func loadFile(path string, f os.FileInfo, e error) error {
  return nil
}

func loadFiles(path string) map[string]string {
    files := map[string]string{}
    filepath.Walk(path, func(file_path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            // Add path to files
            fi, err := os.Open(path)
            if err != nil {
                panic(err)
            }
            buf := make([]byte, 1024)
            for {
                // read a chunk
                n, err := fi.Read(buf)
                if err != nil && err != io.EOF { panic(err) }
                if n == 0 { break }
            }

            //files[file_path]
            // read the file into content here
            println(file_path)
            println(buf)
            println(fi)
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
