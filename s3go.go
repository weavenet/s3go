package main

import (
    "os"
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/brettweavnet/s3go/s3go"
    "launchpad.net/goamz/aws"
    "launchpad.net/goamz/s3"
    "strings"
)

func main() {
    app := cli.NewApp()
    app.Name = "s3go"
    app.Usage = "CLI for S3"

    region := aws.USEast
    auth := s3go.ConnectS3()
    s := s3.New(auth, region)

    app.Commands = []cli.Command{
      {
        Name:        "ls",
        Usage:       "s3go ls s3:://BUCKET",
        Description: "List contents of bucket.",
        Action: func(c *cli.Context) {
          if len(c.Args()) == 0 {
             fmt.Printf("Bucket required.")
             os.Exit(1)
          }
          s3url := s3go.S3Url{Url: c.Args()[0]}
          key := s3url.Key()
          bucket := s.Bucket(s3url.Bucket())
          fmt.Printf("Listing contents of bucket '%s' in region '%s'.\n", bucket.Name, region.Name)
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          data, err := bucket.List(key, "", "", 0)
          if err != nil {
             panic(err.Error())
          }

          for key := range data.Contents {
              fmt.Printf("s3://%s/%s\n", bucket.Name, data.Contents[key].Key)
          }
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
          local_file := c.Args()[0]
          s3url := s3go.S3Url{Url: c.Args()[1]}
          bucket := s.Bucket(s3url.Bucket())
          key := s3url.Key()
          if key == "" {
              key = strings.Split(local_file, "/")[len(strings.Split(local_file, "/"))-1]
          }
          fmt.Printf("Putting file '%s' in 's3://%s/%s'.\n", local_file, bucket.Name, key)
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          s3go.Put(bucket, key, local_file)
        },
      },
      {
        Name:        "get",
        Usage:       "s3go get s3://BUCKET/KEY LOCAL_FILE",
        Description: "Get file in s3.",
        Action: func(c *cli.Context) {
          if len(c.Args()) < 2 {
             fmt.Printf("S3 location required and local file.")
             os.Exit(1)
          }
          s3url := s3go.S3Url{Url: c.Args()[0]}
          bucket := s.Bucket(s3url.Bucket())
          key := s3url.Key()
          local_file := c.Args()[1]
          fmt.Printf("Downloading file 's3://%s/%s' into '%s'.\n", bucket.Name, key, local_file)
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          s3go.Get(local_file, bucket, key)
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
          s3url := s3go.S3Url{Url: c.Args()[0]}
          bucket := s.Bucket(s3url.Bucket())
          key := s3url.Key()
          fmt.Printf("Removing file 's3://%s/%s'.\n", bucket.Name, key)
          defer func() {
              if r := recover(); r != nil {
                  fmt.Printf("%v", r)
              }
          }()
          bucket.Del(key)
        },
      },
      {
        Name:        "sync",
        Usage:       "s3go sync LOCAL_DIR s3://BUCKET/KEY",
        Description: "Sync local dir with S3 URL.",
        Action: func(c *cli.Context) {
          if len(c.Args()) < 2 {
             fmt.Printf("S3 URL and local directory required.")
             os.Exit(1)
          }
          arg0 := c.Args()[0]
          arg1 := c.Args()[1]
          fmt.Printf("Syncing %s with %s\n", arg0, arg1)
          sync := s3go.SyncPair{arg0, arg1}
          result := sync.Sync()
          fmt.Printf("Result: %s\n", result)
        },
      },
    }
    app.Run(os.Args)
}

func visit(path string, f os.FileInfo, err error) error {
  fmt.Printf("Visited: %s\n", path)
  return nil
}
