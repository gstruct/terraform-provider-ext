#!/bin/bash

VERSION=0.1

rm -rf bin

GOOS=darwin  GOARCH=amd64 go build -o bin/macos/terraform-provider-content
GOOS=linux   GOARCH=amd64 go build -o bin/linux/terraform-provider-content
GOOS=windows GOARCH=amd64 go build -o bin/windows/terraform-provider-content.exe

tar czf bin/terraform-vsphere-$VERSION-macos.tar.gz --directory=bin/macos terraform-provider-content
tar czf bin/terraform-vsphere-$VERSION-linux.tar.gz --directory=bin/linux terraform-provider-content
zip     bin/terraform-vsphere-$VERSION-windows.zip -j bin/windows/terraform-provider-content.exe
