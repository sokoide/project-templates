# Modern Language Feature Demos & Cheat Sheets

このリポジトリは、C++ と Go の主要な言語機能の進化を、実際のコード例とともに分かりやすくまとめた学習リソース兼クイックリファレンスです。

## 🎯 このリポジトリの目的

- **迅速なキャッチアップ**: 各言語バージョンの主要な変更点を短時間で把握する。
- **実務への適用**: 「どのような場面でどの機能を使うべきか」というエンジニアの意思決定をサポートする。
- **動作するサンプル**: 全ての機能について、実際にコンパイル・実行可能な最小構成のコードを提供。

## 📚 目次

### [C++ セクション](./README_cpp_ja.md)

C++11 から C++23、そして次世代の C++26 まで。

- [C++11](./cpp11/README_ja.md) - 言語の大規模近代化
- [C++17](./cpp17/README_ja.md) - 実務完成度アップ
- [C++20](./cpp20/README_ja.md) - 設計思想の進化
- [C++23](./cpp23/README_ja.md) - 現実的な適用と整備

### [Go セクション](./README_go_ja.md)

Go 1.0 から最新の 1.26 まで。

- [Go 1.16](./go1.16/README_ja.md) - `//go:embed` リソース埋め込み
- [Go 1.18](./go1.18/README_ja.md) - Generics（史上最大の言語拡張）
- [Go 1.21](./go1.21/README_ja.md) - `slog`、`slices`、`min`/`max`/`clear`
- [Go 1.22](./go1.22/README_ja.md) - ループ変数スコープ修正、HTTP ルーティング
- [Go 1.23](./go1.23/README_ja.md) - カスタムイテレータ
- [Go 1.24](./go1.24/README_ja.md) - ジェネリック定義型、`os.Root`
- [Go 1.25](./go1.25/README_ja.md) - `testing/synctest`
- [Go 1.26](./go1.26/README_ja.md) - Green Tea GC、`errors.AsType`

## 🛠 使い方

各ディレクトリには `Makefile` または `go.mod` が含まれています。

### C++ の場合

```bash
cd cpp23
make run
```

### Go の場合

```bash
cd go1.26
go run .
```

---
*このプロジェクトは学習目的で作成されています。最新の情報については各言語の公式ドキュメントも併せて参照してください。*
