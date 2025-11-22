# central-octocov サンプル

このリポジトリは「API テスト」と「Web テスト」の 2 つの GitHub Actions ジョブで作成したカバレッジレポートを、octocov の central モードで集約したあとに 1 度だけ PR へコメントする最小構成の例です。

## 仕組み

1. `.github/workflows/central.yml` には 3 つのジョブがあります。
   - `api`: `internal/api` 配下のテストを実行し、`coverage/api/coverage.out` を生成してアーティファクトへ保存します。
   - `web`: `internal/web` 配下のテストを実行し、`coverage/web/coverage.out` を生成してアーティファクトへ保存します。
   - `coverage`: 上記 2 つのジョブが完了した後 (`needs`) にアーティファクトをまとめて取得し、octocov を central モードで 1 度実行して `local://central-reports` と `local://central-badges` に集約結果を出力します。続けて通常モードで octocov を実行し、集約済みカバレッジを元に PR コメントを投稿します。
2. `.octocov.central.yml` は central モード専用の設定で、集約レポートとバッジの書き出し先を `local://` に向けています。サンプルのためローカル出力ですが、`github://` や `artifact://` に変更すれば専用リポジトリへ集約する形にもできます。
3. `.octocov.yml` は PR へコメントするための通常設定です。`coverage/**/coverage.out` をまとめて入力にし、central モードと同じ `local://central-reports` をレポート／差分の保存先にしています。

## 使い方

- `main` ブランチへの push と任意の PR でワークフローが起動します。
- `api` と `web` がそれぞれ独立に成功したあと、`coverage` ジョブが動き出し、octocov によるコメントが 1 度だけ投稿されます。
- 集約されたレポートやバッジは central モードの出力ディレクトリ以下に生成されるため、必要に応じてアーティファクト化したり、別リポジトリへ push することで「中央リポジトリ」を構築できます。

## 応用

- `api` / `web` のジョブを別ワークフローに分割したい場合は、各ワークフローでアーティファクトをアップロードし、`workflow_run` トリガーで `coverage` 専用ワークフローを起動してください。その際も central モード用設定と通常設定を切り替えて実行すれば、同じ考え方で 2 つ以上のテスト結果をまとめて計測できます。
- `local://central-reports` を `github://owner/central-repo/reports` に置き換えると、専用の集計リポジトリへ push する official サンプルと同じ構成になります。
