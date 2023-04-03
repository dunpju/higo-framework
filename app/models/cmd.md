go run bin\main.go -gen=model -out=app\models -name=table_name
table_name如果是all,就构建所有表;
构建model时会自动构建Dao层、Entity层