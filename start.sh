#!/usr/bin/env bash

etcd &
redis-server &

path=$(pwd)$"/server"

for file in ${path}/*; do
  name=$(basename "$file")
  rpcdir=$file$"/rpc/"
  rpcfile=$name$".go"
  yamlfile=$rpcdir$"etc/"$name$".yaml"

  cd "$rpcdir"

  go run "$rpcfile" -f "$yamlfile" &

done
