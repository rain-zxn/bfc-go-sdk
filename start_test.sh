#!/bin/bash
echo "start testing"


cd ./account
go test
echo "account test done"


cd ../bfc_types
go test
echo "bfc_type test done"


cd ../types
go test
echo "type test done"


cd ../lib
go test
echo "lib test done"

cd ../client
echo "client test start"
go test
echo "client test done"






read -p "Press any key to continue..." -n1
