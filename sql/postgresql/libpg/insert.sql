INSERT INTO userinfo(username,departname,created) VALUES($1, $2, $3) RETURNING uid;