#!/bin/bash

INTELLIJ_PLUGINS_DIR="<Insert your value>"
TERRAFORM_DIR="<Insert your value>"
LOCAL_OS="<Insert your value>"

cp -f bin/content.json "$INTELLIJ_PLUGINS_DIR/intellij-hcl/classes/terraform/model/providers/content.json"
cp -f "bin/$LOCAL_OS/terraform-provider-content" "$TERRAFORM_DIR/terraform-provider-content"
