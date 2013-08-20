package s3go

import "testing"

func TestS3Url(t *testing.T) {
    url := S3Url{}
    url.SetUrl("s3://bucket1/dir/key")

    if url.Key() != "dir/key" {
      t.Error("Key not returned correctly.")
    }
    if url.Bucket() != "bucket1" {
      t.Error("Bucket not returned correctly.")
    }
}
