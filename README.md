# loyalify
Hackathon Backend

Go based backend for customer loyality program on the Stellar blockchain

this code is only aimed at the testnet right now

this supports creating new accounts on the testnet using

/api/createAccount/{ID}

you can list royality programs using

/api/programs
or a single one with
/api/programs/{ID}

You need to exchange the used wallet id in the server.go to be able to issue new tokens
