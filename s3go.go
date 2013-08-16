package main

import (
    "fmt"
    "github.com/codegangsta/cli"
    "launchpad.net/goamz/s3"
    "launchpad.net/goamz/aws"
    "os"
)

func main() {
    app := cli.NewApp()
    app.Name = "s3go"
    app.Usage = "CLI for S3"

    app.Commands = []cli.Command{
      {
        Name:        "ls",
        Usage:       "s3go ls BUCKET_NAME",
        Description: "List contents of bucket.",
        Action: func(c *cli.Context) {
          if len(c.Args()) == 0 {
             fmt.Printf("Bucket required.")
             os.Exit(1)
          }
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          ListBucketContents(c.Args()[0])
        },
      },
    }

    app.Run(os.Args)
}

func ListBucketContents(bucketName string) {
    auth := ConnectS3()
    s := s3.New(auth, aws.USEast)
    bucket := s.Bucket(bucketName)
    data, err := bucket.List("", "", "", 0)
    if err != nil {
        panic(err.Error())
    }

    for key := range data.Contents {
        fmt.Printf("s3://%s/%s\n", bucketName, data.Contents[key].Key)
    }
}

func ConnectS3() (auth aws.Auth)  {
    // The AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables are used.
    auth, err := aws.EnvAuth()
    if err != nil {
      panic(err.Error())
    }
    return auth
}
