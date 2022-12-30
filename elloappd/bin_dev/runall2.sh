#!/usr/bin/env bash

echo "run idgen ..."
nohup ./idgen -f=../etc_dev/idgen.yaml >> ../logs/idgen.log  2>&1 & sleep 1

echo "run status ..."
nohup ./status -f=../etc_dev/status.yaml >> ../logs/status.log  2>&1 & sleep 1

echo "run authsession ..."
nohup ./authsession -f=../etc_dev/authsession.yaml >> ../logs/authsession.log  2>&1 & sleep 1

echo "run dfs ..."
nohup ./dfs -f=../etc_dev/dfs.yaml >> ../logs/dfs.log  2>&1 & sleep 1

echo "run media ..."
nohup ./media -f=../etc_dev/media.yaml >> ../logs/media.log  2>&1 & sleep 1

echo "run biz ..."
nohup ./biz -f=../etc_dev/biz.yaml >> ../logs/biz.log  2>&1 & sleep 1

echo "run msg ..."
nohup ./msg -f=../etc_dev/msg.yaml >> ../logs/msg.log  2>&1 & sleep 1

echo "run sync ..."
nohup ./sync -f=../etc_dev/sync.yaml >> ../logs/sync.log  2>&1 & sleep 1

echo "run bff ..."
nohup ./bff -f=../etc_dev/bff.yaml >> ../logs/bff.log  2>&1 & sleep 5

echo "run session ..."
nohup ./session -f=../etc_dev/session.yaml >> ../logs/session.log  2>&1 & sleep 1

echo "run gateway ..."
nohup ./gateway -f=../etc_dev/gateway.yaml >> ../logs/gateway.log  2>&1 & sleep 1
