@echo off
echo Waiting for Oracle DB to be ready...

:retry
docker exec oracle-crud-xe sqlplus -L sys/Password1234@XE as sysdba @oracle-setup.sql
IF %ERRORLEVEL% NEQ 0 (
    echo Oracle is not ready yet... sleeping 5s
    timeout /t 5
    goto retry
)

echo Oracle setup completed!
pause
