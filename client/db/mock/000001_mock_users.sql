INSERT INTO "users" (username, email, password, hetzner_token) VALUES ('Kevin', 'kevin@hetz.com', 'password', 'token');
INSERT INTO "users" (username, email, password, hetzner_token) VALUES ('Dejan', 'dejan@hetz.com', 'password', 'token2');

INSERT INTO "sessions" (user_id, token, expires_at) VALUES (1, '203410160105069f5b23a0ac5625145aebb412ee89e0ea7b770e96e7e67dddc8', '2034-10-16 01:05:22');
INSERT INTO "sessions" (user_id, token, expires_at) VALUES (2, '2034101601052253882601472e1c4e6dab305e0f9d55afa4a8262e7b1ef5f304', '2034-10-16 01:05:22');
