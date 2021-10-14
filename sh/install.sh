#!/bin/bash

user=$(whoami)

if [ $user != "root" ]
then
    echo "Please run 'install.sh' as root account."
    exit 1
fi

if [ -f "/etc/virt-go/virt-go.cfg" ]
then
    echo "/etc/virt-go/virt-go.cfg file already exists."
    exit 1
fi

echo "Input virt-go data directory"
echo -n "ex) /etc/virt-go : "
read dataDir
echo ""

echo "Input virt-go network address"
echo -n "ex) 10.62.62.0 : "
read netAddr
echo ""

mkdir -p /etc/virt-go
cat << EOF >> /etc/virt-go/virt-go.cfg
dataDir=$dataDir
netAddr=$netAddr
EOF

mkdir -p $dataDir/volumes
mkdir -p $dataDir/images
mkdir -p $dataDir/cloud-init

curl https://raw.githubusercontent.com/YoungjuWang/virt-go/v2/cloud-init/user-data -o $dataDir/cloud-init/user-data
curl https://raw.githubusercontent.com/YoungjuWang/virt-go/v2/cloud-init/meta-data -o $dataDir/cloud-init/meta-data

echo "Please change value below in 'user-data' file for login to virt-go server without password."
echo "ssh_authorized_keys:
	>>	  - <pub-key>"
echo ""
