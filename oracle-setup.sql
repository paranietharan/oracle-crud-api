-- Switch to PDB
ALTER SESSION SET CONTAINER = XEPDB1;

-- Create user
CREATE USER oracle_crud_api IDENTIFIED BY Password1234
  DEFAULT TABLESPACE users
  TEMPORARY TABLESPACE temp
  QUOTA UNLIMITED ON users;

-- Grant privileges
GRANT CREATE SESSION, CREATE TABLE, CREATE SEQUENCE, CREATE VIEW,
CREATE PROCEDURE TO oracle_crud_api;

GRANT CONNECT, RESOURCE TO oracle_crud_api;

-- Create persons table
CREATE TABLE persons (
                         id    VARCHAR2(50) PRIMARY KEY,
                         name  VARCHAR2(100) NOT NULL,
                         phone VARCHAR2(20),
                         email VARCHAR2(100) UNIQUE
);
