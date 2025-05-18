INSERT INTO
    stats (link_id, click, day_date)
VALUES
    (1, 5, CURRENT_DATE),
    (2, 5, CURRENT_DATE)
ON CONFLICT DO NOTHING;