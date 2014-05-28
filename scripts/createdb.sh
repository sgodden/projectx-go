#!/bin/bash
DB=projectx

dropdb $DB
createdb $DB
psql $DB < ../src/projectx/persistence/createdb.sql
 