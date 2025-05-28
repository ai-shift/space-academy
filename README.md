# Space academy

## Traverser
### How to install
For MacOs with M series
Open your terminal and insert the following commands. You will be prompted to select folder where all your files live. Also you will be prompted for the password to ensure that all file in the selected directory can be accessed.
```bash
curl -L -o traverser https://github.com/ai-shift/space-academy/releases/download/0.1.3/traverser-darwin-arm64
chmod +x traverser
DRR=$(osascript -e 'choose folder')
sudo traverser --path "$DIR"
```
