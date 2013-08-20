package s3go

import "strings"

type S3Url struct {
  url string
}

func (r *S3Url) bucket() string {
    return r.keys()[0]
}

func (r *S3Url) key() string {
    return strings.Join(r.keys()[1: len(r.keys())], "/")
}

func (r *S3Url) keys() []string {
    trimmed_string := strings.Trim(r.url, "s3://")
    return strings.Split(trimmed_string, "/")
}
