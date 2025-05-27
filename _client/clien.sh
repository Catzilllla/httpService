#!/bin/bash

agrs=$1
port=$2

if [ "$1" -eq 1 ]; then
  request="POST"
  path='execute'

  curl -X $request \
  -H "Content-Type: application/json" \
  -d '{
    "object_cost": 5000000,
    "initial_payment": 1000000,
    "months": 240,
    "program": {
        "salary": true
    }
  }' \
  http://38.107.235.2:$port/$path
fi

if [ "$1" -eq 3 ]; then
  request="POST"
  path='execute'

  curl -X $request \
  -H "Content-Type: application/json" \
  -d '{
    "object_cost": 3000000,
    "initial_payment": 3000000,
    "months": 1240,
    "program": {
        "salary": true
    }
  }' \
  http://38.107.235.2:$port/$path
fi

if [ "$1" -eq 2 ]; then
  request='GET'
  path='cache'

  curl http://38.107.235.2:$port/$path
fi

echo ".................................."
echo $request
echo $1

