# SpecializedHTTPClient
自分のDASTの開発を助けるため + 簡単な検査ツールを作成する際に流用が効くので分離した。
## judgment	
時間や文字列が含まれるかなどの判定を行う処理をwrapした関数を提供しています。

現状は簡単なものではあるが、必要に応じて拡張や追加が可能になっている。

## client
judgmentの関数を利用した特殊なクライアントを提供している。

- TimeBase : 時間を測定し、検査などを助けます
- DisplayBase : 表示を確認します
- Client : 通常のクライアント

これらクライアントはBodyをcloseせずに返すので、受け取り後は `defer response.Body.Close()` を実行してください
