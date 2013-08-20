package s3go

import "os"
import "testing"

func TestConnectS3(t *testing.T) {
    auth := ConnectS3()
    if auth.AccessKey != os.Getenv("AWS_ACCESS_KEY_ID") {
        t.Error("Could not read AWS_ACCESS_KEY_ID.")
    }
    if auth.SecretKey != os.Getenv("AWS_SECRET_ACCESS_KEY") {
        t.Error("Could not read AWS_SECRET_ACCESS_KEY.")
    }
}
