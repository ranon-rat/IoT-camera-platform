CREATE TABLE usercameras(
    id INTEGER PRIMARY KEY,
    ip TEXT NOT NULL,
    password TEXT NOT NULL,
    token TEXT ,
    username TEXT NOT NULL,
    last_time_login INTEGER NOT NULL
);
-- if we need to add a new colum 
-- ALTER TABLE <table> ADD <column> <type>;

CREATE TABLE userclients(
    id INTEGER PRIMARY KEY,
    id_camera_client INTEGER,
    -- this is for be related with the table usercameras
    cookie TEXT
);

