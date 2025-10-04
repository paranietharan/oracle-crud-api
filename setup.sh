#!/bin/bash

echo "Waiting for Oracle DB to be ready..."
until docker exec oracle-crud-xe sqlplus -L sys/Password1234@XE as sysdba @oracle-setup.sql
do
  echo "Oracle is not ready yet... sleeping 5s"
  sleep 5
done

echo "Oracle setup completed!"