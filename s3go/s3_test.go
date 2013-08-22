package s3go

import "testing"

type testcase struct {
    url string
    bucket string
    key string
}

var tests = []testcase {
    { "s3://bucket/test.tar.gz", "bucket", "test.tar.gz" },
    { "s3://bucket-123/dir/folder/key", "bucket-123", "dir/folder/key" },
    { "s3://bucket-123/files*", "bucket-123", "files*" },
    { "bucket-123/dir/folder/key", "bucket-123", "dir/folder/key" },
    { "bucket-123", "bucket-123", "" },
}

func TestS3Url(t *testing.T) {
    for _, c := range tests {
        url := S3Url{Url: c.url}

        if url.Key() != c.key {
          t.Error("Key not returned correctly.")
        }

        if url.Bucket() != c.bucket {
          t.Error("Bucket not returned correctly.")
        }
    }
}
