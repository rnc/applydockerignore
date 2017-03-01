# applydockerignore

## Introduction

This is a small utility program to apply a dockerignore on a path. `.dockerignore` files are documented https://docs.docker.com/engine/reference/builder/. As they use Go's filepath.Match rules it is not possible to apply simple bash filter to the dockerignore file.

The use case is where the repository is ADD'ed via a tar file to the docker image temporary directory rather than being processed normally. This program can then be used to post-process that unpacked tar prior to it being manipulated inside the image.

## Usage

applydockerignore <path>

The `.dockerignore` file must exist within the target path.
