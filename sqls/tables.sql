create table Topic (
    id  INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    topic TEXT,
    Positive INT,
    Negative INT
);

create table News (
    id    INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    body  TEXT NOT NULL
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
    FOREIGN KEY (pid) REFERENCES Comment (id) ON DELETE CASCADE
);