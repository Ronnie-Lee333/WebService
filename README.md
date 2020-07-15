# WebService
Web Service Server Sample Code For Go

◆当サンプルは、以下の機能を提供  
1)DBの「テーブルposts」に登録、更新、照会及び削除を行うWebAPIのサンプルを提供する  
2)二つのgoroutineでそれぞれが「テーブルposts」から取得した情報と付加情報（固定情報）をつなげるJSON文字を  
クライアントへ返却するサンプルを提供する  
※ここでのJSON文字は、厳密なJSON文字列ではなく、ただ二つのJSON文字列を結合するだけ。
***
◆コードファイルは、以下  
・server.go                      =>WebServiceサーバ  
・data.go                        =>DBとやり取り  
***
◆コード詳細  
・server.go  
  1)func main                   =>サーバの入口  
  2)func handleRequest          =>DBの「テーブルposts」を操作するリクエストのハンドル  
  3)func handleJSON             =>JSONを返却するリクエストのハンドル※GETのみ対応  
  4)func getJSONWithID          =>二つのgoroutineでそれぞれが取得したJSONデータを結合し返却  
  5)handleGet                   =>DBの「テーブルposts」の検索処理  
  6)handlePost                  =>DBの「テーブルposts」の登録処理  
  7)handlePut                   =>DBの「テーブルposts」の更新処理  
  8)handleDelete                =>DBの「テーブルposts」の削除処理  
  
・data.go  
  1)func init                   =>DB接続  
  2)func retrieve               =>DBの「テーブルposts」のレコード取得  
  3)func create                 =>DBの「テーブルposts」のレコード登録  
  4)func update                 =>DBの「テーブルposts」のレコード更新  
  5)func delete                 =>DBの「テーブルposts」のレコード削除  
***
◆実行方法  
**サーバ側**  
1)ビルドは、以下実行  
  go build  
  成功したら、web_service.exeが生成される  
2)サーバ起動は、以下実行  
  ./web_service  

**クライアント側** ※windows terminalで実行  
1)「テーブルposts」に登録、更新、照会及び削除を行うWebAPIについてのリクエストは、以下はサンプル  
   ・登録  
     `Invoke-WebRequest -Method POST -Body '{"content":"Hello","author":"system"}' http://127.0.0.1:8080/post/`      
   ・検索   ※最後の1は、「テーブルposts」のID  
     `Invoke-WebRequest -Method GET http://127.0.0.1:8080/post/1`  
   ・更新   ※最後の1は、「テーブルposts」のID  
     `Invoke-WebRequest -Method PUT -Body '{"content":"Post","author":"s.lee"}' http://127.0.0.1:8080/post/1`  
   ・削除   ※最後の1は、「テーブルposts」のID  
     `Invoke-WebRequest -Method DELETE http://127.0.0.1:8080/post/1`  
 
 2)並行処理でJSON文字を結合するWebAPIについてのリクエストは、以下はサンプル  
   `Invoke-WebRequest -Method GET http://127.0.0.1:8080/japi/1`  
***
◆その他  
postgreSQLのDB作成、テーブル生成の資源は、以下  
・db_install.sql  
・table.sql  
 
   
