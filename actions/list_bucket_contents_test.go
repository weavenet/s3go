package s3go

import "testing"

func TestConnectS3(t *testing.T) {
    auth := ConnectS3()
    if 1 != 2 {
        t.Error("Expected 1 got 2")
    }
}
