grep -rl 'import "../Controls"' ./src/ui | xargs sed -i 's/import "..\/Controls"/\/\/import "..\/Controls"/g'
grep -rl 'import "../"' ./src/ui | xargs sed -i 's/import "..\/"/\/\/import "..\/"/g'
grep -rl 'import "Delegates/"' ./src/ui | xargs sed -i 's/import "Delegates\/"/\/\/import "Delegates\/"/g'
grep -rl 'import "Dialogs/"' ./src/ui | xargs sed -i 's/import "Dialogs\/"/\/\/import "Dialogs\/"/g'
grep -rl 'import "Controls"' ./src/ui | xargs sed -i 's/import "Controls"/\/\/import "Controls"/g'
grep -rl '// import "qrc:/ui/src/ui/"' ./src/ui | xargs sed -i 's/\/\/ import "qrc:\/ui\/src\/ui\/"/import "qrc:\/ui\/src\/ui\/"/g'
