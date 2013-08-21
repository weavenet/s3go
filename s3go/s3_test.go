package s3go

import "testing"

func TestS3Url(t *testing.T) {
    url := S3Url{}

    url.SetUrl("s3://bucket1/dir/test123")

    if url.Key() != "dir/test123" {
      t.Error("Key not returned correctly.")
    }

    if url.Bucket() != "bucket1" {
      t.Error("Bucket not returned correctly.")
    }
}
