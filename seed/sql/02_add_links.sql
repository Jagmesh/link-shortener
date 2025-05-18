INSERT INTO
    links (user_id, url, hash)
VALUES
    (1, 'https://google.com', 'WPrQnnnNeL'),
    (2, 'https://i.pinimg.com/236x/c6/2e/47/c62e47ccce4e8e568c9c7e381032bde9.jpg', 'cDgscsefBn')
ON CONFLICT DO NOTHING;