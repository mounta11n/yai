build:
	go build -o yai

install:
	chmod +x yai && mv yai ~/.local/bin/

uninstall:
	rm ~/.local/bin/yai && rm ~/.config/yai.json
