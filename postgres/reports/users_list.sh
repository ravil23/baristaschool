#!/bin/sh

psql -U baristaschool <<-EOSQL
    SELECT *
    FROM public.user
    ORDER BY created_at DESC;
EOSQL
