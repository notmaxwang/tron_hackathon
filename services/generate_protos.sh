#!/bin/bash

# Copy and generate protos for go backend
rm keyfi-backend/protos/*
cp protos_main/* keyfi-backend/protos
cd keyfi-backend
./generate_protos.sh

# Copy and generate protos for react frontend
cd ../
rm keyfi-frontend/protos/*
cp protos_main/* keyfi-frontend/protos
cd keyfi-frontend
./generate_protos.sh

