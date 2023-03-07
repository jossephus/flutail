# flutail
## Tailwindcss's color System for Flutter

# Why

One thing I realized about css in the last few months was that having well defined color systems is very beneficial in your day to day workflow. Since tailwindcss has the best set of colors, this is me trying to bring tailwindcss's color system to flutter.. It's about 100 lines of dart code generated from the default tailwindcss configuration file  using the go files above.

## HOW TO USE
just copy the stubs/tailwindColors.dart in ur project and start using the color systems.

every color is prefixed with tail[color]. If you want to use tailwind's 'teal-100' it will be 'tailTeal100' here.

## TODO
- a flag to grap the configuration file directly from the network 
- a flag to specify prefix to use for the constants 
- a program that walks through all dart files in a project and  filter the tail colors used and edit the tailwindColors file  
- do u know how awesome tailiwind shadows are?? what about them 
- the font sizes too.. Problem is dart doesn't know about rems.. we will figure that out  

