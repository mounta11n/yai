build:
	go build -o yai

install:
	chmod +x yai && sudo mv yai /usr/bin/yai

sss:
	chmod +x yai && sudo mv yai /usr/bin/yai && sudo cp /usr/bin/yai /usr/bin/sss-yai

uninstall:
	rm /usr/bin/yai && rm ~/.config/yai.json
