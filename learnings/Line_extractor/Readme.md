# Results time!

It's Year-end, and you need to see what your dealerships sold this year, so you can tell the boss how well you've done! Each dealership is going to publish their sales data, and management need to know what's hot and what's not in the Solar Panel world right now

## Sales results format

Each dealership will publish one row. The fields in this row will be separated by commas. Each row may vary in length as described below.

A result will consist of:

1. The dealership name
2. A repeating set of pairs with the number of panels sold, and the panel code

So for example:

    Dubai-1, 11014, SP-1 , 17803, SX, 4923, HV, 2069, P
    Mumbai-3, 22547, SP-1, 9389, SP-2, 4829, SX, 3375, HV, 3371, P, 309, Mu

We want to transform this into a standard result that shows:

* the dealership name
* translates the panel code into the full name
* shows the percentages of each panel sold, by dealership

### Codes

* SP-1 -> Standard Panel 1
* SP-2 -> Standard Panel 2
* SX -> Solar Extension
* HV -> High Voltage
* P -> Premium
* Mu -> Multi-Panel