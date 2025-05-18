INSERT INTO
    users (email, name, password)
VALUES
    ('test@test.com', 'test', 'somepass'),
    ('second@user.com', 'secondUser', 'ultrapass')
ON CONFLICT DO NOTHING;