# Raindrop.io to Notion

Script to migrate your bookmarks from Raindrop.io to Notion. Written in Go.

## How to use

0. Export bookmarks from Raindrop.io

https://app.raindrop.io/settings/backups (html)

1. Create Notion integration

Follow offical Notion documentation and use default settings.
https://developers.notion.com/docs/getting-started

2. Put your integration secret key

[NotionSecret](https://github.com/kirillshevch/raindrop_to_notion/blob/master/main.go#L14)

3. Create a new page in Notion and share for following integration

4. Put your page UUID (for e.g. `e03226c471c94116b8fdfdeefb02b74565`)

[ParentID](https://github.com/kirillshevch/raindrop_to_notion/blob/master/main.go#L13)

5. Copy your exported HTML backup from Raindrop to local script folder and rename to `import.html`

```sh
cp ~/Downloads/Raindrop.io.html ~/workspace/raindrop_to_notion/import.html
```

6. Install dependencies and run the script

```sh
go get
go run main.go
```