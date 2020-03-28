#!/bin/bash
docker stop
docker rm mic_user_edge_service
docker run -p 9091:9091 --name mic_user_edge_service -v /usr/local/docker/micservice/user-edge-service/log:/go/src/user-edge-service/log -d mic-user-edge-service