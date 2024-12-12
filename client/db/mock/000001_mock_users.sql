INSERT INTO "users" (username, email, password, hetzner_token) VALUES ('Kevin', 'kevin@hetz.com', 'password', 'token');
INSERT INTO "users" (username, email, password, hetzner_token) VALUES ('Dejan', 'dejan@hetz.com', 'password', 'token2');

INSERT INTO "sessions" (user_id, token, expires_at, user_agent, ip_address) VALUES (1, '203410160105069f5b23a0ac5625145aebb412ee89e0ea7b770e96e7e67dddc8', '2034-10-16 01:05:22', 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36', '127.0.0.1');
INSERT INTO "sessions" (user_id, token, expires_at, user_agent, ip_address) VALUES (2, '2034101601052253882601472e1c4e6dab305e0f9d55afa4a8262e7b1ef5f304', '2034-10-16 01:05:22', 'Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36', '127.0.0.1');

