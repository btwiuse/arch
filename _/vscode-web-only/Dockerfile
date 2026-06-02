FROM btwiuse/arch:vscode-src

RUN curl -sL https://gist.githubusercontent.com/btwiuse/ce8ec1815cf9501e35d5116dcf0955fb/raw/7b7fc2570aabadaa57e000ab7056ba41eda60ff5/product.json > product.json

RUN yarn gulp vscode-web-only

CMD ufo term
