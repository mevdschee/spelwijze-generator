# spelwijze-genie

Download Dutch words from:

https://www.opentaal.org/bestanden/file/2-woordenlijst-v-2-10g-bronbestanden

To filter out non-letter characters execute:

    cat 'OpenTaal-210G-basis-gekeurd.txt' | grep -vP '[^a-z]' | sort | uniq > woordenlijst.txt

Now run pick a length for your seeding word (a word with 7 different letters):

    go run . 16

Showing:

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

Resulting in the number of words dependent on the first (mandatory) letter:

    eblmnot: 210
    nbelmot: 165
    obelmnt: 163
    tbelmno: 151
    lbemnot: 118
    mbelnot: 110
    belmnot: 104

Now run to solve:

    go run . belmnot

To find:

    beboet
    beboeten
    been
    beet
    beetnemen
    belenen
    belet
    beletten
    belle
    bellen
    ...(94 more)...

Enjoy!