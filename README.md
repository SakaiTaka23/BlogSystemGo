# BlogSystemGo
goで簡単なブログシステムを作成した

# データベース
dbname: gotest
```
 CREATE TABLE entries (
    id INT AUTO_INCREMENT,
    title TEXT,
    content TEXT,
    created DATETIME,
    primary key (id)
);
```
# 参考
https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/13.5.html