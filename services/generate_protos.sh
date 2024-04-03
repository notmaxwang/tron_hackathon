#!/bin/bash

# Copy and generate protos for go backend
rm -r keyfi-backend/protos
cp -r keyfi_protos keyfi-backend/protos
cd keyfi-backend
./generate_protos.sh

# Copy and generate protos for react frontend
cd ../
rm -r  keyfi-frontend/protos
cp -r  keyfi_protos keyfi-frontend/protos
cd keyfi-frontend
./generate_protos.sh

