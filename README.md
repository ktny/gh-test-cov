# central-octocov サンプル

このリポジトリは「API テスト」と「Web テスト」を別々の GitHub Actions ワークフロー（`.github/workflows/test-api.yml`, `.github/workflows/test-web.yml`）として用意し、両方のワークフローが終わったあとに octocov を 1 回だけ実行して PR へコメントする central モードサンプルです。

## 仕組み

1. `test-api.yml` と `test-web.yml` はそれぞれ独立したワークフローです。どちらも push / PR をトリガーに動作し、`go test` でカバレッジファイル (`coverage/<target>/coverage.out`) を生成した後、`actions/upload-artifact` へ保存します。
2. `.github/workflows/coverage.yml` は push / PR をトリガーに動き、`actions/github-script` で「同じ head SHA を持つ `API Tests` と `Web Tests` がそろって成功するまで待つ」処理を行います。2 つとも成功した時点で `actions/download-artifact@v4` の `run-id` 指定でアーティファクトを取得し、octocov を central モード (`.octocov.central.yml`) → 通常モード (`.octocov.yml`) の順に 1 度だけ実行します。
3. `.octocov.central.yml` は central モード専用の設定で、集約レポートとバッジを書き出す先を `local://central-reports` / `local://central-badges` に向けています。`.octocov.yml` は PR コメント用設定で、central モードと同じデータストアを使って差分を計算します。

## 使い方

- `main` ブランチへの push と任意の PR でワークフローが起動します。
- `test-api.yml` と `test-web.yml` の両方が成功したコミットだけ `coverage.yml` が集計処理を行うため、octocov は 1 度だけ実行されます (片方のワークフローが未完了の間は `coverage.yml` が待機し続けます)。
- 集約されたレポート／バッジは `local://central-*` へ出力されるので、必要に応じてアーティファクト化したり、別リポジトリへ push することで「中央リポジトリ」を構築できます。

## 応用

- ワークフローごとに実行環境やスケジュールを変えたい場合も、この構成をベースにすれば「テストは複数ワークフロー、集計は1ワークフロー」を保てます。
- `local://central-reports` を `github://owner/central-repo/reports` に置き換えると、専用の集計リポジトリへ push する official サンプルと同じ構成になります。
