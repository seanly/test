#!/bin/bash

set -eux

npm config set registry=https://registry.npmmirror.com/
npm config set sass_binary_site=https://npmmirror.com/mirrors/node-sass
npm config set strict-ssl false

npm install node-sass@8.0.0 --unsafe-perm --save-dev --loglevel warn