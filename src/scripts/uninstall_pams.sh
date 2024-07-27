#!/bin/bash
set -e

WRKDIR=$(pwd)
CMDDIR=${WRKDIR}/src/cmd

# Check if the script is run as root
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

# Install PAMS
echo "===================="
echo "Uninstalling PAMS"
echo "===================="

if $(systemctl is-active --quiet pams.timer); then
    echo "===================="
    echo "Stopping and disabling pams.timer"
    echo "===================="
    echo "systemctl stop pams.timer"
    systemctl stop pams.timer

    echo "systemctl disable pams.timer"
    systemctl disable pams.timer
fi

if [ -f "/etc/systemd/system/pams.timer" ]; then
    echo "===================="
    echo "Removing pams.timer"
    echo "===================="
    echo "rm -rf /etc/systemd/system/pams.timer"
    rm -rf /etc/systemd/system/pams.timer
fi

if [ -f "/etc/systemd/system/pams.service" ]; then
    echo "===================="
    echo "Removing pams.service"
    echo "===================="

    echo "rm -rf /etc/systemd/system/pams.service"
    rm -rf /etc/systemd/system/pams.service
fi

if [ -f "/usr/local/bin/pams" ]; then
    echo "===================="
    echo "Removing PAMS binary"
    echo "===================="

    echo "rm -rf /usr/local/bin/pams"
    rm -rf /usr/local/bin/pams
fi