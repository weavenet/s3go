package s3go

import "strings"

type S3Url struct {
    Url string
}

func (r *S3Url) Bucket() string {
    return r.keys()[0]
}

func (r *S3Url) Key() string {
    return strings.Join(r.keys()[1: len(r.keys())], "/")
}

func (r *S3Url) Valid() bool {
    return strings.HasPrefix(r.Url, "s3://")
}

func (r *S3Url) keys() []string {
    trimmed_string := strings.TrimLeft(r.Url, "s3://")
    return strings.Split(trimmed_string, "/")
}
