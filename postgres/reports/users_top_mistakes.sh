#!/bin/sh

USER_ID="$1"
LIMIT="${2:-10}"

if [ -z "$USER_ID" ]
  then
    echo "No USER_ID argument"
    exit 1
fi

psql -v USER_ID="$USER_ID" -v LIMIT="$LIMIT" -U baristaschool <<-EOSQL
    SELECT
      user_id,
      question,
      count(1) as mistakes_count
    FROM usermemorizedquestion
    WHERE user_id = :USER_ID AND NOT coalesce(correctly_answered, false)
    GROUP BY user_id, question
    ORDER BY mistakes_count DESC
    LIMIT :LIMIT;
EOSQL
