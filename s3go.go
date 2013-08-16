package s3go

import (
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/brettweavnet/s3go/actions"
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
          bucket := c.Args()[0]
          fmt.Printf("Listing contents of bucket '%s' in region '%s'.\n", bucket, region.Name)
          ListBucketContents(bucket, region)
        },
      },
    }

    app.Run(os.Args)
}
