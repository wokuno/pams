#!/bin/bash
set -e

WRKDIR=$(pwd)
CMDDIR=${WRKDIR}/src/cmd

# Check if the script is run as root
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

if [ ! -f "${WRKDIR}/bin/pams" ]; then
    echo "===================="
    echo "PAMS is not Built"
    echo "./build_pams.sh"
    echo "===================="
    exit
fi

# Install PAMS
echo "===================="
echo "Installing PAMS"
echo "===================="

echo "mv ${WRKDIR}/pams ${WRKDIR}/bin/"
mv ${WRKDIR}/bin/pams /usr/local/bin/

echo "rm -rf ${WRKDIR}/bin"
rm -rf ${WRKDIR}/bin

# Create systemd service
#if file doesnt exist
if [ ! -f "/etc/systemd/system/pams.service" ]; then

echo "===================="
echo "Setting up systemd service"
echo "===================="

echo "touch /etc/systemd/system/pams.service"
touch /etc/systemd/system/pams.service

echo '[Unit]
Description="pacos Autonomous Management System"

[Service]
ExecStart=/usr/local/bin/pams' > /etc/systemd/system/pams.service

echo "cat /etc/systemd/system/pams.service"
cat /etc/systemd/system/pams.service

fi

if [ ! -f "/etc/systemd/system/pams.timer" ]; then

echo "===================="
echo "Setting up systemd timer"
echo "===================="

echo "touch /etc/systemd/system/pams.timer"
touch /etc/systemd/system/pams.timer

# Every 10 seconds is chosen, because it takes 1 second to get a good reading,
# Not much should change in 9 seconds.

echo '[Unit]
Description="Run pams every 10 seconds"

[Timer]
OnUnitActiveSec=10
Unit=pams.service

[Install]
WantedBy=multi-user.target' > /etc/systemd/system/pams.timer

echo "cat /etc/systemd/system/pams.timer"
cat /etc/systemd/system/pams.timer

echo "===================="
echo "Starting and enabling systemd timer"
echo "===================="

echo "systemctl start pams.timer"
systemctl start pams.timer

echo "systemctl enable pams.timer"
systemctl enable pams.timer

fi

