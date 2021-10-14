#!/bin/bash

user=$(whoami)

if [ user != root ]
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

echo "Input virt-go network address : "
echo -n "ex) 10.62.62.0 : "
read netAddr
echo ""

mkdir -p /etc/virt-go
echo << EOF >> /etc/virt-go/virt-go.cfg
dataDir=$dataDir
netAddr=$netAddr
EOF

mkdir -p $dataDir/volumes
mkdir -p $dataDir/images
mkdir -p $dataDir/cloud-init


