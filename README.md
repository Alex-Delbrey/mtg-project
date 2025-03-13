# mtg-project
Basically have to go directly into REST API since mtg-sdk-go is extremely depricated. \
Thanks to the mtg-sdk-go, though it is deprecated and has not been modified since 2018, I still used (and modified a bit) of their structs in order to marshal/unmarshal json. [Check them out.](https://github.com/MagicTheGathering/mtg-sdk-go/tree/master)

Goal: \
    Basic - have a cli app that can search cards on mtg server and save to a postgres db the results, which would be my personal deck. Potentially having tables with decks, and each deck would be a a reference to a table of cards.
