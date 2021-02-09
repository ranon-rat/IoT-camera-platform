CREATE TABLE usercameras(
    id INTEGER PRIMARY KEY,
    ip VARCHAR(120),
    password TEXT NOT NULL,
    username TEXT NOT NULL,
    last_time_login INTEGER NOT NULL
);
CREATE TABLE userclients(
    cookie TEXT
);

