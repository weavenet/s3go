package s3go

import "testing"
import "fmt"

func TestS3Url(t *testing.T) {
    url := S3Url{url: "s3://bucket1/dir/key"}

    fmt.Printf("test %s", url.key())

    if url.key() != "dir/key" {
      t.Error("Key not returned correctly.")
    }
    if url.bucket() != "bucket1" {
      t.Error("Bucket not returned correctly.")
    }
}
