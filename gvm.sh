#!/bin/bash
# must to have gvm (https://github.com/moovweb/gvm)
# and install a go specific version

export GVM_ROOT=$HOME/.gvm
. $GVM_ROOT/scripts/gvm-default

gvm use go1.18.8