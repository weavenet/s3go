package main

import (
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/brettweavnet/s3go/s3go"
    "launchpad.net/goamz/aws"
    "os"
)

func main() {
    app := cli.NewApp()
    app.Name = "s3go"
    app.Usage = "CLI for S3"

    region := aws.USEast

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
          s3url := s3go.S3Url{}
          s3url.SetUrl(c.Args()[0])
          bucket := s3url.Bucket()
          fmt.Printf("Listing contents of bucket '%s' in region '%s'.\n", bucket, region.Name)
          s3go.ListBucketContents(bucket, region)
        },
      },
      {
        Name:        "put",
        Usage:       "s3go put LOCAL_FILE s3://BUCKET/KEY",
        Description: "Put file in s3.",
        Action: func(c *cli.Context) {
          if len(c.Args()) < 2 {
             fmt.Printf("Local file and S3 location required.")
             os.Exit(1)
          }
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          local_file := c.Args()[0]
          s3url := s3go.S3Url{}
          s3url.SetUrl(c.Args()[1])
          bucket := s3url.Bucket()
          key := s3url.Key()
          fmt.Printf("Putting file '%s' in 's3://%s/%s'.\n", local_file, bucket, key)
          s3go.Put(bucket, key, local_file, region)
        },
      },
      {
        Name:        "rm",
        Usage:       "s3go rm s3://BUCKET/KEY",
        Description: "Remove key from s3.",
        Action: func(c *cli.Context) {
          if len(c.Args()) < 1 {
             fmt.Printf("S3 location required.")
             os.Exit(1)
          }
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          s3url := s3go.S3Url{}
          s3url.SetUrl(c.Args()[0])
          bucket := s3url.Bucket()
          key := s3url.Key()
          fmt.Printf("Removing file 's3://%s/%s'.\n", bucket, key)
          s3go.Del(bucket, key, region)
        },
      },
    }
    app.Run(os.Args)
}
