tumia mtandao
tumia jsoni
pakeji sarufi {
    andaa = unda(file) {
        config = fungua(file)
        configString = config.soma()
        configDict = jsoni.dikodi(configString)
        clientID = configDict["client_id"]
        clientSecret = configDict["client_secret"]
        params = {"client_id": clientID, "client_secret": clientSecret}
        tokenString = mtandao.tuma(yuareli="https://api.sarufi.io/api/access_token", mwili=params)
        tokenDict = jsoni.dikodi(tokenString)
        @.token = tokenDict["access_token"]
        @.Auth = "Bearer " + @.token
        }

    tokenYangu = unda() {
            rudisha @.token
        }

    tengenezaChatbot = unda(data) {
            majibu = mtandao.tuma(yuareli="https://api.sarufi.io/chatbot", vichwa={"Authorization": @.Auth}, mwili = data)
            rudisha majibu
        }

    pataChatbotZote = unda() {
            majibu = mtandao.peruzi(yuareli="https://api.sarufi.io/chatbots", vichwa={"Authorization": @.Auth})
            rudisha majibu
        }
}


