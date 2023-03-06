#!/usr/bin/env bash

cd ../
docker build -t bytesfarm/chf  -f ./GoDockerfile .
docker run -v $(pwd)/bin:/data bytesfarm/chf
#scp ./bin/app vpn_vps:'~'
#scp ./bin/app vpn_vps_1:'~'

#ssh vpn_vps '~/app /etc/nginx/conf/conf.d/v2ray.conf'
#ssh vpn_vps_1 '~/app /etc/nginx/conf/conf.d/v2ray.conf'