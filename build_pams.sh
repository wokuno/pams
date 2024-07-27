#!/bin/bash
set -e

WRKDIR=$(pwd)
CMDDIR=${WRKDIR}/src/cmd

echo "===================="
echo "Building PAMS"
echo "===================="

echo "mkdir -p ${WRKDIR}/bin"
mkdir -p ${WRKDIR}/bin

echo "cd ${CMDDIR}"
cd ${CMDDIR}

echo "go get ."
go get .

echo "go build -o ${WRKDIR}/bin/pams"
go build -o ${WRKDIR}/bin/pams -v
