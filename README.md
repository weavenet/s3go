[![Build Status](https://secure.travis-ci.org/brettweavnet/s3go.png)](http://travis-ci.org/brettweavnet/s3go)

# s3go

CLI to interact with S3 written in Go.

# Installation

Clone the repo:

    get clone https://github.com/brettweavnet/s3go

Chagne in the s3go directory and run make:

    cd s3go
    make

# Setup

Set environment variables:

    AWS_SECRET_ACCESS_KEY=yyy
    AWS_ACCESS_KEY_ID=xxx

# Usage

## ls

    s3go ls s3_url

## put

    s3go put file s3_url

## get

    s3go put s3_url file

## rm

    s3go rm s3_url
