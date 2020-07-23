#!/bin/sh

USER_ID="$1"

if [ -z "$USER_ID" ]
  then
    echo "No USER_ID argument"
    exit 1
fi

psql -v USER_ID="$USER_ID" -U baristaschool <<-EOSQL
    SELECT
      user_id,
      date(timestamp) as date,
      count(1) as answers_count,
      round(100. * sum(correctly_answered::int) / count(1), 1)::text || '%' as correct_answers_rate
    FROM usermemorizedquestion
    WHERE user_id = :USER_ID
    GROUP BY user_id, date
    ORDER BY date;
EOSQL
