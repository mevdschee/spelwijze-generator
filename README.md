# spelwijze-genie

Download Dutch words from:

https://www.opentaal.org/bestanden/file/2-woordenlijst-v-2-10g-bronbestanden

To filter out non-letter characters execute:

    cat 'OpenTaal-210G-basis-gekeurd.txt' | grep -vP '[^a-z]' | sort | uniq > woordenlijst.txt

Now run pick a length for your seeding word:

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

    go run . espressoapparaat

Resulting in the number of words dependent on the first (mandatory) letter:

    taeoprs: 333
    raeopst: 315
    saeoprt: 293
    eaoprst: 275
    paeorst: 270
    oaeprst: 256
    aeoprst: 197

Now run to solve:

    go run . aeoprst

To find:

    aars
    aaseter
    aassoort
    aastor
    aorta
    apart
    aparte
    apert
    apetrots
    apoptose
    ...(187 more)...

Enjoy!