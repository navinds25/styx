# STYX:

[![Build Status](https://img.shields.io/travis/navinds25/styx)](https://travis-ci.org/navinds25/styx)

## Build Steps

Currently building only works on linux. Steps are as follows:

**To generate protobuf files:**

``` make proto ```


**To build binaries:**

``` make build ```


**Sample configs are present in the configs folder**

[configs folder](configs/)


**To create certificates for testing:**

``` make certs ```

this requires cfssl to be installed, https://github.com/cloudflare/cfssl


**To Test locally with docker:**

``` make local-docker ```

make build and make certs will have to be run before make local-docker

## Adding Private Dependencies
git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
