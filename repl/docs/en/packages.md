# Packages in Nuru

You can use third packages written in Nuru with the following conditions:

- The package file MUST be in the same directory
- The package file MUST end with `nr`
- The package name and package file MUST have the same name (eg: if `pakeji hesabu` then the file name must be `hesabu.nr`)
- The package must have the following structure:
```
// imports if any

pakeji [name of package] {
        andaa = unda() { // the andaa function is mandatory even if its empty

            }
        [body of package]
    }
```
- The package must be initialized with the `andaa` keyword (see above).

The `andaa` keyword is for initializing your package. This is also where you'd put your global variables. The global variables should be prefixed with `@.` Eg: `@.myGlobalVar`.

A variable being globally available means that the variable can be accessed and manipulated by all other methods in the package.


Below is an example Sarufi package:
```
// import modules
tumia mtandao
tumia jsoni

// package body
pakeji sarufi {

        // initialize function
        andaa = unda(file) {
            config = fungua(file) // read passwords from json file
            configString = config.soma()

            configDict = jsoni.dikodi(configString) // convert it to a dict
            clientID = configDict["client_id"]
            clientSecret = configDict["client_secret"]

            //  fill in params
            params = {"client_id": clientID, "client_secret": clientSecret}

            // get response
            resp = mtandao.tuma(yuareli="https://api.sarufi.io/api/access_token", mwili=params)
            tokenDict = jsoni.dikodi(resp)

            // extract token and make it globally available
            @.token = tokenDict["access_token"]

            // make the "Bearer <token>" globally available
            @.Auth = "Bearer " + @.token
            }

        // a method to get token
        tokenYangu = unda() {
                rudisha @.token
            }

        // a method to create new chatbots.
        // the data should be a dict
        tengenezaChatbot = unda(data) {
                majibu = mtandao.tuma(yuareli="https://api.sarufi.io/chatbot", vichwa={"Authorization": @.Auth}, mwili = data)
                rudisha majibu
            }

        // a method to get all available chatbots
        pataChatbotZote = unda() {
                majibu = mtandao.peruzi(yuareli="https://api.sarufi.io/chatbots", vichwa={"Authorization": @.Auth})
                rudisha majibu
            }
    }
```
