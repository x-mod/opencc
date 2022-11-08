update\:data:
	# Fetch to update data from https://github.com/BYVoid/OpenCC
	mkdir -p ./tmp && rm -Rf tmp/OpenCC-master
	# wget https://github.com/BYVoid/OpenCC/archive/refs/heads/master.zip -O tmp/opencc.zip
	unzip tmp/opencc.zip -d tmp/
	cp -r tmp/OpenCC-master/data/dictionary/* dictionary/
	git diff dictionary/