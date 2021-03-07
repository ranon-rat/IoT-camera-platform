CREATE TABLE usercameras(
    id INTEGER PRIMARY KEY,
    ip VARCHAR(64)    ,MN BVCXA<NOT NULL,
    password VARCHAR(64) NOT NULL,
    token VARCHAR(64),
    username TEXT NOT NULL,
    last_time_login INTEGER NOT NULL
);
-- if we need to add a new colum 
-- ALTER TABLE <table> ADD <column> <type>;

CREATE TABLE userclients(
    id INTEGER PRIMARY KEY,
    id_camera_client INTEGER NOT NULL,
    -- this is for be related with the table usercameras
    cookie TEXT NOT NULL
);

