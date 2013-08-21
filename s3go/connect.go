package s3go

import (
    "launchpad.net/goamz/aws"
)

func ConnectS3() (auth aws.Auth)  {
    // The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
    auth, err := aws.EnvAuth()
    if err != nil {
      panic(err.Error())
    }
    return auth
}
