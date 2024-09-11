# Spelwijze generator

"Spelwijze" is spelling game that is published on Dutch newspaper websites. You get 1 mandatory letter and 6 optional letters. The goal is to make as many 4 or more letter words with these 7 letters containing at least the mandatory letter once. This repository contains a tool to generate these puzzles.

### Sources

Download Dutch words from:

https://www.opentaal.org/bestanden/file/2-woordenlijst-v-2-10g-bronbestanden

To filter out non-letter characters (go from 164313 to 156280 words) execute:

    cat 'OpenTaal-210G-basis-gekeurd.txt' | grep -vP '[^a-z]' | sort | uniq | gzip > words.txt.gz

Download Dutch word freqencies from:

https://wortschatz.uni-leipzig.de/en/download/Dutch

To filter the word frequency list (go from 1000000 to 515630 words) execute:

    cat 'nld_mixed_2012_1M-words.txt' | cut -f 2,3 | tr A-Z a-z | grep -P '^[a-z]+\t' | gzip > wordfreq.txt.gz

The text files are gzipped to reduce space.

### Running

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

Resulting in 7 different 7 letter combinations (points based on word frequency):

    tbelmno: 728
    eblmnot: 725
    nbelmot: 711
    obelmnt: 534
    belmnot: 194
    mbelnot: 164
    lbemnot: 119

Now if we chose "mbelnot" (where "m" is the mandatory letter) we can run:

    go run . mbelnot

To find all 110 words containing the letter "m" and one or more of the other 6 letters:

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
