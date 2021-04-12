create table Topic (
    id  INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    topic TEXT,
    positive INT,
    negative INT
);

create table News (
    id    INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    body  TEXT NOT NULL,
    hash    VARCHAR(50) NOT NULL,
    create_time  TIME,
    update_time  TIME
);

create table News_Topic (
    nid   INT NOT NULL,
    tid   INT NOT NULL,
    PRIMARY KEY(nid, tid),
    FOREIGN KEY (nid) REFERENCES News (id) ON DELETE CASCADE,
    FOREIGN KEY (tid) REFERENCES Topic (id) ON DELETE CASCADE
);

create table Comment (
    id  INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    nid INT NOT NULL,
    body    TEXT NOT NULL,
    pid INT,
    isPos   TINYINT,
    create_time  TIME,
    update_time  TIME,
    FOREIGN KEY (pid) REFERENCES Comment (id) ON DELETE CASCADE
);