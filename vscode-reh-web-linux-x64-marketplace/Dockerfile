FROM btwiuse/arch:vscode-src

RUN curl -sL https://gist.githubusercontent.com/btwiuse/ce8ec1815cf9501e35d5116dcf0955fb/raw/f5481743b9ab7bc62755c9d32b058f6c8f1a2b1b/patched.json > product.json

RUN yarn gulp vscode-reh-web-linux-x64

CMD ufo term
