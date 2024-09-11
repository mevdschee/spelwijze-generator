# spelwijze-genie

Spelwijze is spelling game that is published on Dutch newspaper websites. You get one mandatory letter and six optional letters. The goal is to make as many 4 or more letter words with these 7 letters containing at least the mandatory letter once. This repository contains a tool to generate these puzzles.

Download Dutch words from:

https://www.opentaal.org/bestanden/file/2-woordenlijst-v-2-10g-bronbestanden

To filter out non-letter characters execute:

    cat 'OpenTaal-210G-basis-gekeurd.txt' | grep -vP '[^a-z]' | sort | uniq > woordenlijst.txt

Now run pick a length for your seeding word (a word with 7 different letters):

    go run . 16

Showing all 16 letter words consisting of exactly 7 different letters:

    begijnenbeweging
    binnenduingebied
    bloembollenteelt
    concernonderdeel
    engineeringgroep
    espressoapparaat
    exercitieterrein
    geestesgestoorde
    herinterpreteren
    intentionaliteit
    ...(23 more)...

Now pick a seeding word and run:

    go run . bloembollenteelt

Resulting in 7 different 7 letter combinations (each having a different first (mandatory) letter):

    eblmnot: 210
    nbelmot: 165
    obelmnt: 163
    tbelmno: 151
    lbemnot: 118
    mbelnot: 110
    belmnot: 104

Now if we chose "belmnot" (where "b" is the mandatory letter) we can run:

    go run . belmnot

To find all words containing a "b" and one or more of the other letters:

    beetnemen
    bemeten
    benemen
    benoemen
    betomen
    betonelement
    betonmolen
    bloem
    bloembol
    bloembollenteelt
    ...(100 more)...

Enjoy!