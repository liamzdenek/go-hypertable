go-hypertable
=============

This repository contains the autogenerated bindings in Google Go for the Hypertable database. These bindings were built for Hypertable 0.9.7.9 with Thrift 0.9.1.

Generation Instructions
-----------------------

If this repo happens to be out of date, or you require bindings for another version:

```sh
# after installing git and thrift... (hypertable doesn't need to be compiled to
# prepare the bindings, nor does this repository need to be)

# clone hypertable's source
git clone github.com/hypertable/hypertable.git

# `git checkout` whatever version of hypertable you need. not performing a
# checkout will build the bindings for the git HEAD, which may be unstable.

# generate the bindings
thrift -gen go hypertable/src/cc/ThriftBroker/Client.thrift
cd gen-go/Client/

# you should now be in a directory similar to this repository, less this
# README, and the Example directory.
```