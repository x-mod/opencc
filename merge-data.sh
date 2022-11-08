# Merge addition-dictionary/*.txt into dictionary/*.txt
cp -r tmp/OpenCC-master/data/dictionary/* dictionary/
cp -r tmp/OpenCC-master/data/config/* config/
git diff dictionary/

for x in addition-dictionary/*.txt; do 
  target="dictionary/$(basename $x .txt).txt"
  echo "Merging $x to $target"

  cat $x >> $target
done