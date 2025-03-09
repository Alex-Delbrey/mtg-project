# mtg-project
Basically have to go directly into REST API since mtg-sdk-go is extremely depricated.

Goal:
    Basic - have a cli app that can search cards on mtg server and save to a postgres db the results, which would be my personal deck. Potentially having tables with decks, and each deck would be a a reference to a table of cards.
    Advance - have some level of AI implementation, to where I can identify a card through a picture and have AI look for that card based on the picture, then I could save the response from what it retrieved into my db.
    Dream - rebuild mtg-sdk-go or contribute to it so that it is updated
