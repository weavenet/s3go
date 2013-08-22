package s3go

import "testing"

type testpair struct {
    url string
    bucket string
    key string
}

var tests = []testpair {
    { "s3://bucket/test.tar.gz", "bucket", "test.tar.gz" },
    { "s3://bucket-123/dir/folder/key", "bucket-123", "dir/folder/key" },
    { "bucket-123/dir/folder/key", "bucket-123", "dir/folder/key" },
}

func TestS3Url(t *testing.T) {
    for _, pair := range tests {
        url := S3Url{Url: pair.url}

        if url.Key() != pair.key {
          t.Error("Key not returned correctly.")
        }

        if url.Bucket() != pair.bucket {
          t.Error("Bucket not returned correctly.")
        }
    }
}
